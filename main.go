package main

import (
	"encoding/json"
	"fmt"
	"github.com/guitarpawat/portscan/api"
	"github.com/guitarpawat/portscan/api/cache"
	"github.com/guitarpawat/portscan/api/model"
	"log"
	"time"
)

func main() {
	defer cache.CloseDB()

	input := model.ScanInput{
		Targets: []model.Target{
			{
				Address: "www.google.com",
				Ports:   []int{80, 433, 443, 27017},
			},
			{
				Address: "127.0.0.1",
				Ports:   []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 200, 300, 400, 433, 443, 500, 4433, 8000, 27017},
			},
			//{
			//	Address: "www.google.com",
			//	Ports: []int{80, 443, 433, 500},
			//},
			//{
			//	Address: "127.0.0.1",
			//	Ports: []int{80, 443, 433, 500},
			//},
		},
	}

	b, err := json.Marshal(input)
	if err != nil {
		log.Println(err)
		return
	}

	res := api.PutNewScanRequest(b).(*model.ScanOutput)
	fmt.Println(res)

	token := res.Token
	getin := model.GetInput{token}
	b, err = getin.Marshal()
	if err != nil {
		log.Println(err)
		return
	}

	errjson := api.KillScanRequest(b)
	if errjson != nil {
		fmt.Println(errjson)
	} else {
		fmt.Println("deleted!")
	}

	for {
		out, _ := api.GetUpdateScanResult(b).Marshal()
		fmt.Println(string(out))
		time.Sleep(time.Second * 1)
	}
}
