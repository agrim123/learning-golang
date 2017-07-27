/*

Interfaces

- An interface is two things: it is a set of methods, but it is also a type.
- The interface{} type, the empty interface, is the source of much confusion. The interface{} type is the interface that has no methods. Since there is no implements keyword, all types implement at least zero methods, and satisfying an interface is done automatically, all types satisfy the empty interface. That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value.
	func DoSomething(v interface{}) {
   		// ...
	}
	will accept any parameter whatsoever.

*/

package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d *Dog) Speak() string {
	return "Woof!"
}

func main() {
	animals := []Animal{new(Dog)}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
