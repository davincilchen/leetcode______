package main

type compare func(a, b int) bool

var Compare compare

func Bigger(a, b int) bool {
	return a >= b
}

func Smaller(a, b int) bool {
	return a <= b
}

// 5, 6, 4, 17, 20, 2, 1, 3,     21

// compare
// pos              i
// 5, 6, 4, 17, 20, 2, 1, 3,

// swap
// pos              i
// 2, 6, 4, 17, 20, 5, 1, 3,

// compare
//    pos              i
// 2, 6, 4, 17, 20, 5, 1, 3,

// swap
//    pos              i
// 2, 1, 4, 17, 20, 5, 6, 3,

// last swap
//       pos              pivot
// 2, 1, 4, 17, 20, 5, 6, 3,

// 2, 1, 3, 17, 20, 5, 6, 4,
//pivot 扭 ,unstable

//quick sort ,不要的留在原地
//最後會被要的or pivot 扭回去
func Sort(arr []int) []int {

	QuickSort(arr, 0, len(arr)-1)
	return arr
}

// 平均情况下快速排序的时间复杂度是Θ(n\lgn)，最坏情况是Θ(n2)。

// 当划分产生的两个子问题分别包含 n-1 和 0 个元素时，最坏情况发生。划分操作的时间复杂度为Θ(n)，T(0)=Θ(1)，这时算法运行时间的递归式为 T(n)=T(n−1)+T(0)+Θ(n)=T(n−1)+Θ(n)，解为T(n)=Θ(n2)。

// 当划分产生的两个子问题分别包含⌊n/2⌋和⌈n/2⌉−1个元素时，最好情况发生。算法运行时间递归式为 T(n)=2T(n/2)+Θ(n)，解为T(n)=Θ(nlgn)。

// 事实上只要划分是常数比例的，算法的运行时间总是O(nlgn)。 假设按照 9:1 划分，每层代价最多为 cn，递归深度为 log10/9n=Θ(lgn)，故排序的总代价为O(nlgn)。

//T(n)=2T(n/2)+cn
//快速排序時間 = 兩個半邊快速排序+ cur partiion時間
//參考merge sort (divide and conquer)可以推導出來

//T(n)=2T(n/2)+cn
//同理
//T(n/2) = 2T(n/4)+cn/2
//帶入1
//T(n)=2 * T(n/2)+cn
//T(n)=2 * (2T(n/4)+cn/2)+cn
//T(n)=4T(n/4)+2*cn/2+cn
//T(n)=4T(n/4)+2cn

//ex n=4
//T(4)=4T(1)+2cn
//T(4)=2[的k次方]* T(1)+ 2[=log4]cn
//T(4)=2[的k次方]* T(1)+ 2[=log4]cn

//T(4)=2[的k次方]* T(1)+ 2[log4]cn

// 2[的k次方] = n
// logN = K

//層數為 0, 1 ~ k層 = 總共k+1層

func getPivot(arr []int, start, end int) int {
	if end-start < 2 {
		return end
	}
	tmp := (start + end) / 2

	// Checking for b
	if arr[start] < arr[tmp] && arr[tmp] < arr[end] || arr[end] < arr[tmp] && arr[tmp] < arr[start] {
		return tmp
	} else if (arr[tmp] < arr[start] && arr[start] < arr[end]) || (arr[end] < arr[start] && arr[start] < arr[tmp]) {
		return start
	} else {
		return end
	}

}
func QuickSort(arr []int, start, end int) {
	if end <= start { //邊界條件
		return
	}
	if Compare == nil {
		return
	}

	pos := Partition(arr, start, end)

	QuickSort(arr, start, pos-1)
	QuickSort(arr, pos+1, end)
}

func Partition(arr []int, start, end int) int {
	pivot := getPivot(arr, start, end)

	// !!
	arr[pivot], arr[end] = arr[end], arr[pivot]
	pivot = end
	//
	pos := start
	for i := start; i < end; i++ {
		if !Compare(arr[pivot], arr[i]) {
			continue
		}
		arr[pos], arr[i] = arr[i], arr[pos]
		pos++
	}
	arr[pos], arr[pivot] = arr[pivot], arr[pos]
	return pos
}

func main() {

}
