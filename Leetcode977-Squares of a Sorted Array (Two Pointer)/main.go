package main

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/squares-of-a-sorted-array/
// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

// Example 1:

// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].
// Example 2:

// Input: nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]

// Constraints:

// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums is sorted in non-decreasing order.

// Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?

//SortedSquares1:
//Runtime: 59 ms, faster than 25.69% of Go online submissions for Squares of a Sorted Array.
//Memory Usage: 7.9 MB, less than 13.46% of Go online submissions for Squares of a Sorted Array.

//SortedSquares2 + SliceMove3:
// Runtime: 356 ms, faster than 6.76% of Go online submissions for Squares of a Sorted Array.
// Memory Usage: 8.4 MB, less than 8.17% of Go online submissions for Squares of a Sorted Array.

//SortedSquares2 + SliceMove2: **
// Runtime: 184 ms, faster than 10.45% of Go online submissions for Squares of a Sorted Array.
// Memory Usage: 6.9 MB, less than 74.43% of Go online submissions for Squares of a Sorted Array.

//SortedSquares3:  ***
//Runtime: 38 ms, faster than 49.66% of Go online submissions for Squares of a Sorted Array.
//Memory Usage: 7.2 MB, less than 43.27% of Go online submissions for Squares of a Sorted Array.

func sortedSquares(nums []int) []int {

	return SortedSquares1(nums)
	//return SortedSquares2(nums)
}

func SortedSquares2(nums []int) []int {
	//math.Pow(2, x)

	k := len(nums) - 1
	for k > 0 {
		i := 0
		if nums[i] >= 0 {
			break
		}

		for j := k; j > 0; j-- {
			fmt.Println(i, j, k, nums[i], nums[j])
			if nums[i]+nums[j] >= 0 {
				//if (nums[i]*nums[i])+(nums[j]*nums[j]) >= 0 {
				//正數比較多, 合法
				k-- //for all ok
				continue
			}

			//重新插入排序
			nums = SliceMove3(i, j, nums)
			//fmt.Print(" ret ")
			// for xx := range nums {
			// 	fmt.Print(nums[xx], " ")
			// }
			//fmt.Println("")
			k = j - 1
			break
		}

	}

	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}

	return nums
}

func SortedSquares3(nums []int) []int {
	highestIdx := len(nums) - 1
	l := 0
	r := highestIdx
	squares := make([]int, len(nums))

	for l <= r {
		leftSquare := nums[l] * nums[l]
		rightSquare := nums[r] * nums[r]

		if leftSquare >= rightSquare {
			squares[highestIdx] = leftSquare
			highestIdx--
			l++
		} else {
			squares[highestIdx] = rightSquare
			highestIdx--
			r--
		}
	}
	return squares
}

func SortedSquares4(nums []int) []int { //還不正確
	highestIdx := len(nums) - 1
	l := 0
	r := highestIdx

	leftSquare := nums[l] * nums[l]
	rightSquare := nums[r] * nums[r]
	for l <= r {
		if leftSquare >= rightSquare {
			if l < len(nums) { //要用的先取出來才部會被蓋掉
				leftSquare = nums[l] * nums[l]
			}
			nums[highestIdx] = leftSquare
			highestIdx--
			l++
		} else {
			if r > 0 { //要用的先取出來才部會被蓋掉
				rightSquare = nums[r] * nums[r]
			}
			nums[highestIdx] = rightSquare
			highestIdx--
			r--

		}
	}
	return nums
}

func SortedSquares1(nums []int) []int {
	//math.Pow(2, x)

	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}

	sort.Ints(nums) //可以看轉換成IntSlice 的code

	return nums
}

func SliceMove2(indexToRemove, indexWhereToInsert int, s []int) []int {
	if indexToRemove < 0 || indexToRemove >= len(s) {
		return s
	}

	if indexWhereToInsert < 0 || indexWhereToInsert >= len(s) {
		return s
	}

	//ss = append(ss[:indexToRemove], tmp) //error
	//ss = append(ss, ss[indexToRemove+1:]...) //error

	tmp := s[indexToRemove]
	for i := indexToRemove; i < indexWhereToInsert; i++ {
		s[i] = s[i+1]
	}
	s[indexWhereToInsert] = tmp
	return s
}

func SliceMove3(indexToRemove, indexWhereToInsert int, s []int) []int {
	if indexToRemove < 0 || indexToRemove >= len(s) {
		return s
	}

	if indexWhereToInsert < 0 || indexWhereToInsert >= len(s) {
		return s
	}

	return moveInt(s, indexToRemove, indexWhereToInsert)

}

// .. //
func insertInt(array []int, value int, index int) []int {
	//aa := append([]int{value}, array[index:]...)
	//fmt.Printf("SliceMove3 address array %p , aa %p \n", array, aa)
	//不同樣位置
	//newSlice slice不同樣參考 append 以前面參數的slice為保留位置
	//sliceMove3 address array 0xc0000a4780 , aa 0xc0000b6510
	//SliceMove3 address array 0xc0000a46e0 , aa 0xc0000b64e0

	//new := append([]int{value}, array[index:]...)
	//return append(array[:index], new...)

	return append(array[:index], append([]int{value}, array[index:]...)...)
}

func removeInt(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func moveInt(array []int, srcIndex int, dstIndex int) []int {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

func main() {

}
