# gonverter [![Build Status](https://travis-ci.org/fern4lvarez/gonverter.png)](https://travis-ci.org/fern4lvarez/gonverter) [Documentation online](http://godoc.org/github.com/fern4lvarez/gonverter)
=====
Easy, straightforward type conversion for Go.

**gonverter** package provides a homogeneous conversion interface for
Go basic types, based on a type names initials pattern:
```
´int´ -> I
´string´ -> S
´bool´ -> B
...
```

This pattern makes possible to achieve conversions easily:
```
´int´ to ´string´: ItoS
´bool´ to ´int´: BtoS
...
```

**Disclaimer**: Go already provides simple native conversions. This package 
does not aim to replace them, just to put all of them together, making easier
to the Go developer to find the right conversion way without digging in
multiple doc sources.

**TODO**: 
```
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
```


##Install
```
go get github.com/fern4lvarez/gonverter
```

##Usage
```
// gonversions.go

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

```

```
$ go run gonversions.go
It's true, the answer is 42.
```

##License
gonverter is MIT licensed, see [here](https://github.com/fern4lvarez/gonverter/blob/master/LICENSE)