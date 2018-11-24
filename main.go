package main

import (
	"github.com/guitarpawat/portscan/api/cache"
	"github.com/guitarpawat/portscan/web"
	"log"
)

func main() {
	defer cache.CloseDB()

	log.Println("START:\t", "running on port 80")

	log.Fatalln(web.ListenAndServe(":80"))
}
