package main

import (
	"fmt"
)

func main() {
	list := []string{"a", "b", "c", "d", "e"}
	/*
	  We can skip the index value by using
	  for _, v := range list { }
	*/
	for i, v := range list {
		// i -> index
		// v -> value
		fmt.Println(i)
		fmt.Println(v)
		fmt.Println()
	}
}
