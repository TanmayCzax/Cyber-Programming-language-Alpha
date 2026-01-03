------------Cyber+ Programming Language------------

Cyber+ is a minimal, powerful, and ethical cyber security–focused programming language designed for reconnaissance, OSINT, and security analysis.

It is inspired by Go’s simplicity and built for clarity, speed, and real-world usage, not gimmicks.

Cyber+ is strictly for ethical hacking, learning, and defensive security.

Why Cyber+

Minimal & readable syntax

Cyber-security–only commands (no noise)

Beginner-friendly, professional output

CLI-based & scriptable

Easy to extend (written in Go)

VS Code extension support

-------- Installation ----------
Requirements

Go 1.21+

Git

VS Code (optional, for extension)

Clone
git clone https://github.com/YOUR_USERNAME/CyberPlus.git
cd CyberPlus

Run
go run .

 Running a Cyber+ Script

-------------Create a file:------------

test.cyber


-----------Run:--------------

cyber run test.cyber

 Cyber+ Syntax Example
-- run Cyber+

Compute("Cyber+ Recon Started");

Ping("google.com");
Scan_Port("192.168.1.1", 80);

Dns_Lookup("google.com");
Subdomain_Enum("google.com");

IP_Info("8.8.8.8");
Reverse_IP("8.8.8.8");

HTTP_Headers("google.com");
URL_Status("https://google.com");

SSL_Info("google.com");
Whois("google.com");

Phone_Info("+919009800061");

Recon("google.com");

Export_Report("json");
Export_Report("txt");

 ------------Available Commands (17)--------------
 Engine Control

-- run Cyber+

 Core Output

Compute("text");

 Network & Connectivity

Ping("host");

Scan_Port("ip", port);

 DNS & Domain Recon

Dns_Lookup("domain");

Subdomain_Enum("domain");

 IP Intelligence

IP_Info("ip");

Reverse_IP("ip");

 Web & HTTP Security

HTTP_Headers("domain");

URL_Status("url");

 SSL / Certificate

SSL_Info("domain");

 WHOIS Intelligence

Whois("domain");

 Phone OSINT

Phone_Info("number");

 Automated Recon

Recon("domain");

 Reporting

Export_Report("json");

Export_Report("txt");


 ---------------Ethical Usage Policy--------------------

Cyber+ must only be used on systems you own or have permission to test.

 Illegal usage is strictly prohibited
 Learning, labs, bug bounties, defense only

-------------------Roadmap----------------------

 Auto-completion (IntelliSense)

 Script variables & flow control

 Binary releases

Contributing

Pull requests are welcome.
Please open an issue before major changes.

----------- License ----------

MIT License
© 2025 Cyber+

