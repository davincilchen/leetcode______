package main

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104

// 我們討論1到i的最大子陣列和，應該分兩種情況：
// a. 這個子陣列包括index i
// b. 這個子陣列不包括index i
// 如果是b的話，那答案其實就是1到i-1的最大子陣列和；
// 但如果是a的話，我們好像還需要一些條件？
// 什麼條件呢？就是1到i-1中，包含i-1元素的最大子陣列和。
// 因為如果是前面所提到的a狀況的話，
// 那麼1到i的最大連續子陣列，就必須從i->i-1連回去才行(或乾脆是只有i)。

//到現在為止"包"含當前這個元素"的最大子陣列和，我們命名為curr(current)
//到目前為止的最大子陣列和，我們命名為res(result)

func maxSubArray(nums []int) int {
	//fmt.Println(nums)
	subMax := nums[0]
	max := nums[0]
	n := len(nums)

	for i := 1; i < n; i++ {
		//lastMax := subMax

		//更新包含nums[i]時的最大值 //因是連續才是subarray
		subMax += nums[i]     //包元nums[i]時的最大值,才能再往右長
		if nums[i] > subMax { //單獨還比較大
			subMax = nums[i]
		}

		// .. //
		if subMax > max {
			max = subMax
		}
		//fmt.Printf("i = %d, max = %d,  subMax = %d , lastMax = %d, nums[i] = %d\n", i, max, subMax, lastMax, nums[i])
	}

	return max
}

func main() {

}
