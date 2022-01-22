package main

//盡可能展到最大
//再縮小
//再儘可能展到最到
//再縮小

//箭頭標示寫不好要改寫的地方

// func minSubArrayLen(target int, nums []int) int {
// 	fmt.Println("========")
// 	min := 0
// 	for sum, l, r := 0, 0, 0; r < len(nums); r++ {
// 		sum += nums[r]
// 		if sum < target {
// 			continue
// 		}
// 		fmt.Println(" --min sum, l, r =", min, sum, l, r)
// 		---> // if min > r-l+1 {
// 		---> // 	min = r - l + 1
// 		---> // }

// 		---> //for sum >= target && l < r {
// 		for sum >= target && l <= r {
// 			if min == 0 {
// 				min = r - l + 1
// 			} else {
// 				if min > r-l+1 {
// 					min = r - l + 1
// 				}

// 			}
// 			fmt.Println("** --min sum, l, r =", min, sum, l, r)
// 			sum -= nums[l]
// 			l++

// 		}

// 	}

// 	return min
// }

// Success
// Details
// Runtime: 14 ms, faster than 25.08% of Go online submissions for Minimum Size Subarray Sum.
// Memory Usage: 4 MB, less than 51.66% of Go online submissions for Minimum Size Subarray Sum.
// Next challenges:
// Minimum Window Substring
// Maximum Size Subarray Sum Equals k
// Maximum Length of Repeated Subarray
// Minimum Operations to Reduce X to Zero
// K Radius Subarray Averages

//min 寫得不好 初始化有問題
// func minSubArrayLen(target int, nums []int) int {
// 	//fmt.Println("========")
// 	min := 0
// 	for sum, l, r := 0, 0, 0; r < len(nums); r++ {
// 		sum += nums[r]
// 		if sum < target {
// 			continue
// 		}
// 		//fmt.Println(" --min sum, l, r =", min, sum, l, r)
// 		// if min > r-l+1 {
// 		// 	min = r - l + 1
// 		// }

// 		//for sum >= target && l < r {
// 		for sum >= target && l <= r {
// 			if min == 0 {
// 				min = r - l + 1
// 			} else {
// 				if min > r-l+1 {
// 					min = r - l + 1
// 				}

// 			}
// 			//fmt.Println("** --min sum, l, r =", min, sum, l, r)
// 			sum -= nums[l]
// 			l++

// 		}

// 	}

// 	return min
// }

//min 還需要多一個bool判斷 還不夠好
// func minSubArrayLen(target int, nums []int) int {
// 	size := len(nums)
// 	minW := size
// 	pass := false

// 	for sum, l, r := 0, 0, 0; r < size; r++ {
// 		sum += nums[r]
// 		for sum >= target && l <= r {
// 			pass = true
// 			if minW > r-l+1 {
// 				minW = r - l + 1
// 			}
// 			sum -= nums[l]
// 			l++
// 		}
// 	}

// 	if pass {
// 		return minW
// 	}

// 	return 0

// }

// Runtime: 11 ms, faster than 37.46% of Go online submissions for Minimum Size Subarray Sum.
// Memory Usage: 4 MB, less than 51.66% of Go online submissions for Minimum Size Subarray Sum.
// Next challenges:
// Minimum Window Substring
// Maximum Size Subarray Sum Equals k
// Maximum Length of Repeated Subarray
// Minimum Operations to Reduce X to Zero
// K Radius Subarray Averages
const INT_MAX = int(^uint(0) >> 1)

//const INT_MIN = ^INT_MAX
//const UINT_MIN uint = 0
//const UINT_MAX = ^uint(0)

func minSubArrayLen(target int, nums []int) int {
	size := len(nums)
	minW := INT_MAX //初始最大值

	for sum, l, r := 0, 0, 0; r < size; r++ {
		sum += nums[r]
		for sum >= target && l <= r {
			if minW > r-l+1 {
				minW = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}

	if minW != INT_MAX {
		return minW
	}

	return 0

}

func main() {

}
