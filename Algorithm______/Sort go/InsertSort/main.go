package main

//遞增
//分為以排序和未排序部分
//把未排序的插入以排序中
//可以使用額外空間, 和 in place兩種版本
//index i之前 , 由i-1開始,比a[i]大的都往右移
// func Sort(nums []int) []int {

// 	for i := 1; i < len(nums); i++ {
// 		tmp := nums[i]
// 		j := i - 1
// 		for {
// 			if j < 0 {
// 				break
// 			}
// 			if nums[j] < tmp { //沒有比較大的了
// 				break
// 			}
// 			nums[j+1] = nums[j]
// 			j--
// 		}

// 		//if i != j+1 {
// 		nums[j+1] = tmp
// 		//}

// 	}

// 	return nums
// }

func Sort(nums []int) []int {

	j := 0
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		for j = i - 1; j >= 0; j-- {
			if tmp > nums[j] { //沒有比較大的了(已經排序過了)
				break
			}
			nums[j+1] = nums[j] //往右移一格
		}
		nums[j+1] = tmp
	}
	return nums
}

func main() {

}
