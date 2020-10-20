package main

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	// given
	message := "Test1"
	stack := Stack{}

	// when
	stack.Push(message)

	// then
	if len(stack.nodes) != 1 {
		t.Errorf("Test failed")
	}
}

func TestStackPeek(t *testing.T) {
	// given
	message := "Test1"
	stack := Stack{}
	// when
	stack.Push(message)
	result := stack.Peek()
	// then
	if len(stack.nodes) != 1 || result != message {
		t.Errorf("Test failed")
	}
}

func TestStackPop(t *testing.T) {
	// given
	message1 := "Test1"
	message2 := "Test2"
	stack := Stack{}
	stack.Push(message1)
	stack.Push(message2)

	// when
	result := stack.Pop()

	// then
	if result != message2 || stack.nodes[0].value != message1 || stack.nodes[1].value != "" {
		t.Errorf("Test failed")
	}
}
