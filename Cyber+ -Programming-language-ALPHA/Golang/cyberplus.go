package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	phonenumbers "github.com/nyaruka/phonenumbers"
)

var engineActive bool
var report = make(map[string]interface{})

func banner() {
	fmt.Println("Cyber+ — Ethical Cyber Security Language")
	fmt.Println("=======================================")
}

func executeCompute(msg string) {
	fmt.Println(msg)
}

func executeScanPort(ip, port string) {
	conn, err := net.DialTimeout("tcp", ip+":"+port, 2*time.Second)
	status := "CLOSED"
	if err == nil {
		status = "OPEN"
		conn.Close()
	}
	fmt.Printf("Port %s %s\n", port, status)
	report["port_"+port] = status
}

func executePhoneInfo(number string) {
	fmt.Println("---- Phone Intelligence ----")

	num, err := phonenumbers.Parse(number, "")
	if err != nil {
		fmt.Println("Invalid phone number")
		return
	}

	valid := phonenumbers.IsValidNumber(num)
	possible := phonenumbers.IsPossibleNumber(num)
	numberType := phonenumbers.GetNumberType(num)

	typeMap := map[phonenumbers.PhoneNumberType]string{
		phonenumbers.MOBILE:     "Mobile",
		phonenumbers.FIXED_LINE: "Fixed Line",
		phonenumbers.VOIP:       "VOIP",
	}
	nType := typeMap[numberType]
	if nType == "" {
		nType = "Unknown"
	}

	carrier, _ := phonenumbers.GetCarrierForNumber(num, "en")
	timezones, _ := phonenumbers.GetTimezonesForNumber(num)

	fmt.Println("Valid:", valid)
	fmt.Println("Possible:", possible)
	fmt.Println("Type:", nType)
	fmt.Println("Carrier:", carrier)
	fmt.Println("Timezones:", timezones)

	report["phone_info"] = map[string]interface{}{
		"valid":     valid,
		"possible":  possible,
		"type":      nType,
		"carrier":   carrier,
		"timezones": timezones,
	}
}

func executeDnsLookup(domain string) []string {
	ips, err := net.LookupHost(domain)
	if err != nil {
		fmt.Println("DNS failed")
		return nil
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	report["dns"] = ips
	return ips
}

func executePing(host string) {
	cmd := exec.Command("ping", "-n", "1", host)
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))
}

func executeSubdomainEnum(domain string) []string {
	url := "https://crt.sh/?q=%25." + domain + "&output=json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Subdomain fetch failed")
		return nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data []map[string]interface{}
	json.Unmarshal(body, &data)

	found := map[string]bool{}
	results := []string{}

	for _, e := range data {
		if v, ok := e["name_value"].(string); ok {
			for _, s := range strings.Split(v, "\n") {
				if strings.HasSuffix(s, domain) && !found[s] {
					found[s] = true
					results = append(results, s)
					fmt.Println(s)
				}
			}
		}
	}
	report["subdomains"] = results
	return results
}

func executeIPInfo(ip string) {
	resp, _ := http.Get("https://ipinfo.io/" + ip + "/json")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	for k, v := range data {
		fmt.Println(k, ":", v)
	}
	report["ip"] = data
}

func executeReverseIP(ip string) {
	names, err := net.LookupAddr(ip)
	if err != nil {
		fmt.Println("Reverse IP failed")
		return
	}
	for _, n := range names {
		fmt.Println(n)
	}
	report["reverse_ip"] = names
}

func executeURLStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("URL unreachable")
		return
	}
	fmt.Println("Status:", resp.Status)
	report["url_status"] = resp.Status
}

