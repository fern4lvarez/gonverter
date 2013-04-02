package gonverter

import "testing"

// Test int to string conversion
func TestItoS(t *testing.T) {
	if ItoS(42) != "42" {
		t.Error("ItoS function didn't work as expected")
	}
}

// Test string to int conversion
func TestStoIE(t *testing.T) {
	i, err := StoIE("42")
	if i != 42 || err != nil {
		t.Error("StoI function didn't work as expected")
	}
}

// Test string to int conversion with errors
func TestStoIE_Error(t *testing.T) {
	i, err := StoIE("44abs")
	msg := "StoI function didn't work as expected"

	if err == nil {
		t.Error(msg)
	}
	if i == 44 {
		t.Error(msg)
	}
}

// Test string to int conversion
func TestStoI(t *testing.T) {
	i := StoI("42")
	if i != 42 {
		t.Error("StoI function didn't work as expected")
	}
}
