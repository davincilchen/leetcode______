package main

type Heap struct {
	arr []int
}

func NewHeap(arr []int) *Heap {
	return &Heap{arr: arr}
}

func (t *Heap) leftchildIndex(index int) int {
	return index*2 + 1
}

func (t *Heap) rightchildIndex(index int) int {
	return index*2 + 2
}

func (t *Heap) swap(i, j int) {
	tmp := t.arr[i]
	t.arr[i] = t.arr[j]
	t.arr[j] = tmp
}

func (t *Heap) leaf(index, size int) bool {
	if index >= size/2 && index <= size-1 {
		return true
	}

	return false
}

func (m *Heap) downHeapifyForMin(current int, size int) {
	if m.leaf(current, size) {
		return
	}
	smallest := current
	leftChildIndex := m.leftchildIndex(current)
	rightRightIndex := m.rightchildIndex(current)
	if leftChildIndex < size && m.arr[leftChildIndex] < m.arr[smallest] {
		smallest = leftChildIndex
	}
	if rightRightIndex < size && m.arr[rightRightIndex] < m.arr[smallest] {
		smallest = rightRightIndex
	}
	if smallest != current {
		m.swap(current, smallest)
		m.downHeapifyForMin(smallest, size)
	}

}

func (m *Heap) downHeapifyForMax(current int, size int) {
	if m.leaf(current, size) {
		return
	}
	bigger := current
	child := m.leftchildIndex(current)
	if child < size && m.arr[child] > m.arr[bigger] {
		bigger = child
	}
	child = m.rightchildIndex(current)
	if child < size && m.arr[child] > m.arr[bigger] {
		bigger = child
	}

	if bigger != current {
		m.swap(bigger, current)
		m.downHeapifyForMax(bigger, size)
	}

}

func (m *Heap) buildMinHeap(size int) {
	for index := ((size / 2) - 1); index >= 0; index-- {
		m.downHeapifyForMin(index, size)
	}
}

func (m *Heap) buildMaxHeap(size int) {
	for i := size/2 - 1; i >= 0; i-- {
		m.downHeapifyForMax(i, size)
	}
}
