package main

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104

//兩種方法都是 O(n) ??
// 04/17/2022 21:32	Accepted	97 ms	8.4 MB	golang
// 04/17/2022 21:31	Accepted	92 ms	8.4 MB	golang
//時間差不多

//
//

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

//=====================

// Follow up: If you have figured out the O(n) solution,
// try coding another solution using the divide and conquer approach,
// which is more subtle.

// 第1步。選擇數組的中間元素。
// 因此，最大子數組可能包含或不包含該中間元素。

// Step 2.1 如果最大子數組不包含中間元素，那麼我們可以對中間元素左側的子數組和中間元素右側的子數組應用相同的算法。

// 步驟 2.2 如果最大子數組確實包含中間元素，那麼結果將只是左子數組的最大後綴子數組加上右子數組的最大前綴子數組

// 第 3 步返回這三個答案中的最大值。

// 這是分治解決方案的示例代碼。請在看代碼之前嘗試理解算法

// Step1. Select the middle element of the array.
// So the maximum subarray may contain that middle element or not.

// Step 2.1 If the maximum subarray does not contain the middle element, then we can apply the same algorithm to the the subarray to the left of the middle element and the subarray to the right of the middle element.

// Step 2.2 If the maximum subarray does contain the middle element, then the result will be simply the maximum suffix subarray of the left subarray plus the maximum prefix subarray of the right subarray

// Step 3 return the maximum of those three answer.

// Here is a sample code for divide and conquer solution. Please try to understand the algorithm before look at the code

// class Solution {
// 	public:
// 		int maxSubArray(int A[], int n) {
// 			// IMPORTANT: Please reset any member data you declared, as
// 			// the same Solution instance will be reused for each test case.
// 			if(n==0) return 0;
// 			return maxSubArrayHelperFunction(A,0,n-1);
// 		}

// 		int maxSubArrayHelperFunction(int A[], int left, int right) {
// 			if(right == left) return A[left];
// 			int middle = (left+right)/2;
// 			int leftans = maxSubArrayHelperFunction(A, left, middle);
// 			int rightans = maxSubArrayHelperFunction(A, middle+1, right);
// 			int leftmax = A[middle];
// 			int rightmax = A[middle+1];
// 			int temp = 0;
// 			for(int i=middle;i>=left;i--) {
// 				temp += A[i];
// 				if(temp > leftmax) leftmax = temp;
// 			}
// 			temp = 0;
// 			for(int i=middle+1;i<=right;i++) {
// 				temp += A[i];
// 				if(temp > rightmax) rightmax = temp;
// 			}
// 			return max(max(leftans, rightans),leftmax+rightmax);
// 		}
// 	};
func maxSubArray1(nums []int) int {
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

//divide and conquer
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray2(nums []int) int {
	return maxSubArrayDivide(nums, 0, len(nums)-1)
}

func maxSubArrayDivide(nums []int, left, right int) int {
	if left == right { //只有一個元素
		return nums[left]
	}

	mid := (left + right) / 2
	ansL := maxSubArrayDivide(nums, left, mid)
	ansR := maxSubArrayDivide(nums, mid+1, right)

	tmp := 0
	lVal := 0
	rVal := 0
	for i := mid - 1; i >= left; i-- {
		tmp += nums[i]
		if tmp > lVal {
			lVal = tmp
		}
	}

	tmp = 0
	for i := mid + 1; i <= right; i++ {
		tmp += nums[i]
		if tmp > rVal {
			rVal = tmp
		}
	}
	cVal := lVal + nums[mid] + rVal

	return max(max(ansL, ansR), cVal)
}
