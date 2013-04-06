package gonverter

import "testing"

var msgFail = "%s function fails. Expects %s, returns %s."

// Test int to string conversion
func TestItoS(t *testing.T) {
	if s := ItoS(42); s != "42" {
		t.Errorf(msgFail, "ItoS", "42", s)
	}

	if s := ItoS(-5); s != "-5" {
		t.Errorf(msgFail, "ItoS", "-5", s)
	}
}

// Test string to int conversion
func TestStoI(t *testing.T) {
	var i int
	i = StoI("42")
	if i != 42 {
		t.Errorf(msgFail, "StoI", ItoS(42), ItoS(i))
	}

	i = StoI("foo")
	if i != 0 {
		t.Errorf(msgFail, "StoI", ItoS(0), ItoS(i))
	}
}

// Test int to bool conversion
func TestItoB(t *testing.T) {
	var b bool
	b = ItoB(42)
	if !b {
		t.Errorf(msgFail, "ItoB", "true", BtoS(b))
	}

	b = ItoB(0)
	if b {
		t.Errorf(msgFail, "ItoB", "false", BtoS(b))
	}
}

//Test bool to int conversion
func TestBtoI(t *testing.T) {
	var i int
	i = BtoI(true)
	if i != 1 {
		t.Errorf(msgFail, "BtoI", "1", ItoS(i))
	}

	i = BtoI(false)
	if i != 0 {
		t.Errorf(msgFail, "BtoI", "0", ItoS(i))
	}
}

// Test string to bool conversion
func TestStoB(t *testing.T) {
	var b bool
	b = StoB("Hello")
	if !b {
		t.Errorf(msgFail, "StoB", "false", BtoS(b))
	}

	b = StoB("")
	if b {
		t.Errorf(msgFail, "StoB", "false", BtoS(b))
	}

}

//Test bool to string conversion
func TestBtoS(t *testing.T) {
	var s string
	s = BtoS(true)
	if s != "true" {
		t.Errorf(msgFail, "BtoS", "true", s)
	}

	s = BtoS(false)
	if s != "false" {
		t.Errorf(msgFail, "BtoS", "false", s)
	}

}
