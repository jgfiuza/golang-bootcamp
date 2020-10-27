// Given a list of strings find all strings that contain a given substring. You should split the
// list into chunks and process them in parallel. Each goroutine should write the response to an output
// channel which will be consumed in the end to consolidate the response.

package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	strs := []string{"abc", "bcd", "efg", "aabcr", "acb", "ggg", "hjuklbc"}
	substr := "bc"

	ch := make(chan string)
	waitingGroup := sync.WaitGroup{}
	defer close(ch)

	for _, oneString := range strs {
		waitingGroup.Add(1)
		go func(str string) {
			if strings.Contains(str, substr) {
				ch <- str
			}
			waitingGroup.Done()
		}(oneString)
	}

	go func() {
		for str := range ch {
			fmt.Println(str)
		}
	}()

	waitingGroup.Wait()
}
