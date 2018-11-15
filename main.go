package main

import (
	"fmt"
	"github.com/guitarpawat/portscan/scanner"
)

func main() {
	fmt.Println(scanner.NSLookUp("localhost"))
}
