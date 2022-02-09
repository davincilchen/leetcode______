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
type stack []int

func (s stack) Peek() int {
	return s[len(s)-1]
}

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Len() int {
	return len(s)
}

func newStack(capacity int) stack {
	return stack(make([]int, 0, capacity))
}

func largestRectangleArea(heights []int) int {
	// if heights == nil || len(heights) <= 0 {
	// 	return 0
	// }

	maxArea := 0
	tmp := -1
	stack := newStack(len(heights))
	heights = append(heights, 0) // dummy end  to finish all

	for i, h := range heights {
		for stack.Len() > 0 && heights[stack.Peek()] >= h {
			w := i
			//stack, tmpH := stack.Pop() //new stack // infinite loop
			stack, tmp = stack.Pop()
			if stack.Len() > 0 {
				w = i - 1 - stack.Peek()
			}
			area := w * heights[tmp]
			if area > maxArea {
				maxArea = area
			}

		}
		// if h == -1 {
		// 	continue
		// }
		stack = stack.Push(i)
	}

	return maxArea
}

func largestRectangleArea_ori(heights []int) int {
	if heights == nil || len(heights) <= 0 {
		return 0
	}

	maxArea := 0
	stack := newStack(len(heights))
	heights = append(heights, 0) // dummy end

	for i := 0; i < len(heights); i++ {
		ch := heights[i]
		for stack.Len() > 0 && ch <= heights[stack.Peek()] {
			h := heights[stack.Peek()]
			stack, _ = stack.Pop()
			w := i
			if stack.Len() > 0 {
				w = i - 1 - stack.Peek()
			}

			if area := h * w; area > maxArea {
				maxArea = area
			}
		}

		stack = stack.Push(i)
	}

	return maxArea
}

func main() {
	nums := []int{1, 3, 2, 4, 3, 2, 1}
	fmt.Println("0 -", nums, " : max area = ", largestRectangleArea_ori(nums))

	fmt.Println("1 -", nums, " : max area = ", largestRectangleArea(nums))

	nums = []int{1, 3, 0, 2, 4, 0, 3, 2, 1}
	fmt.Println("0 -", nums, " : max area = ", largestRectangleArea_ori(nums))

	fmt.Println("1 -", nums, " : max area = ", largestRectangleArea(nums))

}
