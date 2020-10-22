package main

import (
	"golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch)
	ch <- t.Value
	walkImpl(t.Right, ch)
}

// Walk collect the elements of a Tree and channel it through
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	close(ch)
}

// Same verifies that two Trees have the same elements
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	verificationMap := make(map[int]struct{})

	for num := range ch1 {
		verificationMap[num] = struct{}{}
	}

	for num := range ch2 {
		_, ok := verificationMap[num]
		if !ok {
			return false
		}
	}

	return true
}

func main() {
}
