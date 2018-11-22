package model

import (
	"encoding/json"
	"testing"
)

func TestScanOutput_Marshal(t *testing.T) {
	out := ScanOutput{
		Token: "asdf",
		IP:    []string{"127.0.0.1"},
	}

	b, _ := out.Marshal()
	res := ScanOutput{}
	json.Unmarshal(b, &res)

	if out.Token != res.Token {
		t.Errorf("expected token: %s, but get: %s", out.Token, res.Token)
	}

	if out.IP[0] != res.IP[0] {
		t.Errorf("expected ip: %s, but get: %s", out.IP[0], res.IP[0])
	}
}

func TestUnMarshalScanInput(t *testing.T) {
	in := ScanInput{
		Targets: []Target{
			{
				Address: "www.google.com",
				Ports:   []int{80},
			},
		},
	}

	b, _ := json.Marshal(in)
	out, _ := UnMarshalScanInput(b)

	if in.Targets[0].Address != out.Targets[0].Address {
		t.Errorf("expected address: %s, but get: %s",
			in.Targets[0].Address, out.Targets[0].Address)
	}

	if in.Targets[0].Ports[0] != out.Targets[0].Ports[0] {
		t.Errorf("expected port: %d, but get: %d",
			in.Targets[0].Ports[0], out.Targets[0].Ports[0])
	}
}
