/*
 * Struct
 */

package main

import (
	"fmt"
)

type Vertex struct {
	X float64
	Y float64
}

func main() {
	v := Vertex{1, 2}
	v.X = 5
	fmt.Println(v.X)
}
