package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxDepth(t *testing.T) {

	assert.Equal(t, 3, MaxDepth([]int{3, 9, 20, -1, -1, 15, 7}))
	assert.Equal(t, 2, MaxDepth([]int{1, -1, 2}))

	assert.Equal(t, 3, MaxDepth2([]int{3, 9, 20, -1, -1, 15, 7}))
	assert.Equal(t, 2, MaxDepth2([]int{1, -1, 2}))

}
