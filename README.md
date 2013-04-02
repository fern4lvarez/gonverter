# gonverter [![Build Status](https://travis-ci.org/fern4lvarez/gonverter.png)](https://travis-ci.org/fern4lvarez/gonbverter) [Documentation online](http://godoc.org/github.com/fern4lvarez/gonverter)
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
  var i, n, res int
  var s string
  var b bool

  i = gnv.StoI("38")
  // i = 38

  n = i + 3
  // n = 41

  s = gnv.ItoS(n)
  if s == "42" {
    b = gnv.StoB(s)
    // b is true

    i = gnv.BtoI(b)
    // i is 1

    res = n + i
    // res is 42

    b = gnv.ItoB(res)
    if b // true {
      fmt.Printf("It's %s, the answer is %s.", gnv.BtoS(b), gnv.ItoS(res))
    }
  }

}
```

```
$ go run gonversions.go
It's true, the answer is 42.

```

###License
gonverter is MIT licensed, see [here](https://github.com/fern4lvarez/gonverter/blob/master/LICENSE)