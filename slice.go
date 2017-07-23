/*
 * Slice
 */

package main

import (
	"fmt"
	"strings"
)

// Nil slices
func nilSlices() {
	var nilSlice []int
	fmt.Println(nilSlice)
}

func sampleSlices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var slice []int = primes[1:4]
	fmt.Println(slice)
}

/*
	Slice Literals
	- A slice literal is like an array literal without the length.
*/
func sliceLiterals() {
	l := []int{1, 2, 3}
	fmt.Println(l)
	sl := []struct {
		i int
		b bool
	}{
		{4, true},
		{1, false},
		{12, true},
		{31, true},
		{17, false},
		{13, true},
	}
	fmt.Printf("len=%d cap=%d %v\n", len(sl), cap(sl), sl)
	fmt.Println(sl[1:])
}

/*
  Slices can be created with the built-in make function
  The make function allocates a zeroed array and returns a slice that refers to that array.
*/
func sliceWithMake() {
	a := make([]int, 5)
	a[2] = 1
	fmt.Println(a)
}

// Slices can contain any type, including other slices.
func slicesOfSlices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "0"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

/*
  Go has builtin append function for slices
  func append(s []T, vs ...T) []T
*/
func appendingToSlice() {
	var s []int
	// We can pass as many elements we want to add to slice
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func main() {
	nilSlices()
	sampleSlices()
	sliceLiterals()
	sliceWithMake()
	slicesOfSlices()
	appendingToSlice()
}
