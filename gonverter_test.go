package gonverter

import "testing"

var msgFail = "%s function fails. Expects %s, returns %s."

// Test int to string conversion
func TestItoS(t *testing.T) {
	var s string
	if s = ItoS(42); s != "42" {
		t.Errorf(msgFail, "ItoS", "42", s)
	}

	if s = ItoS(-5); s != "-5" {
		t.Errorf(msgFail, "ItoS", "-5", s)
	}
}

// Test string to int conversion
func TestStoI(t *testing.T) {
	var i int
	if i = StoI("42"); i != 42 {
		t.Errorf(msgFail, "StoI", ItoS(42), ItoS(i))
	}

	if i = StoI("foo"); i != 0 {
		t.Errorf(msgFail, "StoI", ItoS(0), ItoS(i))
	}
}

// Test int to bool conversion
func TestItoB(t *testing.T) {
	var b bool
	if b = ItoB(42); !b {
		t.Errorf(msgFail, "ItoB", "true", BtoS(b))
	}

	if b = ItoB(0); b {
		t.Errorf(msgFail, "ItoB", "false", BtoS(b))
	}
}

//Test bool to int conversion
func TestBtoI(t *testing.T) {
	var i int
	if i = BtoI(true); i != 1 {
		t.Errorf(msgFail, "BtoI", "1", ItoS(i))
	}

	if i = BtoI(false); i != 0 {
		t.Errorf(msgFail, "BtoI", "0", ItoS(i))
	}
}

// Test string to bool conversion
func TestStoB(t *testing.T) {
	var b bool
	if b = StoB("Hello"); !b {
		t.Errorf(msgFail, "StoB", "false", BtoS(b))
	}

	if b = StoB(""); b {
		t.Errorf(msgFail, "StoB", "false", BtoS(b))
	}
}

//Test bool to string conversion
func TestBtoS(t *testing.T) {
	var s string
	if s = BtoS(true); s != "true" {
		t.Errorf(msgFail, "BtoS", "true", s)
	}

	if s = BtoS(false); s != "false" {
		t.Errorf(msgFail, "BtoS", "false", s)
	}
}
