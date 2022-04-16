package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxDepth(t *testing.T) {

	assert.Equal(t, true, IsBalanced([]int{3, 9, 20, -1, -1, 15, 7}))
	assert.Equal(t, false, IsBalanced([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}))

	assert.Equal(t, true, IsBalanced([]int{}))

	assert.Equal(t, true, IsBalanced2([]int{3, 9, 20, -1, -1, 15, 7}))
	assert.Equal(t, false, IsBalanced2([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}))

	assert.Equal(t, true, IsBalanced2([]int{}))
}
