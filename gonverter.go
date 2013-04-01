package gonverter

import (
	"strconv"
)

// Converts a `string` into a `int`
func StoI(s string) int {
	var i int
	i, _ = strconv.Atoi(s)
	return i
}

// Converts a `int` into a `string`
func ItoS(i int) string {
	var s string
	s = strconv.Itoa(i)
	return s
}
