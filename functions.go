/*
 * Functions
 */

package main

import (
	"fmt"
)

// Returning Single Value
func add(x, y int) int {
	return x + y
}

// Returning multiple Values
func swap(x, y int) (int, int) {
	return y, x
}

// Returning named value
func multiply(x, y int) (mul int) {
	mul = x * y
	/*
	 * A return statement without arguments returns the named return values. This is known as a "naked" return
	 */
	return
}

/*
  	Function values :
	  Functions are values too. They can be passed around just like other values.
	  Function values may be used as function arguments and return values.
*/
func functionValues(fn func(int, int) int) int {
	return fn(1, 3)
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(swap(1, 2))
	fmt.Println(multiply(3, 2))
	fmt.Println(functionValues(add))
}
