package model

import (
	"encoding/json"
	"time"
)

type GetInput struct {
	token string `json:"token"`
}

type GetOutput struct {
	Results []Result     `json:"results"`
	LastUpdate time.Time `json:"last_update"`
}

type Result struct {
	IP         string    `json:"ip"`
	Ports      []Port    `json:"ports"`
	Finished   bool      `json:"finished"`
}

type Port struct {
	Port        int    `json:"port"`
	Description string `json:"description"`
}

func(out *GetOutput) Marshal() ([]byte, error) {
	return json.Marshal(out)
}

func UnMarshalGetOutput(b []byte) (*GetOutput, error) {
	out := new(GetOutput)
	err := json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

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