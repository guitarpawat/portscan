package main

import (
	"fmt"
	"github.com/guitarpawat/portscan/scanner"
)

func main() {
	ipv4, _, _ := scanner.NSLookUp("www.google.com")
	fmt.Println(scanner.GetOpenPorts(ipv4, 80, 100, 200, 300, 400, 443, 500))
	ipv6, _, _ := scanner.NSLookUp("www.google.com")
	fmt.Println(scanner.GetOpenPorts(ipv6, 80, 100, 200, 300, 400, 443, 500))
}
