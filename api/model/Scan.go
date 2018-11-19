package model

type ScanInput struct {
	Target string `json:"target"`
}

type ScanOutput struct {
	Token string   `json:"token"`
	IP    []string `json:"ip"`
}
