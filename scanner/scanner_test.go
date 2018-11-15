package scanner

import "testing"

func TestNSLookUp(t *testing.T) {
	expectedv4 := "127.0.0.1"
	expectedv6 := "::1"
	v4, v6, err := NSLookUp("localhost")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if v4 != expectedv4 {
		t.Errorf("expected ipv4: %s, but get %s", expectedv4, v4)
	}
	if v6 != expectedv6 {
		t.Errorf("expected ipv4: %s, but get %s", expectedv6, v6)
	}
}
