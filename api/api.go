package api

import (
	"fmt"
	"github.com/chilts/sid"
	"github.com/guitarpawat/portscan/api/cache"
	"github.com/guitarpawat/portscan/api/model"
	"github.com/guitarpawat/portscan/scanner"
	"time"
)

const limit = 8

func init() {
	go func() {
		for {
			time.Sleep(3 * time.Minute)
			go killTimeoutBatch()
		}
	}()
}

// PutNewScanRequest is RESTful API request for add new port scan task
//
// Accept: ScanInput, Returns: ScanOutput or Error
func PutNewScanRequest(b []byte) model.Json {
	input, err := model.UnMarshalScanInput(b)
	if err != nil {
		return model.MakeError(err)
	}

	var token = sid.Id()
	var ip []string

	for i, t := range input.Targets {
		if scanner.IsV4(t.Address) {
			ip = append(ip, t.Address)

		} else if scanner.IsV6(t.Address) {
			return model.MakeErrorString("ipv6 is not currently support")

		} else {
			add, err := scanner.NSLookUp(t.Address)
			if err != nil {
				return model.MakeError(err)
			}

			ip = append(ip, add)
			input.Targets[i].Address = add
		}

	}

	fmt.Println(ip)

	err = cache.PutNewToken(token, ip...)
	if err != nil {
		return model.MakeError(err)
	}

	//sem := make(chan struct{}, limit)
	registerToken(token)

	for _, ip := range input.Targets {
		fmt.Println("IP", ip)
		routines.Lock()
		_, ok := routines.Tasks[token]
		routines.Unlock()
		if !ok {
			break
		}
		kill := make(chan struct{}, 1)
		registerChan(token, kill)
		go func(sem, kill chan struct{}, ip string, ports ...int) {
			//sem <- struct{}{}
			//defer func() { <-sem }()
			open := scanner.GetOpenPorts(ip, kill, ports...)
			cache.UpdateTokenInfo(token, model.MakeResult(ip, open...))
		}(nil, kill, ip.Address, ip.Ports...)
	}

	return &model.ScanOutput{
		Token: token,
		IP:    ip,
	}
}

// GetUpdateScanResult is RESTful API request for getting latest update of port scan task
//
// Accept: GetInput, Returns: GetOutput or Error
func GetUpdateScanResult(b []byte) model.Json {
	input, err := model.UnMarshalGetInput(b)
	if err != nil {
		return model.MakeError(err)
	}

	out, err := cache.GetTokenInfo(input.Token)
	//fmt.Println(out)
	if err != nil {
		return model.MakeError(err)
	}

	return &out
}

// GetUpdateScanResult is RESTful API request for delete unused port scan task
//
// Accept: GetInput, Returns: nil or Error
func KillScanRequest(b []byte) model.Json {
	in, err := model.UnMarshalGetInput(b)
	if err != nil {
		return model.MakeError(err)
	}

	revokeChan(in.Token)

	err = cache.DeleteToken(in.Token)
	if err != nil {
		return model.MakeError(err)
	}

	return nil
}

func killTimeoutBatch() {
	routines.Lock()
	defer routines.Unlock()
	killTimeOut()
}
