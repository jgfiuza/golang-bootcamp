package main

import (
	"fmt"
	"testing"
)

func testSingleCase(a, b, want int, t *testing.T) {
	result := IntMin(a, b)
	testname := fmt.Sprintf("%d,%d", a, b)
	if result != want {
		t.Errorf("\n%s -> Got %d, want %d", testname, result, want)
	}
}

func TestMultiCaseMode2(t *testing.T) {
	var tests = []struct {
		a, b, want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, test := range tests {
		testSingleCase(test.a, test.b, test.want, t)
	}

}

func TestIntMinMultipleCases(t *testing.T) {
	var tests = []struct {
		a, b, want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d,%d", test.a, test.b)
		t.Run(testname, func(t *testing.T) {
			result := IntMin(test.a, test.b)
			if result != test.want {
				t.Errorf("\n%s -> Got %d, want %d", testname, result, test.want)
			}
		})
	}
}
