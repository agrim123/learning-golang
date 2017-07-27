/*
 * Basic Types
 */

package main

import (
	"fmt"
)

var (
	a      int        = 1
	b      string     = "string"
	c      bool       = false
	d      float64    = 2.0212
	e      complex128 = 12 + 12i
	MaxInt uint64     = 1 << 62
	MinInt uint64     = MaxInt >> 61
)

func main() {
	f := "This is declared inside function"
	fmt.Println(a, b, c, d, e, f, MinInt, MaxInt)

	// Rune is an alias for int32. It is an UTF-8 encoded code point.
	// You could loop over each byte (which is only equivalent to a character when strings are encoded in 8-bit ASCII, which they are not in Go!)
	str := []rune("hello")
	fmt.Println(str)
}
