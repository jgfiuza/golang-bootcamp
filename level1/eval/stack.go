package main

// Create a Stack structure that means a LIFO (Last In First Out).
// Note: The func main should be provided and Just need to create the Stack struct and
// Methods to Push, Pop and Peek. IsEmpty is optional.
// To test it, push two or three strings and then pop the last one and finally, push another one.

// Stack implements a string stack
type Stack struct {
	nodes     []node
	lastIndex int
}

type node struct {
	value string
}

// Push a new element into the Stack.
func (s *Stack) Push(newString string) {
	if newString == "" {
		return
	}
	newNode := node{newString}
	if s.lastIndex >= cap(s.nodes) {
		s.nodes = append(s.nodes, newNode)
	} else {
		s.nodes[s.lastIndex] = newNode
	}

	s.lastIndex++
}

// Pop the last element and remove it from the Stack.
func (s *Stack) Pop() string {
	var result string
	if s.lastIndex > 0 {
		s.lastIndex--
		result, s.nodes[s.lastIndex].value = s.nodes[s.lastIndex].value, ""
	}
	return result
}

// Peek the last element and keep it in the Stack.
func (s Stack) Peek() string {
	if len(s.nodes) > 0 {
		return s.nodes[s.lastIndex-1].value
	}
	return ""
}
