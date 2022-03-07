package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Heapify(t *testing.T) {

	//.. min .. //
	assert.Equal(t, []int{1, 3, 2, 20, 17, 21}, HeapifyMin([]int{17, 20, 2, 1, 3, 21}))
	assert.Equal(t, []int{1, 3, 2, 6, 20, 5, 4, 17, 21}, HeapifyMin([]int{5, 6, 4, 17, 20, 2, 1, 3, 21}))

	//.. max .. //
	assert.Equal(t, []int{21, 20, 17, 1, 3, 2}, HeapifyMax([]int{17, 20, 2, 1, 3, 21}))
	assert.Equal(t, []int{21, 20, 4, 17, 5, 2, 1, 3, 6}, HeapifyMax([]int{5, 6, 4, 17, 20, 2, 1, 3, 21}))

}

func Test_Sort(t *testing.T) {

	// sort
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 17, 20, 21}, Sort([]int{5, 6, 4, 17, 20, 2, 1, 3, 21}, false))
	assert.Equal(t, []int{4, 5, 6}, Sort([]int{5, 6, 4}, false))
	assert.Equal(t, []int{6}, Sort([]int{6}, false))

	// sort
	assert.Equal(t, []int{21, 20, 17, 6, 5, 4, 3, 2, 1}, Sort([]int{5, 6, 4, 17, 20, 2, 1, 3, 21}, true))
	assert.Equal(t, []int{6, 5, 4}, Sort([]int{5, 6, 4}, true))
	assert.Equal(t, []int{5}, Sort([]int{5}, true))
}

func Test_Push(t *testing.T) {

	h := NewHeap([]int{5, 6, 4, 17, 20, 11, 1, 3, 21}, true) //1, 3, 4, 6, 20, 11, 5, 17, 21
	h.Push(2)
	assert.Equal(t, []int{1, 2, 4, 6, 3, 11, 5, 17, 21, 20}, h.arr)

	h = NewHeap([]int{5, 6, 4, 17, 20, 11, 1, 3}, true) //1, 3, 4, 6, 20, 11, 5, 17
	h.Push(2)
	assert.Equal(t, []int{1, 2, 4, 3, 20, 11, 5, 17, 6}, h.arr)

	h = NewHeap([]int{5, 6, 4, 17, 20, 11, 1, 3}, false) //20, 17, 11, 5, 6, 4, 1, 3
	h.Push(33)
	assert.Equal(t, []int{33, 20, 11, 17, 6, 4, 1, 3, 5}, h.arr)
}

func Test_Pop(t *testing.T) {

	h := NewHeap([]int{5, 6, 4, 17, 20, 2, 1, 3, 21}, true)
	assert.Equal(t, 1, *h.Pop())
	assert.Equal(t, []int{2, 3, 4, 6, 20, 5, 21, 17}, h.arr)
	assert.Equal(t, 2, *h.Pop())
	assert.Equal(t, []int{3, 6, 4, 17, 20, 5, 21}, h.arr)
	assert.Equal(t, 3, *h.Pop())
	assert.Equal(t, []int{4, 6, 5, 17, 20, 21}, h.arr)
	assert.Equal(t, 4, *h.Pop())
	assert.Equal(t, []int{5, 6, 21, 17, 20}, h.arr)
	assert.Equal(t, 5, *h.Pop())
	assert.Equal(t, []int{6, 17, 21, 20}, h.arr)
	assert.Equal(t, 6, *h.Pop())
	assert.Equal(t, []int{17, 20, 21}, h.arr)
	assert.Equal(t, 17, *h.Pop())
	assert.Equal(t, []int{20, 21}, h.arr)
	assert.Equal(t, 20, *h.Pop())
	assert.Equal(t, []int{21}, h.arr)
	assert.Equal(t, 21, *h.Pop())
	assert.Equal(t, []int{}, h.arr)
	assert.Nil(t, h.Pop())
	assert.Equal(t, []int{}, h.arr)

	h = NewHeap([]int{17, 20, 2, 1, 3, 21}, false)
	assert.Equal(t, 21, *h.Pop())
	assert.Equal(t, []int{20, 3, 17, 1, 2}, h.arr)
	assert.Equal(t, 20, *h.Pop())
	assert.Equal(t, []int{17, 3, 2, 1}, h.arr)
	assert.Equal(t, 17, *h.Pop())
	assert.Equal(t, []int{3, 1, 2}, h.arr)
	assert.Equal(t, 3, *h.Pop())
	assert.Equal(t, []int{2, 1}, h.arr)
	assert.Equal(t, 2, *h.Pop())
	assert.Equal(t, []int{1}, h.arr)
	assert.Equal(t, 1, *h.Pop())
	assert.Equal(t, []int{}, h.arr)
	assert.Nil(t, h.Pop())
	assert.Equal(t, []int{}, h.arr)
}
