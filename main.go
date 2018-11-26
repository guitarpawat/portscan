package main

import (
	"github.com/guitarpawat/portscan/api/cache"
	"github.com/guitarpawat/portscan/web"
	"github.com/pkg/browser"
	"log"
	"time"
)

func main() {
	defer cache.CloseDB()

	log.Println("START:\t", "running on port 80")

	go func() {
		time.Sleep(2 * time.Second)
		err := browser.OpenURL("http://localhost:80")
		if err != nil {
			log.Println(err)
		}
	}()

	log.Fatalln(web.ListenAndServe(":80"))
}
