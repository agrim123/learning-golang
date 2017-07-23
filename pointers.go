/*
 *  Pointers
 */

package main

import (
	"fmt"
)

func update(val int) {
	val = val + 10
}

func updateMe(val *int) {
	*val = *val + 10
}

func main() {
	val := 10
	update(val)
	fmt.Println(val)
	updateMe(&val)
	fmt.Println(val)
}
