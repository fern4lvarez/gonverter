package gonverter

import (
	"strconv"
)

// Converts a `string` into a `int`, error handler
func StoIE(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}

// Converts a `string` into a `int`
func StoI(s string) int {
	i, _ := StoIE(s)
	return i
}

// Converts a `int` into a `string`
func ItoS(i int) string {
	s := strconv.Itoa(i)
	return s
}
