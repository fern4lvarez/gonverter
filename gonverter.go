package gonverter

import (
	"strconv"
)

// Converts a `string` into a `int`
func StoI(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Converts a `int` into a `string`
func ItoS(i int) string {
	s := strconv.Itoa(i)
	return s
}

// Converts a `int` into a `bool`
// If i is 0, return false, otherwise returns true
func ItoB(i int) bool {
	return i != 0
}

// Converts a `bool` into a `int`
// If b is true returns 1, if false returns 0
func BtoI(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Converts a `string` into a `bool`
// If s is empty, return false, otherwise returns true
func StoB(s string) bool {
	return s != ""
}

// Converts a `bool` into a `string`
// If b is true return "true", if false returns "false"
// Note: It is not reciprocal to StoB
func BtoS(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
