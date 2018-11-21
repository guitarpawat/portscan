package model

// ScanInput is the json model for port scanning request
type ScanInput struct {
	Target string `json:"target"`
}

// ScanInput is the json model for returning port scanning request with token id
type ScanOutput struct {
	Token string   `json:"token"`
	IP    []string `json:"ip"`
}
