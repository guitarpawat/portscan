package main

import (
	"fmt"
	"github.com/guitarpawat/portscan/scanner"
)

func main() {
	ipv4, _, _ := scanner.NSLookUp("www.google.com")
	fmt.Println(scanner.GetOpenPorts(ipv4, 80, 443))
	ipv6, _, _ := scanner.NSLookUp("www.google.com")
	fmt.Println(scanner.GetOpenPorts(ipv6, 80, 443))
}
