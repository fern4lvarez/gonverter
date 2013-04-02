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
	msg := "StoI function didn't work as expected"
	// OK
	i, err := StoIE("42")
	if i != 42 || err != nil {
		t.Error(msg)
	}
	// Wrong
	i, err = StoIE("44abs")
	if err == nil || i == 44 {
		t.Error(msg)
	}
}

// Test string to int conversion
func TestStoI(t *testing.T) {
	i1 := StoI("42")
	i2, _ := StoIE("42")
	if i1 != i2 {
		t.Error("StoI function didn't work as expected")
	}
}

// Test int to bool conversion
func TestItoB(t *testing.T) {
	b1 := ItoB(42)
	b2 := ItoB(0)
	if !b1 || b2 {
		t.Error("ItoB function didn't work as expected")
	}
}

//Test bool to int conversion
func TestBtoI(t *testing.T) {
	i1 := BtoI(true)
	i2 := BtoI(false)
	if i1 != 1 || i2 != 0 {
		t.Error("BtoI function didn't work as expected")
	}
}

// Test string to bool conversion
func TestStoB(t *testing.T) {
	s1 := StoB("Hello")
	s2 := StoB("")
	if !s1 || s2 {
		t.Error("StoB function didn't work as expected")
	}
}

//Test bool to string conversion
func TestBtoS(t *testing.T) {
	s1 := BtoS(true)
	s2 := BtoS(false)
	if s1 != "true" || s2 != "false" {
		t.Error("BtoS function didn't work as expected")
	}
}
