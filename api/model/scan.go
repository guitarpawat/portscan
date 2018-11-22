package model

import "encoding/json"

// ScanInput is the json model for port scanning request
type ScanInput struct {
	Targets []Target `json:"targets"`
}

type Target struct {
	Address string `json:"address"`
	Ports   []int  `json:"ports"`
}

// ScanOutput is the json model for returning port scanning request with Token id
type ScanOutput struct {
	Token string   `json:"Token"`
	IP    []string `json:"ip"`
}

// TODO: Add docs and tests
func (out *ScanOutput) Marshal() ([]byte, error) {
	return json.Marshal(out)
}

func UnMarshalScanInput(b []byte) (*ScanInput, error) {
	out := new(ScanInput)
	err := json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}