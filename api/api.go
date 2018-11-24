package api

import (
	"github.com/chilts/sid"
	"github.com/guitarpawat/portscan/api/cache"
	"github.com/guitarpawat/portscan/api/model"
	"github.com/guitarpawat/portscan/scanner"
	"time"
)

const limit = 2048

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

	token := sid.Id()
	ip := make([]string, len(input.Targets))
	host := make([]string, len(input.Targets))

	for i, t := range input.Targets {
		if scanner.IsV4(t.Address) {
			ip[i] = t.Address
			host[i] = ""

		} else if scanner.IsV6(t.Address) {
			return model.MakeErrorString("ipv6 is not currently support")

		} else {
			add, err := scanner.NSLookUp(t.Address)
			if err != nil {
				return model.MakeError(err)
			}

			ip[i] = add
			host[i] = t.Address
			input.Targets[i].Address = add
		}

	}

	err = cache.PutNewToken(token, ip...)
	if err != nil {
		return model.MakeError(err)
	}

	sem := make(chan struct{}, limit)
	registerToken(token)

	for i := 0; i < len(input.Targets); i++ {
		routines.Lock()
		_, ok := routines.Tasks[token]
		routines.Unlock()
		if !ok {
			break
		}
		kill := make(chan struct{}, 1)
		registerChan(token, kill)
		go func(sem, kill chan struct{}, host, ip string, ports ...int) {
			sem <- struct{}{}
			defer func() { <-sem }()
			open := scanner.GetOpenPorts(ip, kill, ports...)
			cache.UpdateTokenInfo(token, model.MakeResult(host, ip, open...))
		}(sem, kill, host[i], ip[i], input.Targets[i].Ports...)
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
	if err != nil {
		return model.MakeError(err)
	}

	return &out
}

// KillScanRequest is RESTful API request for delete unused port scan task
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
