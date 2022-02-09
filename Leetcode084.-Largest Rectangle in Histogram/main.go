package main

import (
	"fmt"
)

//https://leetcode.com/problems/largest-rectangle-in-histogram/
// Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

// Example 1:

// Input: heights = [2,1,5,6,2,3]
// Output: 10
// Explanation: The above is a histogram where width of each bar is 1.
// The largest rectangle is shown in the red area, which has an area = 10 units.
// Example 2:

// Input: heights = [2,4]
// Output: 4

// =================================== //

// import (
// 	"fmt"
// 	"sync"
// )

// // Item the type of the stack
// type Item interface{}

// // ItemStack the stack of Items
// type ItemStack struct {
// 	items []Item
// 	lock  sync.RWMutex
// }

// // New creates a new ItemStack
// func NewStackx() *ItemStack {
// 	s := &ItemStack{}
// 	s.items = []Item{}
// 	return s
// }

// // Pirnt prints all the elements
// func (s *ItemStack) Print() {
// 	fmt.Println(s.items)
// }

// // Push adds an Item to the top of the stack
// func (s *ItemStack) Push(t Item) {
// 	s.lock.Lock()
// 	defer s.lock.Unlock()
// 	s.items = append(s.items, t)
// }

// // Pop removes an Item from the top of the stack
// func (s *ItemStack) Pop() Item {
// 	s.lock.Lock()
// 	defer s.lock.Unlock()
// 	if len(s.items) == 0 {
// 		return nil
// 	}
// 	item := s.items[len(s.items)-1]
// 	s.items = s.items[0 : len(s.items)-1]
// 	return item
// }

// =================================== //

// type Stack struct {
// 	list *list.List
// 	lock *sync.RWMutex
// }

// func NewStack() *Stack {
// 	list := list.New()
// 	l := &sync.RWMutex{}
// 	return &Stack{list, l}
// }

// func (stack *Stack) Push(value interface{}) {
// 	stack.lock.Lock()
// 	defer stack.lock.Unlock()
// 	stack.list.PushBack(value)
// }

// func (stack *Stack) Pop() interface{} {
// 	stack.lock.Lock()
// 	defer stack.lock.Unlock()
// 	e := stack.list.Back()
// 	if e != nil {
// 		stack.list.Remove(e)
// 		return e.Value
// 	}
// 	return nil
// }

// func (stack *Stack) Peak() interface{} {
// 	e := stack.list.Back()
// 	if e != nil {
// 		return e.Value
// 	}

// 	return nil
// }

// func (stack *Stack) Len() int {
// 	return stack.list.Len()
// }

// func (stack *Stack) Empty() bool {
// 	return stack.list.Len() == 0
// }

// ================================= //
//簡單版 stack , 要keep return 值
type stack struct {
	item []int
}

func (s *stack) Peek() *int {
	if s.Len() == 0 {
		return nil
	}
	v := s.item[s.Len()-1]
	return &v
}

func (s *stack) Push(v int) {
	s.item = append(s.item, v)
}

func (s *stack) Pop() *int {
	if s.Len() == 0 {
		return nil
	}
	v := s.item[s.Len()-1]
	s.item = s.item[0 : s.Len()-1]
	return &v
}

func (s *stack) Len() int {
	return len(s.item)
}

func NewStack(capacity int) *stack {
	item := make([]int, 0, capacity)

	return &stack{item: item}
}

func largestRectangleArea(heights []int) int {
	// if heights == nil || len(heights) <= 0 {
	// 	return 0
	// }

	maxArea := 0
	stack := NewStack(len(heights))

	heights = append(heights, 0) // dummy end  to finish all

	for i, h := range heights {
		for stack.Len() > 0 && heights[*stack.Peek()] >= h {
			w := i
			idx := stack.Pop()
			if stack.Len() > 0 {
				w = i - 1 - *stack.Peek()
			}
			area := w * heights[*idx]

			if area > maxArea {
				maxArea = area
			}
		}
		stack.Push(i)
	}
	return maxArea
}

func main() {
	nums := []int{1, 3, 2, 4, 3, 2, 1}

	fmt.Println("1 -", nums, " : max area = ", largestRectangleArea(nums))

	nums = []int{1, 3, 0, 2, 4, 0, 3, 2, 1}

	fmt.Println("1 -", nums, " : max area = ", largestRectangleArea(nums))

}
