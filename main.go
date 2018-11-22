package main

import (
	"github.com/guitarpawat/portscan/api/cache"
	"github.com/guitarpawat/portscan/web"
	"log"
)

func main() {
	defer cache.CloseDB()

	log.Fatalln(web.ListenAndServe(":80"))
}
