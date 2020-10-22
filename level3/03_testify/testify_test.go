package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	// given
	a := 1
	b := -1

	// when
	result := IntMax(a, b)

	// then
	assert.Equal(t, a, result, "Should be equal")
}

func TestIntMaxMultiEmpty(t *testing.T) {
	// given
	// nothing

	// when
	_, err := IntMaxMulti()

	// then
	assert.NotNil(t, err)
}

func TestIntMaxMultiArray(t *testing.T) {
	// given
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// when
	result, err := IntMaxMulti(numbers...)

	// then
	expected := 8
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestIntMaxMultiIsoArray(t *testing.T) {
	// given
	numbers := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	// when
	result, _ := IntMaxMulti(numbers...)

	// then
	expected := 1
	assert.Equal(t, expected, result)
}
