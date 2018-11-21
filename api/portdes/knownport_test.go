package portdes

import "testing"

func TestGetPortDescription(t *testing.T) {
	r1 := GetPortDescription(80)
	r2 := GetPortDescription(-1)

	if r1 != knownPorts[80] {
		t.Errorf("expected: %s, but get: %s", knownPorts[80], r1)
	}

	if r2 != unknown {
		t.Errorf("expected: %s, but get: %s", unknown, r2)
	}
}
