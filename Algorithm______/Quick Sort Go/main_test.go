package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sort(t *testing.T) {

	//no compare
	assert.Equal(t, []int{5, 6, 4, 17, 20, 2, 1, 3, 21}, Sort([]int{5, 6, 4, 17, 20, 2, 1, 3, 21}))

	// sort // bigger//遞增 increase
	Compare = Bigger
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 17, 20, 21}, Sort([]int{5, 6, 4, 17, 20, 2, 1, 3, 21}))

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 17, 20, 21}, Sort([]int{21, 20, 17, 6, 5, 4, 3, 2, 1}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 17, 20, 21}, Sort([]int{1, 2, 3, 4, 5, 6, 17, 20, 21}))
	assert.Equal(t, []int{4, 5, 6}, Sort([]int{5, 6, 4}))
	assert.Equal(t, []int{5}, Sort([]int{5}))

	assert.Equal(t, []int{1, 2, 3, 3, 3, 4, 5, 6, 6, 6, 17, 20, 21}, Sort([]int{5, 3, 6, 4, 17, 6, 3, 20, 2, 1, 6, 3, 21}))

	// sort // smaller//遞減 decrease
	Compare = Smaller
	assert.Equal(t, []int{21, 20, 17, 6, 5, 4, 3, 2, 1}, Sort([]int{5, 6, 4, 17, 20, 2, 1, 3, 21}))

	assert.Equal(t, []int{21, 20, 17, 6, 5, 4, 3, 2, 1}, Sort([]int{21, 20, 17, 6, 5, 4, 3, 2, 1}))
	assert.Equal(t, []int{21, 20, 17, 6, 5, 4, 3, 2, 1}, Sort([]int{1, 2, 3, 4, 5, 6, 17, 20, 21}))
	assert.Equal(t, []int{6, 5, 4}, Sort([]int{5, 6, 4}))
	assert.Equal(t, []int{5}, Sort([]int{5}))
}
