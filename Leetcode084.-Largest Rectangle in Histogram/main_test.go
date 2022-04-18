package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestRectangleArea(t *testing.T) {
	nums := []int{}
	v := 0

	nums = []int{1, 3, 2, 4, 3, 2, 1}
	v = largestRectangleArea(nums)
	assert.Equal(t, 10, v)

	nums = []int{1, 3, 0, 2, 4, 0, 3, 2, 1}
	v = largestRectangleArea(nums)
	assert.Equal(t, 4, v)

	nums = []int{2, 1, 5, 6, 2, 3}
	v = largestRectangleArea(nums)
	assert.Equal(t, 10, v)

	nums = []int{2, 4}
	v = largestRectangleArea(nums)
	assert.Equal(t, 4, v)

	nums = []int{3}
	v = largestRectangleArea(nums)
	assert.Equal(t, 3, v)

	nums = []int{0}
	v = largestRectangleArea(nums)
	assert.Equal(t, 0, v)
}
