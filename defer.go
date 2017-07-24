/*
   Deferred function calls are pushed onto a stack.

   When a function returns, its deferred calls are executed in last-in-first-out order.

   The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

   https://blog.golang.org/defer-panic-and-recover
*/

package main

import (
	"fmt"
	"io"
	"os"
)

/*
  Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns. It's an unusual but effective way to deal with situations such as resources that must be released regardless of which path a function takes to return. The canonical examples are unlocking a mutex or closing a file.
*/
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	/*
	  Deferring a call to a function such as Close has two advantages. First, it guarantees that you will never forget to close the file, a mistake that's easy to make if you later edit the function to add a new return path. Second, it means that the close sits near the open, which is much clearer than placing it at the end of the function.
	*/
	defer f.Close() // f.Close will run when we're finished.
	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}

func main() {
	// defer fmt.Println("world")
	// fmt.Println("hello")
	// fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	// fmt.Println("done")
	defer fmt.Println("---------")
	fmt.Println(Contents("test.txt"))
}
