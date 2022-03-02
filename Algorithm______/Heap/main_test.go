package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxHeapify(t *testing.T) {

	assert.Equal(t, []int{1, 3, 2, 20, 17, 21}, HeapifyMin([]int{17, 20, 2, 1, 3, 21}))
	assert.Equal(t, []int{1, 3, 2, 6, 20, 5, 4, 17, 21}, HeapifyMin([]int{5, 6, 4, 17, 20, 2, 1, 3, 21}))

	assert.Equal(t, []int{21, 20, 17, 1, 3, 2}, HeapifyMax([]int{17, 20, 2, 1, 3, 21}))
	assert.Equal(t, []int{21, 20, 4, 17, 5, 2, 1, 3, 6}, HeapifyMax([]int{5, 6, 4, 17, 20, 2, 1, 3, 21}))
}