func executeSSLInfo(domain string) {
	conn, err := tls.Dial("tcp", domain+":443", nil)
	if err != nil {
		fmt.Println("SSL connection failed")
		return
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	fmt.Println("Issuer:", cert.Issuer)
	fmt.Println("Expires:", cert.NotAfter)

	report["ssl"] = map[string]interface{}{
		"issuer":  cert.Issuer,
		"expires": cert.NotAfter,
	}
}

func executeWhois(domain string) {
	resp, err := http.Get("https://api.hackertarget.com/whois/?q=" + domain)
	if err != nil {
		fmt.Println("Whois failed")
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	report["whois"] = string(body)
}

/* ------------------------------------- */

func executeHTTPHeaders(domain string) {
	resp, _ := http.Get("http://" + domain)
	defer resp.Body.Close()
	headers := map[string]string{}
	for k, v := range resp.Header {
		headers[k] = strings.Join(v, ", ")
		fmt.Println(k+":", headers[k])
	}
	report["headers"] = headers
}

func executeRecon(domain string) {
	ips := executeDnsLookup(domain)
	executeSubdomainEnum(domain)
	if len(ips) > 0 {
		executeIPInfo(ips[0])
		executeScanPort(ips[0], "80")
		executeScanPort(ips[0], "443")
	}
	executeHTTPHeaders(domain)
	fmt.Println("Recon completed")
}

func exportReport(format string) {
	data, _ := json.MarshalIndent(report, "", "  ")
	os.WriteFile("cyber_report."+format, data, 0644)
	fmt.Println("Report exported")
}

func main() {
	banner()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Cybp > ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "-- run Cyber+" {
			engineActive = true
			fmt.Println("Cyber+ activated")
			continue
		}

		if !engineActive {
			fmt.Println("To run Cyber+ code Run '-- run Cyber+' first")
			continue
		}

		rules := []struct {
			re *regexp.Regexp
			fn func([]string)
		}{
			// ---------- Existing commands ----------
			{regexp.MustCompile(`^Compute\("(.+)"\);$`), func(m []string) { executeCompute(m[1]) }},
			{regexp.MustCompile(`^Scan_Port\("(.+)",\s*(\d+)\);$`), func(m []string) { executeScanPort(m[1], m[2]) }},
			{regexp.MustCompile(`^Phone_Info\("(.+)"\);$`), func(m []string) { executePhoneInfo(m[1]) }},
			{regexp.MustCompile(`^Dns_Lookup\("(.+)"\);$`), func(m []string) { executeDnsLookup(m[1]) }},
			{regexp.MustCompile(`^Ping\("(.+)"\);$`), func(m []string) { executePing(m[1]) }},
			{regexp.MustCompile(`^Subdomain_Enum\("(.+)"\);$`), func(m []string) { executeSubdomainEnum(m[1]) }},
			{regexp.MustCompile(`^IP_Info\("(.+)"\);$`), func(m []string) { executeIPInfo(m[1]) }},
			{regexp.MustCompile(`^Reverse_IP\("(.+)"\);$`), func(m []string) { executeReverseIP(m[1]) }},
			{regexp.MustCompile(`^URL_Status\("(.+)"\);$`), func(m []string) { executeURLStatus(m[1]) }},
			{regexp.MustCompile(`^SSL_Info\("(.+)"\);$`), func(m []string) { executeSSLInfo(m[1]) }},
			{regexp.MustCompile(`^Whois\("(.+)"\);$`), func(m []string) { executeWhois(m[1]) }},
			{regexp.MustCompile(`^HTTP_Headers\("(.+)"\);$`), func(m []string) { executeHTTPHeaders(m[1]) }},
			{regexp.MustCompile(`^Recon\("(.+)"\);$`), func(m []string) { executeRecon(m[1]) }},
			{regexp.MustCompile(`^Export_Report\("(json|txt)"\);$`), func(m []string) { exportReport(m[1]) }},

			// ---------- New commands from interpreter.go ----------
			{regexp.MustCompile(`^Hash_Compute\("(.+)",\s*"(.+)"\);$`), func(m []string) { HashCompute(m[1], m[2]) }},
			{regexp.MustCompile(`^Base64_Encode\("(.+)"\);$`), func(m []string) { Base64Encode(m[1]) }},
			{regexp.MustCompile(`^Base64_Decode\("(.+)"\);$`), func(m []string) { Base64Decode(m[1]) }},
			{regexp.MustCompile(`^Generate_Password\((\d+)\);$`), func(m []string) {
				n, _ := strconv.Atoi(m[1])
				GeneratePassword(n)
			}},
			{regexp.MustCompile(`^Traceroute\("(.+)"\);$`), func(m []string) { Traceroute(m[1]) }},
			{regexp.MustCompile(`^GeoIP\("(.+)"\);$`), func(m []string) { GeoIP(m[1]) }},
			{regexp.MustCompile(`^DomainAge\("(.+)"\);$`), func(m []string) { DomainAge(m[1]) }},
			{regexp.MustCompile(`^EmailVerify\("(.+)"\);$`), func(m []string) { EmailVerify(m[1]) }},
			{regexp.MustCompile(`^NetworkScan\("(.+)",\s*\[(.*)\]\);$`), func(m []string) {
				ports := []int{}
				if m[2] != "" {
					for _, p := range strings.Split(m[2], ",") {
						p = strings.TrimSpace(p)
						n, _ := strconv.Atoi(p)
						ports = append(ports, n)
					}
				}
				NetworkScan(m[1], ports)
			}},
			{regexp.MustCompile(`^Check_Headers_Security\("(.+)"\);$`), func(m []string) { CheckHeadersSecurity(m[1]) }},
			{regexp.MustCompile(`^CheckOpenRedirect\("(.+)"\);$`), func(m []string) { CheckOpenRedirect(m[1]) }},
			{regexp.MustCompile(`^CheckSSLvulnerabilities\("(.+)"\);$`), func(m []string) { CheckSSLvulnerabilities(m[1]) }},
			{regexp.MustCompile(`^ClearReport\(\);$`), func(m []string) { ClearReport() }},
			{regexp.MustCompile(`^Export_HTML_Report\(\);$`), func(m []string) { ExportHTMLReport() }},
		}

		for _, r := range rules {
			if m := r.re.FindStringSubmatch(input); len(m) > 0 {
				r.fn(m)
				break
			}
		}
	}
}
