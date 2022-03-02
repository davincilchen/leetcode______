package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxHeapify(t *testing.T) {

	assert.Equal(t, []int{1, 3, 2, 20, 17, 21}, MinHeapify([]int{17, 20, 2, 1, 3, 21}))
	//assert.Equal(t, []int{21, 20, 17, 1, 3, 2}, MinHeapify([]int{17, 20, 2, 1, 3, 21}))
	//assert.Equal(t, []int{21, 20, 17, 1, 3, 2}, MaxHeapifyTopDown([]int{17, 20, 2, 1, 3, 21}))
}
