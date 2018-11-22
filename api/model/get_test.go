package model

import (
	"github.com/guitarpawat/portscan/api/portdes"
	"testing"
	"time"
)

func TestGetOutput_Marshal_And_UnMarshalGetOutput(t *testing.T) {
	expected := GetOutput{
		Results: []Result{
			{
				IP: "127.0.0.1",
				Ports: []Port{
					{
						Port:        80,
						Description: portdes.GetPortDescription(80),
					},
				},
				Finished: true,
			},
		},
		LastUpdate: time.Now(),
	}

	b, err := expected.Marshal()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	result, err := UnMarshalGetOutput(b)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if expected.Results[0].IP != result.Results[0].IP {
		t.Errorf("expected ip: %s, but get: %s", expected.Results[0].IP, result.Results[0].IP)
	}

	if expected.Results[0].Finished != result.Results[0].Finished {
		t.Errorf("expected finished: %t, but get: %t", expected.Results[0].Finished, result.Results[0].Finished)
	}

	if expected.Results[0].Ports[0].Description != result.Results[0].Ports[0].Description {
		t.Errorf("expected port description: %s, but get: %s", expected.Results[0].Ports[0].Description,
			result.Results[0].Ports[0].Description)
	}

	if expected.Results[0].Ports[0].Port != result.Results[0].Ports[0].Port {
		t.Errorf("expected port: %d, but get: %d", expected.Results[0].Ports[0].Port,
			result.Results[0].Ports[0].Port)
	}

	if !expected.LastUpdate.Equal(result.LastUpdate) {
		t.Errorf("expected last update: %s, but get: %s", expected.LastUpdate, result.LastUpdate)
	}
}

func TestUnMarshalGetOutput_Error(t *testing.T) {
	_, err := UnMarshalGetOutput([]byte(""))
	if err == nil {
		t.Errorf("expected error")
	}

	_, err = UnMarshalGetOutput(nil)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestMakeGetOutput(t *testing.T) {
	out := MakeGetOutput("192.168.1.1", "127.0.0.1")

	if out.Results[0].IP != "192.168.1.1" && out.Results[0].IP != "127.0.0.1" {
		t.Errorf("expected results[0] ip to be: 192.168.1.1 or 127.0.0.1, but get: %s", out.Results[0].IP)
	}

	if out.Results[1].IP != "192.168.1.1" && out.Results[1].IP != "127.0.0.1" {
		t.Errorf("expected results[1] ip to be: 192.168.1.1 or 127.0.0.1, but get: %s", out.Results[1].IP)
	}

	if out.Results[0].Finished {
		t.Errorf("expected finished: %t, but get: %t", false, out.Results[0].Finished)
	}
}

func TestMakeResult(t *testing.T) {
	expectedIP := "127.0.0.1"
	out := MakeResult(expectedIP, 80, -1)

	if out.IP != expectedIP {
		t.Errorf("expected ip: %s, but get: %s", expectedIP, out.IP)
	}

	if !out.Finished {
		t.Errorf("expected finished: %t, but get: %t", true, out.Finished)
	}

	if out.Ports[0].Port != 80 && out.Ports[0].Port != -1 {
		t.Errorf("expected port[0] to be: -1 or 80, but get: %d", out.Ports[0].Port)
	} else {
		checkPortDescription(t, out.Ports[0].Port, out.Ports[0].Description)
	}

	if out.Ports[1].Port != 80 && out.Ports[1].Port != -1 {
		t.Errorf("expected port[1] to be: -1 or 80, but get: %d", out.Ports[1].Port)
	} else {
		checkPortDescription(t, out.Ports[1].Port, out.Ports[1].Description)
	}
}

func checkPortDescription(t *testing.T, port int, description string) {
	if description != portdes.GetPortDescription(port) {
		t.Errorf("expected description of port %d: %s, but get: %s", port,
			portdes.GetPortDescription(port), description)
	}
}

func TestGetInput_Marshal_And_UnMarshalGetOutput(t *testing.T) {
	expectedToken := "asdf"
	input := GetInput{expectedToken}
	jsonData, _ := input.Marshal()
	out, _ := UnMarshalGetInput(jsonData)
	if expectedToken != out.Token {
		t.Errorf("expected token: %s, but get: %s", expectedToken, out.Token)
	}
}