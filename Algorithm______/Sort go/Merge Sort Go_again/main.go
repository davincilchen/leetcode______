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

	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2

	arr1 := arr[0:mid] //index  要注意
	arr2 := arr[mid:]  //index  要注意
	return merge(MergeSort(arr1), MergeSort(arr2))

}

func merge(arr1, arr2 []int) []int {
	cnt1 := len(arr1)
	cnt2 := len(arr2)
	ret := make([]int, cnt1+cnt2)

	j1, j2 := 0, 0
	i := 0
	for i = 0; i < len(ret); i++ {
		if j1 >= cnt1 || j2 >= cnt2 {
			break
		}
		v1 := arr1[j1]
		v2 := arr2[j2]
		if Compare(v1, v2) {
			ret[i] = v1
			j1++
		} else {
			ret[i] = v2
			j2++
		}
	}

	for k := j1; k < cnt1; k++ {
		ret[i] = arr1[k]
		i++
	}

	for k := j2; k < cnt2; k++ {
		ret[i] = arr2[k]
		i++
	}

	return ret
}

func main() {

}
