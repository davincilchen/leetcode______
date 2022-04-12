package main

import "fmt"

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

// product 一開始是 1，ans 一開始是 0。
// 假設一開始左手指向 10，右手也指向 10，這時候 product *= 10 == 10，所以我們知道可以把 ans += 1 == 1。
// 把右手指往右移，右手指向 5，這時候 product *= 5 == 50，因為依然小於 k，所以我們知道 [10, 5] 跟 [5] 這兩個 subarray 的乘積都小於 k。 所以可以把 ans += (r-l+1) == 1 + (1-0+1) == 3。（r 表示右手指位置，l 則表示左手指位置）
//			cnt += (r - l + 1)
// 此次數量(r - l + 1)
// 如果是3 就是 最左邊單一個, 最左邊兩個相乘, 總共三個相乘
// 總共這三種
// 因為是長到最符合才會計算

// 第一次包含右邊的排列組合會不一樣
// 第一次少左邊的排列組合也會不一樣

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
			fmt.Println("l = ", l, "r = ", r, cnt)
		}
	}

	return cnt
}
func main() {

}
