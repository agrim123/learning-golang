/*
 * Conditionals
 */

package main

import (
	"fmt"
	"math"
	"time"
)

var (
	a int        = 1
	b string     = "asd"
	c bool       = false
	d float64    = 2.02
	e complex128 = 12 + 12i
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// v is accessible here
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here
	return lim
}

func main() {
	// If / Else
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	fmt.Println("--------")
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println("--------")
	// Switch case
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(a)
	whatAmI(c)
	fmt.Println("--------")
	// Switch without condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
