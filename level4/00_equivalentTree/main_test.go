package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tour/tree"
)

func TestTreeWalking(t *testing.T) {
	// given
	ch := make(chan int)
	tree := tree.New(1)

	// when
	go Walk(tree, ch)
	result := []int{}
	for i := range ch {
		result = append(result, i)
	}

	// then
	assert.Equal(t, len(result), 10)
}

func TestSame(t *testing.T) {
	// given
	t1 := tree.New(1)
	t2 := tree.New(1)

	// when
	result := Same(t1, t2)

	// then
	assert.True(t, result)
}

func TestSameNot(t *testing.T) {
	// given
	t1 := tree.New(1)
	t2 := tree.New(2)

	// when
	result := Same(t1, t2)

	// then
	assert.False(t, result)
}
