/*
gonverter package provides a homogeneous conversion interface for
Go basic types, based on a type names initials pattern:
	´int´ -> I
	´string´ -> S
	´bool´ -> B
	...

This pattern makes possible to achieve conversions easily:
	´int´ to ´string´: ItoS
	´bool´ to ´int´: BtoS
	...

Disclaimer: Go already provides simple native conversions. This package 
does not aim to replace them, just to put all of them together, making easier
to the Go developer to find the right conversion way without digging in
multiple doc sources.

TODO: 
	´int8´ -> I8
	´int16´ -> I16
	´int32´ -> I32
	´int64´ -> I64

	´uint´ -> U
	´uint8´ -> U8
	´byte´ -> B // alias for uint8
	´uint16´ -> U16
	´uint32´ -> U32
	´rune´ -> R // alias for int32
	´uint64´ -> U64
	´uintptr´ -> Uptr

	´float32´ -> F
	´float64´ -> F64

	´complex64´ -> C
	´complex128´ -> C128

*/

package gonverter

import (
	"log"
	"strconv"
)

// StoI `string` => `int`
// If the string is not convertable it returns 0
func StoI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("%s couldn't be converted, returning 0\n", s)
		return 0
	}
	return i
}

// ItoS `int` => `string`
func ItoS(i int) string {
	s := strconv.Itoa(i)
	return s
}

// ItoB `int` => `bool`
// If `int` is 0, returns false, otherwise returns true
func ItoB(i int) bool {
	return i != 0
}

// BtoI `bool` => `int`
// If b is true returns 1, if false returns 0
func BtoI(b bool) int {
	if b {
		return 1
	}
	return 0
}

// StoB `string` => `bool`
// If ´string´ is empty, return false, otherwise returns true
func StoB(s string) bool {
	return s != ""
}

// BtoS `bool` => `string`
// If `bool` is true return "true", if false returns "false"
// Note: It is not reciprocal to StoB
func BtoS(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
