package scanner

import (
	"net"
	"net/http"
	"testing"
)

func TestNSLookUp(t *testing.T) {
	expectedv4 := "127.0.0.1"
	v4, err := NSLookUp("localhost")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if v4 != expectedv4 {
		t.Errorf("expected ipv4: %s, but get %s", expectedv4, v4)
	}
}

func TestNSLookUp_Malformed(t *testing.T) {
	_, err := NSLookUp("some random string")
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGetOpenPorts(t *testing.T) {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	http.HandleFunc("/", h1)
	listener1, _ := net.Listen("tcp", ":0")
	defer listener1.Close()
	listener2, _ := net.Listen("tcp", ":0")
	listener2.Close()
	port1 := listener1.Addr().(*net.TCPAddr).Port
	port2 := listener2.Addr().(*net.TCPAddr).Port
	open := GetOpenPorts("localhost", nil, port1, port2)

	if len(open) != 1 {
		t.Fatalf("expected slice of open ports length: %d, but get: %d", 1, len(open))
	}
	if open[0] != port1 {
		t.Errorf("expected open port: %d, but get %d", port1, open[0])
	}
}

func TestIsV4(t *testing.T) {
	ip := "127.0.0.1"
	r := IsV4(ip)
	if !r {
		t.Errorf("expected IsV4 for %s to be %t", ip, true)
	}

	ip = "127.0.0.1.5"
	r = IsV4(ip)
	if r {
		t.Errorf("expected IsV4 for %s to be %t", ip, false)
	}
}

func TestIsV6(t *testing.T) {
	ip := "::1"
	r := IsV6(ip)
	if !r {
		t.Errorf("expected IsV6 for %s to be %t", ip, true)
	}

	ip = "g05::1"
	r = IsV6(ip)
	if r {
		t.Errorf("expected IsV6 for %s to be %t", ip, false)
	}
}
