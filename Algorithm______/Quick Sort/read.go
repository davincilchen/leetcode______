package main

//n*Logn  //每個點會被處理logn 次, 有n個點

// 效能
// Complexity
// Worst	O(n2)
// Best	O(nlogn)
// Average	O(nlogn)
// Worst space	O(logn) or O(n) auxiliary

// 最差情況
// 最差的分割序列狀況發生在挑選的 pivot 總是最大或最小值（或在 Lomuto partition 下，所有元素值都一樣）。由於 Lomuto 總是選擇最後一個元素作為 pivot，這種情形好發於已排序或接近排序完成的資料上。

// 而當每次的 partition 都是最不平衡的分割序列，就會產生最差時間複雜度的狀況。遞迴在序列長度等於 1 時停止，因此整個排序法的 call stack 需要 n−1 的嵌套遞迴呼叫（nested call）；而第 i 次分割會執行 n−i 次基本操作（ O(n)），所以總共需執行

// ∑i=0n(n−i)=n2−n(n+1)2
// 次基本操作，最差時間複雜度為 O(n2)。

//https://www.geeksforgeeks.org/quick-sort/

// Analysis of QuickSort
// Time taken by QuickSort, in general, can be written as following.

//  T(n) = T(k) + T(n-k-1) + \theta(n)

// The first two terms are for two recursive calls, the last term is for the partition process. k is the number of elements which are smaller than pivot.
// The time taken by QuickSort depends upon the input array and partition strategy. Following are three cases.

// Worst Case: The worst case occurs when the partition process always picks greatest or smallest element as pivot. If we consider above partition strategy where last element is always picked as pivot, the worst case would occur when the array is already sorted in increasing or decreasing order. Following is recurrence for worst case.

//  T(n) = T(0) + T(n-1) + \theta(n)

// which is equivalent to

// T(n) = T(n-1) + \theta(n)

// The solution of above recurrence is  \theta         (n2).

// Best Case: The best case occurs when the partition process always picks the middle element as pivot. Following is recurrence for best case.

//  T(n) = 2T(n/2) + \theta(n)

// The solution of above recurrence is \theta         (nLogn). It can be solved using case 2 of Master Theorem.

// Average Case:
// To do average case analysis, we need to consider all possible permutation of array and calculate time taken by every permutation which doesn’t look easy.
// We can get an idea of average case by considering the case when partition puts O(n/9) elements in one set and O(9n/10) elements in other set. Following is recurrence for this case.

//  T(n) = T(n/9) + T(9n/10) + \theta(n)

// Solution of above recurrence is also O(nLogn)
// Although the worst case time complexity of QuickSort is O(n2) which is more than many other sorting algorithms like Merge Sort and Heap Sort, QuickSort is faster in practice, because its inner loop can be efficiently implemented on most architectures, and in most real-world data. QuickSort can be implemented in different ways by changing the choice of pivot, so that the worst case rarely occurs for a given type of data. However, merge sort is generally considered better when data is huge and stored in external storage.

// Is QuickSort stable?
// The default implementation is not stable. However any sorting algorithm can be made stable by considering indexes as comparison parameter.

// Is QuickSort In-place?
// As per the broad definition of in-place algorithm it qualifies as an in-place sorting algorithm as it uses extra space only for storing recursive function calls but not for manipulating the input.

// What is 3-Way QuickSort?
// In simple QuickSort algorithm, we select an element as pivot, partition the array around pivot and recur for subarrays on left and right of pivot.
// Consider an array which has many redundant elements. For example, {1, 4, 2, 4, 2, 4, 1, 2, 4, 1, 2, 2, 2, 2, 4, 1, 4, 4, 4}. If 4 is picked as pivot in Simple QuickSort, we fix only one 4 and recursively process remaining occurrences. In 3 Way QuickSort, an array arr[l..r] is divided in 3 parts:
// a) arr[l..i] elements less than pivot.
// b) arr[i+1..j-1] elements equal to pivot.
// c) arr[j..r] elements greater than pivot.
// See this for implementation.

// How to implement QuickSort for Linked Lists?
// QuickSort on Singly Linked List
// QuickSort on Doubly Linked List

// Can we implement QuickSort Iteratively?
// Yes, please refer Iterative Quick Sort.

// Why Quick Sort is preferred over MergeSort for sorting Arrays
// Quick Sort in its general form is an in-place sort (i.e. it doesn’t require any extra storage) whereas merge sort requires O(N) extra storage, N denoting the array size which may be quite expensive. Allocating and de-allocating the extra space used for merge sort increases the running time of the algorithm. Comparing average complexity we find that both type of sorts have O(NlogN) average complexity but the constants differ. For arrays, merge sort loses due to the use of extra O(N) storage space.
// Most practical implementations of Quick Sort use randomized version. The randomized version has expected time complexity of O(nLogn). The worst case is possible in randomized version also, but worst case doesn’t occur for a particular pattern (like sorted array) and randomized Quick Sort works well in practice.
// Quick Sort is also a cache friendly sorting algorithm as it has good locality of reference when used for arrays.
// Quick Sort is also tail recursive, therefore tail call optimizations is done.

