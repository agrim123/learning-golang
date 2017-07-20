/*
 * Arrays
 */

package main

import (
	"fmt"
)

func sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var twoD [2][3][4]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				twoD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println(twoD)
	s := make([]string, 3)
	fmt.Println("emp:", s)
	array := [...]float64{1.2, 4.5, 9.0}
	fmt.Println(sum(&array))
}
