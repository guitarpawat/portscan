package model

import (
	"encoding/json"
	"github.com/guitarpawat/portscan/api/portdes"
	"time"
)

// GetInput is json model for getting result of specific token id
type GetInput struct {
	token string `json:"token"`
}

// GetOutput is json model for returning result of specific token id
type GetOutput struct {
	Results []Result     `json:"results"`
	LastUpdate time.Time `json:"last_update"`
}

// Result stores the open ports of specify IP
type Result struct {
	IP         string    `json:"ip"`
	Ports      []Port    `json:"ports"`
	Finished   bool      `json:"finished"`
}

// Port stores description of the port for returning with the result
type Port struct {
	Port        int    `json:"port"`
	Description string `json:"description"`
}

// Marshal the GetOutput to json byte slice
func(out *GetOutput) Marshal() ([]byte, error) {
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
func MakeGetOutput(ip ...string) (out GetOutput) {
	out = GetOutput{
		Results: make([]Result, len(ip)),
		LastUpdate: time.Now(),
	}

	for i, v := range out.Results {
		v.IP = ip[i]
		v.Ports = []Port{}
		v.Finished = false
	}
	return
}

// MakeResult is the helper to make the Result struct, returns with port description
func MakeResult(ip string, ports ...int) Result {
	p := make([]Port, len(ports))
	for i:=0; i<len(ports); i++ {
		port := Port{ports[i], portdes.GetPortDescription(ports[i])}
		p = append(p, port)
	}

	return Result{
		IP: ip,
		Ports: p,
		Finished: true,
	}
}