// Why MergeSort is preferred over QuickSort for Linked Lists?
// In case of linked lists the case is different mainly due to difference in memory allocation of arrays and linked lists. Unlike arrays, linked list nodes may not be adjacent in memory. Unlike array, in linked list, we can insert items in the middle in O(1) extra space and O(1) time. Therefore merge operation of merge sort can be implemented without extra space for linked lists.
// In arrays, we can do random access as elements are continuous in memory. Let us say we have an integer (4-byte) array A and let the address of A[0] be x then to access A[i], we can directly access the memory at (x + i*4). Unlike arrays, we can not do random access in linked list. Quick Sort requires a lot of this kind of access. In linked list to access i’th index, we have to travel each and every node from the head to i’th node as we don’t have continuous block of memory. Therefore, the overhead increases for quick sort. Merge sort accesses data sequentially and the need of random access is low.

// How to optimize QuickSort so that it takes O(Log n) extra space in worst case?
// Please see QuickSort Tail Call Optimization (Reducing worst case space to Log n )

// 快速排序所用的快速排序時間分析
// ，一般可以寫成如下。

// T(n) = T(k) + T(n-k-1) + \θ(n)

// 前兩項用於兩次遞歸調用，最後一項用於分區過程。k 是小於樞軸的元素的數量。
// QuickSort 花費的時間取決於輸入數組和分區策略。以下是三種情況。

// 最壞情況：最壞情況發生在分區過程總是選擇最大或最小元素作為樞軸時。如果我們考慮上面的分區策略，其中最後一個元素總是被選為樞軸，最壞的情況將發生在數組已經按升序或降序排序時。以下是最壞情況的重現。

// T(n) = T(0) + T(n-1) + \θ(n)

// 這相當於

// T(n) = T(n-1) + \θ(n)

// 上述遞歸的解是  \θ         (n 2 )。

// 最佳情況：當分區過程總是選擇中間元素作為樞軸時，就會出現最佳情況。以下是最佳情況的重現。

// T(n) = 2T(n/2) + \θ(n)

// 上述遞歸的解是 \θ         (nLogn)。可以使用Master Theorem的案例 2 來解決。

// 平均情況：
// 要進行平均情況分析，我們需要考慮數組的所有可能排列，併計算每個排列所花費的時間，這看起來並不容易。
// 我們可以通過考慮分區將 O(n/9) 個元素放在一個集合中並將 O(9n/10) 個元素放在另一個集合中的情況來了解平均情況。以下是此案例的重現。

// T(n) = T(n/9) + T(9n/10) + \θ(n)

// 上述遞歸的解也是 O(nLogn)
// 雖然 QuickSort 的最壞情況時間複雜度是 O(n 2 )，比許多其他排序算法如Merge Sort和Heap Sort都要多，但 QuickSort 在實踐中更快，因為它的內循環可以可以在大多數架構和大多數真實數據中有效實施。QuickSort 可以通過改變主元的選擇以不同的方式實現，因此對於給定類型的數據，最壞的情況很少發生。但是，當數據很大並且存儲在外部存儲中時，通常認為合併排序更好。

// 快速排序穩定嗎？
// 默認實現不穩定。然而，任何排序算法都可以通過將索引作為比較參數來穩定。

// 快速排序是否就地？
// 根據就地算法的廣泛定義，它有資格作為就地排序算法，因為它僅使用額外空間來存儲遞歸函數調用，但不用於操作輸入。

// 什麼是三向快速排序？
// 在簡單的快速排序算法中，我們選擇一個元素作為樞軸，圍繞樞軸對數組進行分區，並在樞軸的左側和右側遞歸子數組。
// 考慮一個具有許多冗餘元素的數組。例如，{1, 4, 2, 4, 2, 4, 1, 2, 4, 1, 2, 2, 2, 2, 4, 1, 4, 4, 4}。如果在 Simple QuickSort 中選擇 4 作為主元，我們只修復一個 4 並遞歸處理剩餘的事件。在 3 Way QuickSort 中，數組 arr[l..r] 分為 3 個部分：
// a) arr[l..i] 元素小於主元。
// b) arr[i+1..j-1] 個元素等於主元。
// c) arr[j..r] 元素大於樞軸。
// 請參閱此執行。

// 如何為鍊錶實現快速排序？
// 單鍊錶上的
// 快速排序 雙鍊錶上的快速排序

// 我們可以迭代地實現快速排序嗎？
// 是的，請參考迭代快速排序。

