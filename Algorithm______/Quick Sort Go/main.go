package main

type compare func(a, b int) bool

var Compare compare

func Bigger(a, b int) bool {
	return a >= b
}

func Smaller(a, b int) bool {
	return a <= b
}

func Sort(arr []int) []int {

	QuickSort(arr, 0, len(arr)-1)
	return arr
}

func QuickSort(arr []int, start, end int) {

	if Compare == nil {
		return
	}
	if end <= start { //終止條件在遞迴層
		return
	}

	p := Partition(arr, start, end)
	QuickSort(arr, start, p-1)
	QuickSort(arr, p+1, end)

}

func Partition(arr []int, start, end int) int { // pivot

	i := -1
	tmp := 0
	pivot := arr[end]
	for j := start; j <= end-1; j++ {
		if Compare(arr[j], pivot) {
			continue
		}
		i++

		pos := start + i
		tmp = arr[pos]
		arr[pos] = arr[j]
		arr[j] = tmp
	}

	i++

	pos := start + i
	tmp = arr[pos]
	arr[pos] = pivot
	arr[end] = tmp

	return pos
}

func main() {

}
