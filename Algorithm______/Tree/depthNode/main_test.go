package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxDepth(t *testing.T) {

	assert.Equal(t, []int{}, DepthNode([]int{}, 1))
	assert.Equal(t, []int{}, DepthNode([]int{3, 9, 20, -1, -1, 15, 7}, 0))
	assert.Equal(t, []int{3}, DepthNode([]int{3, 9, 20, -1, -1, 15, 7}, 1))
	assert.Equal(t, []int{9, 20}, DepthNode([]int{3, 9, 20, -1, -1, 15, 7}, 2))
	assert.Equal(t, []int{15, 7}, DepthNode([]int{3, 9, 20, -1, -1, 15, 7}, 3))
	assert.Equal(t, []int{}, DepthNode([]int{3, 9, 20, -1, -1, 15, 7}, 5))

}
