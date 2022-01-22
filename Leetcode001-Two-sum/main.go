package main

import "fmt"

//https://leetcode.com/problems/two-sum/
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:

// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.

// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

// Runtime: 4 ms, faster than 97.19% of Go online submissions for Two Sum.
// Memory Usage: 4.3 MB, less than 34.17% of Go online submissions for Two Sum.

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i, v := range nums {
		found, ok := m[target-v]
		if ok {
			return []int{found, i}
		}
		m[v] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(nums, twoSum(nums, target), "for target =", target)
	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(nums, twoSum(nums, target), "for target =", target)
	nums = []int{3, 3}
	target = 6
	fmt.Println(nums, twoSum(nums, target), "for target =", target)
	nums = []int{1, 2, 7, 3, 3}
	target = 6
	fmt.Println(nums, twoSum(nums, target), "for target =", target)
	fmt.Println("---- ")

}
