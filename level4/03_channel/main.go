// Statement
// Create a channel of int and then create a goroutine to add a value to
// the channel and then print the channel value in the main function

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 1
	}(ch)

	fmt.Printf("%v\n", <-ch)
}
