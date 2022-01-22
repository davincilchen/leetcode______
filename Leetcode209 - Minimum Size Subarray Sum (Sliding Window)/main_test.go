package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MainFun(t *testing.T) {

	nums := []int{2, 3, 1, 2, 4, 3}
	assert.Equal(t, 2, minSubArrayLen(7, nums))

	nums = []int{1, 4, 4}
	assert.Equal(t, 1, minSubArrayLen(4, nums))

	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	assert.Equal(t, 0, minSubArrayLen(11, nums))

}
