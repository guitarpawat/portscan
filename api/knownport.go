package api

const unknown = "<unknown>"

var knownPorts = map[int]string{
	20:    "File Transfer Protocol (FTP) data transfer",
	21:    "File Transfer Protocol (FTP) control (command)",
	22:    "Secure Shell (SSH)",
	23:    "Telnet",
	25:    "Simple Mail Transfer Protocol (SMTP)",
	37:    "Time Protocol",
	53:    "Domain Name Server (DNS)",
	67:    "Dynamic Host Configuration Protocol (DHCP)",
	68:    "Dynamic Host Configuration Protocol (DHCP)",
	69:    "Trivial File Transfer Protocol (TFTP)",
	80:    "HyperText Transfer Protocol (HTTP)",
	110:   "Post Office Protocol (POP3)",
	119:   "Network News Transport Protocol (NNTP)",
	123:   "Network Time Protocol (NTP)",
	137:   "NetBIOS Name Service",
	138:   "NetBIOS Datagram Service",
	139:   "NetBIOS Session Service",
	143:   "Internet Message Access Protocol (IMAP4)",
	161:   "Simple Network Management Protocol (SNMP)",
	162:   "Simple Network Management Protocol Trap (SNMPTRAP)",
	194:   "Internet Relay Chat (IRC)",
	389:   "Lightweight Directory Access Protocol",
	443:   "HTTP with Secure Sockets Layer (SSL)",
	1194:  "Open VPN",
	1433:  "Microsoft SQL Server database management system (MSSQL) server",
	1434:  "Microsoft SQL Server database management system (MSSQL) monitor",
	2222:  "DirectAdmin Access",
	3306:  "MySQL database system",
	5432:  "PostgreSQL database system",
	8008:  "Alternative port for HTTP",
	8080:  "Alternative port for HTTP/ Apache Tomcat",
	27017: "MongoDB daemon process",
}

func GetKnownPorts() []int {
	ports := []int{}
	for port, _ := range knownPorts {
		ports = append(ports, port)
	}
	return ports
}
