# gonverter [![Build Status](https://travis-ci.org/fern4lvarez/gonverter.png)](https://travis-ci.org/fern4lvarez/gonverter) [Documentation online](http://godoc.org/github.com/fern4lvarez/gonverter)
=====
Easy, straightforward type conversion for Go.


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

    i = gnv.StoI("39")
    // i = 39

    n = i + 3
    // n = 42

    s = gnv.ItoS(n)
    if s == "42" {
        b = gnv.StoB(s)
        i = gnv.BtoI(b)
        res := n + i
        b = gnv.ItoB(res)
        if b {
            fmt.Printf("It's %s, result is %s", gnv.BtoS(b), gnv.ItoS(res))
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