package model

import "time"

type GetInput struct {
	token string `json:"token"`
}

type GetOutput struct {
	results []Result `json:"results"`
}

type Result struct {
	IP         string    `json:"ip"`
	Results    []Result  `json:"results"`
	Finished   bool      `json:"finished"`
	LastUpdate time.Time `json:"last_update"`
}

type Port struct {
	Port        int    `json:"port"`
	Description string `json:"description"`
}
