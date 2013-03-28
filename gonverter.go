package gonverter

import (
  "strconv"
)

func Sti(s string) int {
  var i int
  i, _ = strconv.Atoi(s)
  return i
}

func Its(i int) string {
  var s string
  s = strconv.Itoa(i)
  return s
}
