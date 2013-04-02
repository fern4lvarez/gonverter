package main

import (
	"fmt"
	gnv "github.com/fern4lvarez/gonverter"
)

func main() {
	var i, n int
	var s string
	var b bool

	i = gnv.StoI("38")
	// 38

	n = i + 3
	// 41

	s = gnv.ItoS(n)
	if s == "41" {
		b = gnv.StoB(s)
		// true

		i = gnv.BtoI(b)
		// 1

		res := n + i
		// 42

		b = gnv.ItoB(res)
		if b { // true
			fmt.Printf("It's %s, the answer is %s.", gnv.BtoS(b), gnv.ItoS(res))
		}
	}
}
