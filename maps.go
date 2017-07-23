/*
 * Maps :A map maps keys to values.
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
	m := make(map[string]Vertex)
	fmt.Println(m)
	m["Yoo"] = Vertex{1, 2}
	fmt.Println(m)

	//  Map literals are like struct literals, but the keys are required.
	var m1 = map[string]Vertex{
		"Bell Labs": {4123.12, -714.39967},
		"Google":    {3127.41232, -3511.08408},
	}
	delete(m1, "Google")
	fmt.Println(m1)
}
