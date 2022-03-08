package main

type Heap struct {
	min bool
	arr []int
}

func NewHeap(arr []int, min bool) *Heap {
	h := &Heap{min: min, arr: arr}
	if min {
		h.buildMinHeap()
	} else {
		h.buildMaxHeap()
	}
	return h
}

func (t *Heap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (t *Heap) leftchildIndex(index int) int {
	return index*2 + 1
}

func (t *Heap) rightchildIndex(index int) int {
	return index*2 + 2
}

func (t *Heap) swap(i, j int) {
	// tmp := t.arr[i]
	// t.arr[i] = t.arr[j]
	// t.arr[j] = tmp

	t.arr[i], t.arr[j] = t.arr[j], t.arr[i]
}

func (t *Heap) leaf(index, size int) bool {
	if index >= size/2 && index <= size-1 {
		return true
	}

	return false
}

//shift down
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

//shift down
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

func (m *Heap) buildMinHeap() { //Heapify的時間複雜度, O(N) 非 O(NlgN)
	size := len(m.arr)
	for index := ((size / 2) - 1); index >= 0; index-- {
		m.downHeapifyForMin(index, size)
	}
}

func (m *Heap) buildMaxHeap() {
	size := len(m.arr)
	for i := size/2 - 1; i >= 0; i-- {
		m.downHeapifyForMax(i, size)
	}
}

func (m *Heap) decreaseSort() []int {
	size := len(m.arr)
	for i := size - 1; i > 0; i-- {
		// Move current root to end
		m.swap(0, i)              //i : last element
		m.downHeapifyForMin(0, i) //i : new size
	}
	return m.arr
}

func (m *Heap) increaseSort() []int {
	size := len(m.arr)
	for i := size - 1; i > 0; i-- {
		// Move current root to end
		m.swap(0, i)              //i : last element
		m.downHeapifyForMax(0, i) //i : new size
	}
	return m.arr
}

func (m *Heap) Sort() []int {
	if m.min {
		return m.decreaseSort()
	}
	return m.increaseSort()
}

func (m *Heap) Pop() *int {
	cnt := len(m.arr)
	if cnt == 0 {
		return nil
	}
	ret := m.arr[0]
	m.arr[0] = m.arr[cnt-1]
	m.arr = m.arr[:cnt-1]

	if m.min {
		m.downHeapifyForMin(0, cnt-1)
	} else {
		m.downHeapifyForMax(0, cnt-1)
	}
	return &ret

}

//heapq.heappushpop(heap, item)
//這個組合函式比呼叫 heappush() 之後呼叫 heappop() 更有效率。 heapq. heapify (x)¶. 在線性時間內將list x ..
func (m *Heap) Push(in int) {
	m.arr = append(m.arr, in)

	size := len(m.arr)
	p := m.parentIndex(size - 1)
	for p >= 0 {
		if m.min {
			m.downHeapifyForMin(p, size)
		} else {
			m.downHeapifyForMax(p, size)
		}
		if p == 0 {
			break
		}
		p = m.parentIndex(p)
	}

}

//需要讀
//https://docs.python.org/zh-tw/3/library/heapq.html

// heapq.merge(*iterables, key=None, reverse=False)
// 合併多個已排序的輸入並產生單一且已排序的輸出（舉例：合併來自多個 log 檔中有時間戳記的項目）。回傳一個 iterator 包含已經排序的值。

// 和 sorted(itertools.chain(*iterables)) 類似但回傳值是一個 iterable ，不會一次把所有資料都放進記憶體中，並且假設每一個輸入都已經（由小到大）排序過了。

// 有兩個選用參數，指定時必須被當作關鍵字參數指定。

// key 參數指定了一個 key function 引數，用來從每一個輸入的元素中決定一個比較的依據。預設的值是 None （直接比較元素）。

// reverse 为一个布尔值。 如果设为 True，则输入元素将按比较结果逆序进行合并。 要达成与 sorted(itertools.chain(*iterables), reverse=True) 类似的行为，所有可迭代对象必须是已从大到小排序的。

// 3.5 版更變: 加入選用參數 key 和 reverse 。

// heapq.nlargest(n, iterable, key=None)
// 从 iterable 所定义的数据集中返回前 n 个最大元素组成的列表。 如果提供了 key 则其应指定一个单参数的函数，用于从 iterable 的每个元素中提取比较键 (例如 key=str.lower)。 等价于: sorted(iterable, key=key, reverse=True)[:n]。

// heapq.nsmallest(n, iterable, key=None)
// 从 iterable 所定义的数据集中返回前 n 个最小元素组成的列表。 如果提供了 key 则其应指定一个单参数的函数，用于从 iterable 的每个元素中提取比较键 (例如 key=str.lower)。 等价于: sorted(iterable, key=key)[:n]。

// 後兩個函式在 n 值比較小時有最好的表現。對於較大的 n 值，只用 sorted() 函式會更有效率。同樣地，當 n==1 時，使用內建函式 min() 和 max() 會有更好的效率。如果需要重複使用這些函式，可以考慮將 iterable 轉成真正的 heap 。

// 基礎範例

//優先佇列 (Priority Queue) 實作細節
// 优先队列 是堆的常用场合，并且它的实现包含了多个挑战：

// 排序的穩定性：你如何將兩個擁有相同 priority 的 task 按照他們被加入的順序回傳。

// Tuple的排序在某些情況下會壞掉，例如當 Tuple (priority, task) 的 priorities 相等且 tasks 沒有一個預設的排序時。

// 當一個 heap 中 task 的 priority 改變時，你如何將它移到 heap 正確的位置上。

// 或者一個還沒被解決的 task 需要被刪除時，你要如何從佇列中找到並刪除指定的 task。

// 一個針對前兩個問題的解法是：儲存一個包含 priority 、 entry count 和 task 三個元素的 tuple 。兩個 task 有相同 priority 時， entry count 會讓兩個 task 能根據加入的順序排序。因為沒有任何兩個 task 擁有相同的 entry count ，所以永遠不會直接使用 task 做比較。

// 不可比较任务问题的另一种解决方案是创建一个忽略任务条目并且只比较优先级字段的包装器类:
