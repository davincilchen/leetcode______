package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pow(t *testing.T) {

	v := Pow(2, 10)
	assert.Equal(t, 1024, v)

	v = Pow(10, 2)
	assert.Equal(t, 100, v)

	v = Pow(3, 4)
	assert.Equal(t, 81, v)
}

func Test_TransToNormalTree(t *testing.T) {
	v := []int{}
	nums := []int{3, 9, 20, -1, -1, 15, 7} //3, 9, 20, null, null, 15, 7

	nums = []int{2, -1, 3, -1, 4, -1, 5, -1, 6}
	v = TransToNormalTree(nums)
	assert.Equal(t, []int{2, -1, 3, -1, -1, -1, 4,
		-1, -1, -1, -1, -1, -1, -1, 5,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 6}, v)

	nums = []int{2, 3, 4, -1, 5, -1, 6, -1, 7, -1, 8}
	v = TransToNormalTree(nums)
	assert.Equal(t, []int{2, 3, 4, -1, 5, -1, 6,
		-1, -1, -1, 7, -1, -1, -1, 8}, v)

}

// func Test_minDepth(t *testing.T) {

// 	nums := []int{}
// 	v := 0

// 	nums = []int{3, 9, 20, -1, -1, 15, 7} //3, 9, 20, null, null, 15, 7
// 	v = minDepth(buildTree(nums))
// 	assert.Equal(t, 2, v)

// 	nums = []int{1, 3, 0, 2, 4, 0, 3, 2, 1}
// 	v = minDepth(buildTree(nums))
// 	assert.Equal(t, 3, v)

// 	nums = []int{2, -1, 3, -1, 4, -1, 5, -1, 6} //2,null,3,null,4,null,5,null,6
// 	v = minDepth(buildTree(nums))
// 	assert.Equal(t, 5, v)

// 	nums = []int{2, -1, 3, -1, 4, -1, 5, -1, 6, -1, 7, -1, 8} //2,null,3,null,4,null,5,null,6
// 	v = minDepth(buildTree(nums))
// 	assert.Equal(t, 7, v)
// }
