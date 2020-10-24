// Statement
// Create a goroutine that will execute an anonymous function to print
// “Hello World” and in the main routine print “main function”

package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("Hello World!")
	}()
	// To print we have to make the main func take the time to allow the go routine to execute
	// time.Sleep(time.Second)
	fmt.Println("main function")
}
