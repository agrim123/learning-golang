/*
 * Flow
 */

package main

import (
	"fmt"
)

func main() {
	// For / While loops
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("--------")

	sum1 := 1
	// While loop
	for sum1 <= 10 { // forloop => for ;sum1 <=10;
		sum1 += sum1
	}
	fmt.Println(sum1)
	fmt.Println("--------")

	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("--------")

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
