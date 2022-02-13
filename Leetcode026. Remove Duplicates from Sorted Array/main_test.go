package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestRectangleArea(t *testing.T) {
	nums := []int{}
	v := 0

	nums = []int{1, 1, 2}
	v = removeDuplicates(nums)
	assert.Equal(t, 2, v)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	v = removeDuplicates(nums)
	assert.Equal(t, 5, v)

}
