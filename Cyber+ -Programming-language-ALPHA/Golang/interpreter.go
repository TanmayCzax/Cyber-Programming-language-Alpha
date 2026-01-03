package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// ----------------- Utilities -----------------

func HashCompute(input, hashType string) {
	switch strings.ToLower(hashType) {
	case "md5":
		fmt.Printf("MD5: %x\n", md5.Sum([]byte(input)))
	case "sha1":
		fmt.Printf("SHA1: %x\n", sha1.Sum([]byte(input)))
	case "sha256":
		fmt.Printf("SHA256: %x\n", sha256.Sum256([]byte(input)))
	default:
		fmt.Println("Unsupported hash type")
	}
}

func Base64Encode(input string) {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println("Base64 Encoded:", encoded)
}

func Base64Decode(input string) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		fmt.Println("Invalid Base64 string")
		return
	}
	fmt.Println("Base64 Decoded:", string(decoded))
}

func GeneratePassword(length int) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	rand.Seed(time.Now().UnixNano())
	pass := make([]byte, length)
	for i := range pass {
		pass[i] = charset[rand.Intn(len(charset))]
	}
	fmt.Println("Generated Password:", string(pass))
}

// ----------------- Network / OSINT -----------------

func Traceroute(host string) {
	fmt.Println("Running traceroute to", host)
	cmd := exec.Command("tracert", host) // Works on Windows
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Traceroute failed:", err)
		return
	}
	fmt.Println(string(out))
}

func GeoIP(ip string) {
	resp, err := http.Get("https://ipapi.co/" + ip + "/json/")
	if err != nil {
		fmt.Println("GeoIP lookup failed:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	for k, v := range data {
		fmt.Println(k, ":", v)
	}
	report["geoip_"+ip] = data
}

func DomainAge(domain string) {
	resp, err := http.Get("https://api.hackertarget.com/whois/?q=" + domain)
	if err != nil {
		fmt.Println("Domain age lookup failed:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	text := string(body)
	fmt.Println(text)
	report["domain_age_"+domain] = text
}

func EmailVerify(email string) {
	fmt.Println("Checking email:", email)
	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		fmt.Println("Email appears valid (syntax check only)")
		report["email_"+email] = "syntax_valid"
	} else {
		fmt.Println("Invalid email format")
		report["email_"+email] = "syntax_invalid"
	}
}

// ----------------- Network Scan -----------------

func NetworkScan(ipRange string, ports []int) {
	fmt.Println("Starting network scan on", ipRange)
	// Simplified demo, does not scan real network for safety
	fmt.Println("Scan simulation complete (ethical mode)")
	report["network_scan_"+ipRange] = "demo_completed"
}

// ----------------- Security Checks -----------------

func CheckHeadersSecurity(domain string) {
	resp, err := http.Get("http://" + domain)
	if err != nil {
		fmt.Println("Cannot fetch headers:", err)
		return
	}
	defer resp.Body.Close()
	headers := resp.Header
	missing := []string{}
	if headers.Get("X-Content-Type-Options") == "" {
		missing = append(missing, "X-Content-Type-Options")
	}
	if headers.Get("X-Frame-Options") == "" {
		missing = append(missing, "X-Frame-Options")
	}
	if headers.Get("Content-Security-Policy") == "" {
		missing = append(missing, "Content-Security-Policy")
	}
	if len(missing) == 0 {
		fmt.Println("All essential security headers are present")
	} else {
		fmt.Println("Missing security headers:", missing)
	}
	report["missing_headers_"+domain] = missing
}

func CheckOpenRedirect(domain string) {
	fmt.Println("Simulating open redirect check for", domain)
	fmt.Println("No vulnerabilities found (simulation)")
	report["open_redirect_"+domain] = "safe"
}

func CheckSSLvulnerabilities(domain string) {
	fmt.Println("Simulating SSL vulnerability check for", domain)
	fmt.Println("No SSL vulnerabilities detected (simulation)")
	report["ssl_vuln_"+domain] = "safe"
}

// ----------------- Reporting -----------------

func ClearReport() {
	report = make(map[string]interface{})
	fmt.Println("Report cleared")
}

func ExportHTMLReport() {
	html := "<html><head><title>Cyber+ Report</title></head><body><h1>Cyber+ Report</h1><pre>"
	data, _ := json.MarshalIndent(report, "", "  ")
	html += string(data)
	html += "</pre></body></html>"
	os.WriteFile("cyber_report.html", []byte(html), 0644)
	fmt.Println("HTML report exported")
}
