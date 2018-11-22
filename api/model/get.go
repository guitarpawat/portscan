package model

import (
	"encoding/json"
	"github.com/guitarpawat/portscan/api/portdes"
	"time"
)

// GetInput is json model for getting result of specific Token id
type GetInput struct {
	Token string `json:"Token"`
}

// GetOutput is json model for returning result of specific Token id
type GetOutput struct {
	Results    []Result  `json:"results"`
	LastUpdate time.Time `json:"last_update"`
}

// Result stores the open ports of specify IP
type Result struct {
	IP       string `json:"ip"`
	Ports    []Port `json:"ports"`
	Finished bool   `json:"finished"`
}

// Port stores description of the port for returning with the result
type Port struct {
	Port        int    `json:"port"`
	Description string `json:"description"`
}

// Marshal the GetOutput to json byte slice
func (out *GetOutput) Marshal() ([]byte, error) {
	return json.Marshal(out)
}

// UnMarshalGetOutput unmarshal the GetOutput from json byte slice
func UnMarshalGetOutput(b []byte) (*GetOutput, error) {
	out := new(GetOutput)
	err := json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// MakeGetOutput is the helper to make the GetOutput struct
func MakeGetOutput(ip ...string) GetOutput {
	out := GetOutput{
		Results:    make([]Result, len(ip)),
		LastUpdate: time.Now(),
	}

	for i := 0; i < len(ip); i++ {
		v := Result{}
		v.IP = ip[i]
		v.Ports = []Port{}
		v.Finished = false
		out.Results[i] = v
	}
	return out
}

// MakeResult is the helper to make the Result struct, returns with port description
func MakeResult(ip string, ports ...int) Result {
	p := make([]Port, len(ports))
	for i := 0; i < len(ports); i++ {
		port := Port{ports[i], portdes.GetPortDescription(ports[i])}
		p[i] = port
	}

	return Result{
		IP:       ip,
		Ports:    p,
		Finished: true,
	}
}

// TODO: Do some tests
// UnMarshalGetInput unmarshal the GetInput from json byte slice
func UnMarshalGetInput(b []byte) (*GetInput, error) {
	out := new(GetInput)
	err := json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Marshal the GetInput to json byte slice
func (out *GetInput) Marshal() ([]byte, error) {
	return json.Marshal(out)
}