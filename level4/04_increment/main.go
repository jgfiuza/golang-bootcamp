// Statement
// Create a function that will increase a number and that function will
// be executed by a goroutine inside a for loop (x1000 times). To avoid
// race conditioning, implement the sync.Mutex and Lock and Unlock
// inside the increase() function.
// Note: Add a time.Sleep() to be able to see the final n

package main

import (
	"fmt"
	"sync"
	"time"
)

func increase(num *int, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	*num++
}

func main() {
	var number int
	mutex := &sync.Mutex{}

	for i := 1; i <= 1000; i++ {
		go increase(&number, mutex)
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("%v\n", number)

}
