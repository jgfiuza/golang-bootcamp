package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterByAgeRange(t *testing.T) {
	// given
	fromAge := 0
	toAge := 100
	agesSlice := []int{-1, 0, 1, 99, 100, 101}

	// when
	result := filterByAgeRange(fromAge, toAge, agesSlice)

	// then
	expected := []int{0, 1, 99, 100}
	assert.Equal(t, len(expected), len(result))
	assert.True(t, reflect.DeepEqual(expected, result))
}

func TestFilterByAgeRangeTable(t *testing.T) {
	// given
	ageSlice := []int{0, 5, 10, 15, 20, 25, 30}

	tests := []struct {
		fromAge int
		toAge   int
		want    []int
	}{
		{10, 15, []int{10, 15}},
		{5, 25, []int{5, 10, 15, 20, 25}},
		{0, 30, []int{0, 5, 10, 15, 20, 25, 30}},
		{11, 12, []int{}},
	}

	// when
	for idx, tt := range tests {
		testname := fmt.Sprintf("%v: from %v - to %v", idx, tt.fromAge, tt.toAge)
		t.Run(testname, func(t *testing.T) {
			result := filterByAgeRange(tt.fromAge, tt.toAge, ageSlice)
			assert.Equal(t, len(tt.want), len(result))
			assert.True(t, reflect.DeepEqual(tt.want, result))
		})
	}
}
