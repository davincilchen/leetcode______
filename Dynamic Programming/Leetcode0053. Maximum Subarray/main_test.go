package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	nums := []int{}
	v := 0

	nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	v = maxSubArray(nums)
	assert.Equal(t, 10, v)

	// nums = []int{1, -1, 2, -1, 3, -1, 1, 1, -1}
	// v = maxSubArray(nums)
	// assert.Equal(t, 4, v)

	// nums = []int{1, -1, 2, -1, 3, -1, 1,-1, 5}
	// v = maxSubArray(nums)
	// assert.Equal(t, 7, v)
}
