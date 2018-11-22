package main

import (
	"fmt"
	"github.com/guitarpawat/portscan/api/cache"
	"github.com/guitarpawat/portscan/web"
	"log"
)

func main() {
	defer cache.CloseDB()

	fmt.Println("Running on port 80")

	log.Fatalln(web.ListenAndServe(":80"))
}
