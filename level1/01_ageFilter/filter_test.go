package main

import "testing"

func TestEmptyArray(t *testing.T) {
	// given
	min := 1
	max := 1
	array := []int{}

	// when
	result := filter(array, min, max)

	// then
	if len(result) != 0 {
		t.Errorf("Result: %v - Expected: %v", len(result), result)
	}
}

func TestFilter(t *testing.T) {
	// given
	min := 10
	max := 20
	array := []int{5, 10, 15, 20, 25}

	//when
	result := filter(array, min, max)

	//then
	expected := 3
	if len(result) != expected {
		t.Errorf("Result: %v - Expected: %v", len(result), expected)
	}
}

func TestFilterInverseBoundaries(t *testing.T) {
	// given
	min := 20
	max := 10
	array := []int{5, 10, 11, 12, 15, 20, 25}

	//when
	result := filter(array, min, max)

	//then
	expected := 5
	if len(result) != expected {
		t.Errorf("Result: %v - Expected: %v", len(result), expected)
	}
}
