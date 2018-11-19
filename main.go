package main

import (
	"fmt"

	"github.com/guitarpawat/portscan/api"
	"github.com/guitarpawat/portscan/scanner"
)

func main() {
	ipv4, ipv6, _ := scanner.NSLookUp("www.google.com")
	knowPorts := api.GetKnownPorts()
	fmt.Println(knowPorts)
	fmt.Println(scanner.GetOpenPorts(ipv4, knowPorts...))
	fmt.Println(scanner.GetOpenPorts(ipv6, knowPorts...))
}
