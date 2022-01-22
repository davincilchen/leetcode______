package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MainFun(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	assert.Equal(t, []int{1, 2}, twoSum(nums, target))

	nums = []int{2, 3, 4}
	target = 6

	assert.Equal(t, []int{1, 3}, twoSum(nums, target))

	nums = []int{3, 3}
	target = 6
	assert.Equal(t, []int{1, 2}, twoSum(nums, target))

	nums = []int{1, 2, 3, 3, 7}
	target = 6
	assert.Equal(t, []int{3, 4}, twoSum(nums, target))

}
