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

func QuickSort(arr []int, start, end int) {
	if end <= start {
		return
	}
	if Compare == nil {
		return
	}

	pivot := end
	pos := start
	for i := start; i < end; i++ {
		if !Compare(arr[pivot], arr[i]) {
			continue
		}
		arr[i], arr[pos] = arr[pos], arr[i]
		pos++
	}
	arr[pivot], arr[pos] = arr[pos], arr[pivot]

	QuickSort(arr, start, pos-1)
	QuickSort(arr, pos+1, end)
}

func main() {

}
