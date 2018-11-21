package model

// ScanInput is the json model for port scanning request
type ScanInput struct {
	Target string `json:"target"`
}

// ScanOutput is the json model for returning port scanning request with Token id
type ScanOutput struct {
	Token string   `json:"Token"`
	IP    []string `json:"ip"`
}
