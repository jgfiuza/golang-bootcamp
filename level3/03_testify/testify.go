package testify

import "errors"

// IntMax returns the max number of two integers
func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// IntMaxMulti returns the max number of several integers
func IntMaxMulti(numbers ...int) (int, error) {
	if len(numbers) == 0 {
		return -1, errors.New("Empty input")
	}
	var max int
	for idx, number := range numbers {
		if idx == 0 {
			max = number
		}
		if number > max {
			max = number
		}
	}
	return max, nil
}
