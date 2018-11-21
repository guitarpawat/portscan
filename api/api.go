package api

import (
	"github.com/chilts/sid"
	"github.com/guitarpawat/portscan/api/cache"
	"github.com/guitarpawat/portscan/api/model"
	"github.com/guitarpawat/portscan/scanner"
)

func PutNewScanRequest(b []byte) model.Json {
	input, err := model.UnMarshalScanInput(b)
	if err != nil {
		return model.MakeError(err)
	}

	var token = sid.Id()
	var ip []string

	for _, t := range input.Targets {
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
		}
	}

	err = cache.PutNewToken(token, ip...)
	if err != nil {
		return model.MakeError(err)
	}

	// TODO: Add goroutine to scan the host concurrently

	return &model.ScanOutput{
		Token: token,
		IP:    ip,
	}
}
