// Statement
// Create a function that will increase a number and that function will
// be executed by a goroutine inside a for loop (x1000 times). To avoid
// race conditioning, implement the sync.Mutex and Lock and Unlock
// inside the increase() function.
// Note: Add a time.Sleep() to be able to see the final n

package main

import "sync"

type number struct {
	value int
	mux   *sync.Mutex
}

func (n *number) increase(ch chan int) {
	n.mux.Lock()
	defer n.mux.Unlock()
	ch <- n.value
}

func main() {

}
