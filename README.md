# Cyber Plus Alpha (CPA)
The first cybersecurity-only programming language.

## Why Cyber+?
Cyber+ eliminates the need to write long scripts and install heavy libraries
for common cybersecurity tasks.

## Example
Python (15+ lines) vs Cyber+ (1 line)

## Features
- Fast (Go-powered)
- Security-focused DSL
- Beginner friendly
- Cross-platform

## Installation
(one command)

## Commands
(list everything)

## Use cases
- Students
- Security researchers
- Automation

Cyber+ Alpha is a new version of Cyber+ 0.1

# Use Cyber+ for Ethical purpose.
Here are example codes of Cyber+ with there uses -- 

# 1 Engine Control

-- run Cyber+
Activates the Cyber+ execution engine.

# 2 Core / Output

Compute("text");
Prints text output to the console.

# 3 Network Scanning & Connectivity

Scan_Port("ip", port);
Checks whether a TCP port is open or closed.

Ping("host");
Sends a single ping request to a host.

Traceroute("host");
Displays the network path to a host.

NetworkScan("ip_start-ip_end", [ports]);
Simulated ethical network scan over an IP range.

# 4 DNS & Domain Intelligence

Dns_Lookup("domain");
Resolves domain to IP addresses.

Subdomain_Enum("domain");
Enumerates subdomains using certificate transparency logs.

Reverse_IP("ip");
Resolves domain names hosted on an IP.

DomainAge("domain");
Retrieves domain registration information.

# 5 IP Intelligence & OSINT

IP_Info("ip");
Fetches detailed IP information.

GeoIP("ip");
Retrieves geographical data of an IP address.

# 6 Web & URL Analysis

URL_Status("url");
Checks HTTP response status of a URL.

HTTP_Headers("domain");
Displays all HTTP response headers.

Check_Headers_Security("domain");
Detects missing security-related HTTP headers.

# 7 SSL / TLS Analysis

SSL_Info("domain");
Displays SSL certificate issuer and expiry.

CheckSSLvulnerabilities("domain");
Simulated SSL vulnerability analysis.

# 8 WHOIS & Registration

Whois("domain");
Retrieves WHOIS data for a domain.

# 9 OSINT – Phone & Email

Phone_Info("phone_number");
Provides phone number intelligence.

EmailVerify("email");
Performs email format validation.

# 10 Encoding, Hashing & Utilities

Hash_Compute("text", "algorithm");
Generates hashes using md5, sha1, or sha256.

Base64_Encode("text");
Encodes text using Base64.

Base64_Decode("encoded_text");
Decodes Base64-encoded text.

Generate_Password(length);
Generates a random secure password.

# 11 Recon Automation

Recon("domain");
Performs automated reconnaissance (DNS, subdomains, IP info, ports, headers).

# 12 Web Vulnerability Checks

CheckOpenRedirect("domain");
Simulated open redirect vulnerability check.

# 13 Reporting & Data Management

Export_Report("json|txt");
Exports collected data to a report file.

Export_HTML_Report();
Exports collected data as an HTML report.

ClearReport();
Clears all stored report data.

# Summary

Total commands: 29

Type: Ethical cybersecurity DSL

Platform: Go-based standalone executable

Focus: Reconnaissance, OSINT, analysis, reporting

# Comparison with Python
Python code for Hashing --
import hashlib

h = hashlib.sha256(b"password").hexdigest()

print(h)

Cyber+ Alpha (CPA) command for Hashing --

Hash_Compute("password", "sha256");

# Thank you for Downloading
# Cyber+ Team {VERSION - ALPHA (1.0)}
