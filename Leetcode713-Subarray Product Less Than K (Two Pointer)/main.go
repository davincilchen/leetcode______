package main

//看一下寫得不好
// func numSubarrayProductLessThanK(nums []int, k int) int {
// 	// if k <= 1 {
// 	// 	return 0
// 	// }
// 	l := 0
// 	r := 0
// 	cnt := 0
// 	product := 1
// 	for r < len(nums) {
// 		product *= nums[r]
// 		if product < k {
// 			cnt += (r - l + 1)
// 			r++
// 			continue
// 		}

// 		for product >= k {
// 			//l++ //not before
// 			product /= nums[l]
// 			l++

// 			if l >= r {
// 				break
// 			}
// 		}

// 		if product < k {
// 			cnt += (r - l + 1)
// 			r++
// 			continue
// 		}
// 		r++
// 	}

// 	return cnt
// }

// Runtime: 80 ms, faster than 96.43% of Go online submissions for Subarray Product Less Than K.
// Memory Usage: 7.5 MB, less than 72.14% of Go online submissions for Subarray Product Less Than K.
// Next challenges:
// Maximum Product Subarray
// Maximum Size Subarray Sum Equals k
// Subarray Sum Equals K
// Two Sum Less Than K
// Number of Smooth Descent Periods of a Stock

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	cnt := 0
	product := 1
	for l, r := 0, 0; r < len(nums); r++ {
		product *= nums[r]
		//for product >= k && l <= r {
		for product >= k && l < r {
			product /= nums[l]
			l++
		}
		if product < k {
			//fmt.Println("l = ", l, "r = ", r)
			cnt += (r - l + 1)
		}
	}

	return cnt
}
func main() {

}