// 為什麼快速排序優於 MergeSort 用於對數組進行排序
// 快速排序的一般形式是就地排序（即它不需要任何額外的存儲空間），而合併排序需要 O(N) 額外的存儲空間，N 表示數組大小可能相當昂貴。分配和取消分配用於歸併排序的額外空間會增加算法的運行時間。比較平均複雜度，我們發現這兩種類型的平均複雜度為 O(NlogN)，但常數不同。對於數組，由於使用了額外的 O(N) 存儲空間，歸併排序會失敗。
// 快速排序的大多數實際實現都使用隨機版本。隨機版本的預期時間複雜度為 O(nLogn)。最壞的情況也可能出現在隨機版本中，但最壞的情況不會發生在特定模式（如排序數組）中，並且隨機快速排序在實踐中效果很好。
// 快速排序也是一種緩存友好的排序算法，因為它在用於數組時具有良好的引用局部性。
// 快速排序也是尾遞歸的，因此完成了尾調用優化。

// 為什麼 MergeSort 比 QuickSort 更適合鏈接列表？
// 在鍊錶的情況下，情況不同主要是由於數組和鍊錶的內存分配不同。
// 與數組不同，鍊錶節點在內存中可能不相鄰。與數組不同，在鍊錶中，
// 我們可以在 O(1) 額外空間和 O(1) 時間的中間插入項目。
// 因此可以實現歸併排序的歸併操作，而不需要額外的鍊錶空間。
// 在數組中，我們可以進行隨機訪問，因為元素在內存中是連續的。
// 假設我們有一個整數（4 字節）數組 A，讓 A[0] 的地址為 x，那麼要訪問 A[i]，
// 我們可以直接訪問 (x + i*4) 處的內存。與數組不同，我們不能在鍊錶中進行隨機訪問。
// 快速排序需要大量此類訪問。在鍊錶中訪問第 i 個索引，
// 我們必須將每個節點從頭到第 i 個節點，因為我們沒有連續的內存塊。
// 因此，快速排序的開銷會增加。
// 歸併排序按順序訪問數據，對隨機訪問的需求較低。

// Quicksort 是一種高效的排序算法，通常用於生產排序實現。
// 與合併排序一樣，快速排序是一種分治算法。顧名思義，
// Quicksort 是最快的排序算法之一，但你必須在實現中註意細節，
// 因為如果你不小心，你的速度會很快下降。

//https://qvault.io/golang/quick-sort-golang/
//時間複雜度計算
//https://qvault.io/golang/quick-sort-golang/

// Why use Quicksort?
// On average, quicksort has a Big O of O(n*log(n)). In the worst case, and assuming we don’t take any steps to protect ourselves, it can break down to O(n^2). The partition() function has a single for-loop that ranges from the lowest index to the highest index in the array. By itself, the partition() function is O(n). The overall complexity of quicksort is dependent on how many times partition() is called.

// In the worst case, the input is already sorted. An already sorted array results in the pivot being the largest or smallest element in the partition each time. When this is the case, partition() is called a total of n times. In the best case, the pivot is the middle element of each sublist which results in log(n) calls to partition().

// Quick sort has the following properties.

// Very fast in the average case
// In-Place: Saves on memory, doesn’t need to do a lot of copying and allocating
// More complex implementation
// Typically unstable: changes the relative order of elements with equal keys
// Ensuring a fast runtime in Quicksort
// While the version of quicksort that we implemented is almost always able to perform at speeds of O(n*log(n)), it’s Big O complexity is still technically O(n^2). We can fix this by altering the algorithm slightly. There are two approaches:

// Shuffle input randomly before sorting. This can trivially be done in O(n) time.
// Actively find the median of a sample of data from the partition, this can be done in O(1) time.
// Random shuffling optimization
// The random approach is easy to code, works practically all of the time, and as such is often used. The idea is to quickly shuffle the list before sorting it. The likelihood of shuffling into a sorted list is astronomically unlikely, and is also more unlikely the larger the input.

// Finding the median optimization
// One of the most popular solutions is to use the “median of three” approach. Three elements (for example: the first, middle, and last elements) of each partition are chosen and the median is found between them. That item is then used as the pivot. This approach has the advantage that it can’t break down to O(n^2) time because we are guaranteed to never use the worst item in the partition as the pivot. That said,
// it can still be slower because a true median isn’t used.

// 為什麼使用快速排序？
// 平均而言，快速排序的大 O 為O(n*log(n)). 在最壞的情況下，
// 假設我們不採取任何措施保護自己，它可能會崩潰O(n^2)。
// 該partition()函數有一個 for 循環，範圍從數組中的最低索引到最高索引。
// 就其本身而言，partition()函數是O(n)。
// 快速排序的整體複雜性取決於調用了多少次。 partition()

// 在最壞的情況下，輸入已經排序。
// 已排序的數組每次都會導致樞軸成為分區中的最大或最小元素。
// 在這種情況下，partition()稱為總n次數。在
// 最好的情況下，樞軸是每個子列表的中間元素，這會導致log(n)調用partition().
