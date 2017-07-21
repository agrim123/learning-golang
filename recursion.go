/*
 * Recursion
 */

package main

import (
	"fmt"
)

func fact(x float64) float64 {
	if x == 0 {
		return 1
	}
	return x * fact(x-1)
}

func main() {
	fmt.Println(fact(13))
}
