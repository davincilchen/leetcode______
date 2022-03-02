package main

func MaxHeapifyTopDown(currLastNode int, in []int) {
	currEleIndex := 0
	for {
		child := 2*currEleIndex + 1
		// check left child is out of the bound or not
		if child > currLastNode {
			break
		}
		// check right child is out of the bound or not
		// and if right child > left child, shoose right child
		if child+1 <= currLastNode && in[child] < in[child+1] {
			child++
		}
		// if current root > both right and left child, just return
		if in[child] < in[currEleIndex] {
			return
		}
		// if root < the largest child
		// swap root with the largest child
		in[currEleIndex], in[child] = in[child], in[currEleIndex]
		// keep going to next sub-heap
		currEleIndex = child
	}
}

func HeapifyMin(arr []int) []int {

	t := NewHeap(arr)
	t.buildMinHeap(len(arr))
	return t.arr
}

func HeapifyMax(arr []int) []int {

	t := NewHeap(arr)
	t.buildMaxHeap(len(arr))
	return t.arr
}
func main() {

}
