package model

import "encoding/json"

// ScanInput is the json model for port scanning request
type ScanInput struct {
	Targets []Target `json:"targets"`
}

// Target of the port scanner, helper for ScanInput
type Target struct {
	Address string `json:"address"`
	Ports   []int  `json:"ports"`
}

// ScanOutput is the json model for returning port scanning request with Token id
type ScanOutput struct {
	Token string   `json:"token"`
	IP    []string `json:"ip"`
}

// Marshal the ScanOutput into JSON bytes
func (out *ScanOutput) Marshal() ([]byte, error) {
	return json.Marshal(out)
}

// UnMarshalScanInput converts JSON bytes into ScanInput
func UnMarshalScanInput(b []byte) (*ScanInput, error) {
	out := new(ScanInput)
	err := json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
