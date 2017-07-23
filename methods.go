/*
  Methods : A method is a function with a special receiver argument.
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Root() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Root(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	vx := Vertex{1, 2}
	fmt.Println(vx.Root())
	fmt.Println(Root(vx))
}
