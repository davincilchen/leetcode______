package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MainFun(t *testing.T) {

	nums := []int{-4, -1, 0, 3, 10}
	assert.Equal(t, []int{0, 1, 9, 16, 100}, SortedSquares1(nums))

	nums = []int{-7, -3, 2, 3, 11}
	assert.Equal(t, []int{4, 9, 9, 49, 121}, SortedSquares1(nums))

	// == //
	nums = []int{-4, -1, 0, 3, 10}
	assert.Equal(t, []int{0, 1, 9, 16, 100}, SortedSquares2(nums))

	nums = []int{-7, -3, 2, 3, 11}
	assert.Equal(t, []int{4, 9, 9, 49, 121}, SortedSquares2(nums))

	nums = []int{-11, -7, -3, -3, -2}
	assert.Equal(t, []int{4, 9, 9, 49, 121}, SortedSquares2(nums))

	nums = []int{-1, 1}
	assert.Equal(t, []int{-1, 1}, SortedSquares2(nums))

	// .. //
	nums = []int{-4, -1, 0, 3, 10}
	assert.Equal(t, []int{0, 1, 9, 16, 100}, SortedSquares3(nums))

	nums = []int{-7, -3, 2, 3, 11}
	assert.Equal(t, []int{4, 9, 9, 49, 121}, SortedSquares3(nums))

	nums = []int{-11, -7, -3, -3, -2}
	assert.Equal(t, []int{4, 9, 9, 49, 121}, SortedSquares3(nums))

	nums = []int{-1, 1}
	assert.Equal(t, []int{-1, 1}, SortedSquares3(nums))

	// .. //
	nums = []int{-4, -1, 0, 3, 10}
	assert.Equal(t, []int{0, 1, 9, 16, 100}, SortedSquares4(nums))

	nums = []int{-7, -3, 2, 3, 11}
	assert.Equal(t, []int{4, 9, 9, 49, 121}, SortedSquares4(nums))

	nums = []int{-11, -7, -3, -3, -2}
	assert.Equal(t, []int{4, 9, 9, 49, 121}, SortedSquares4(nums))

	nums = []int{-1, 1}
	assert.Equal(t, []int{-1, 1}, SortedSquares4(nums))

}
