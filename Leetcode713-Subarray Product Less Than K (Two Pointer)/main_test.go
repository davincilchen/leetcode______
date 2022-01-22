package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MainFun(t *testing.T) {

	nums := []int{10, 5, 2, 6}
	assert.Equal(t, 8, numSubarrayProductLessThanK(nums, 100))

	nums = []int{1, 2, 3}
	assert.Equal(t, 0, numSubarrayProductLessThanK(nums, 1))

	nums = []int{1, 2, 3}
	assert.Equal(t, 0, numSubarrayProductLessThanK(nums, 0))

	nums = []int{4, 5, 6}
	assert.Equal(t, 0, numSubarrayProductLessThanK(nums, 3))

	// Constraints:

	// 1 <= nums.length <= 3 * 104
	// 1 <= nums[i] <= 1000
	// 0 <= k <= 106

	// nums = []int{0, 1, 2, 3}
	// assert.Equal(t, 0, numSubarrayProductLessThanK(nums, 1))
}
