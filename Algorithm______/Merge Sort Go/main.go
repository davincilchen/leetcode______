package main

// 1. Coding Test (RD)
// 2. System Design + General Technical Question (RD)
// 3. Behavior Question (RD manager)
// 4. Background (CTO)

//n*Logn  //每個點會被處理logn 次, 有n個點
type compare func(a, b int) bool

var Compare compare

func BiggerFirst(a, b int) bool {
	return a >= b
}

func SmallerFirst(a, b int) bool {
	return a <= b
}

func Sort(arr []int) []int {
	return MergeSort(arr)
}

func MergeSort(arr []int) []int {

	if Compare == nil {
		return arr
	}

	cnt := len(arr)
	if cnt <= 1 { //終止條件在遞迴層
		return arr
	}

	mid := cnt / 2
	first := MergeSort(arr[:mid])
	second := MergeSort(arr[mid:])
	return merge(first, second)
}

func merge(arr1, arr2 []int) []int {

	if Compare == nil {
		return []int{}
	}

	i, j := 0, 0
	l1 := len(arr1)
	l2 := len(arr2)

	ret := make([]int, 0, l1+l2)
	for {
		if i >= l1 || j >= l2 {
			break
		}
		if Compare(arr1[i], arr2[j]) {
			ret = append(ret, arr1[i])
			i++
		} else {
			ret = append(ret, arr2[j])
			j++
		}
	}

	// for i < l1 {
	// 	ret = append(ret, arr1[i])
	// 	i++
	// }
	if i < l1 {
		ret = append(ret, arr1[i:l1]...)
	}

	// for j < l2 {
	// 	ret = append(ret, arr2[j])
	// 	j++
	// }
	if j < l2 {
		ret = append(ret, arr2[j:l2]...)
	}

	return ret
}

func main() {

}
