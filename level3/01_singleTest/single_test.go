package main

import "testing"

// TestIntMin ...
func TestIntMin(t *testing.T) {
	// given
	expected := 0
	b := 1

	// when
	result := IntMin(expected, b)

	// then
	if result != expected {
		t.Errorf("Case 1: Got %v, wanted %v", result, expected)
	}
}
