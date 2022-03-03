package main

//https://medium.com/@Kadai/%E8%B3%87%E6%96%99%E7%B5%90%E6%A7%8B%E5%A4%A7%E4%BE%BF%E7%95%B6-binary-heap-ec47ca7aebac
//https://codebbkaf.medium.com/stack-vs-heap-9ebd44381387
//https://medium.com/@zector1030/memo-heap-92da1a10e8f1
//https://ithelp.ithome.com.tw/articles/10250174
//https://rust-algo.club/sorting/heapsort/
//https://ithelp.ithome.com.tw/articles/10214861
//https://golangbyexample.com/heapsort-in-golang/
//https://leetcode.com/problems/k-closest-points-to-origin/discuss/1457945/go-solution-min-heap
//https://docs.python.org/zh-tw/3/library/heapq.html
//heappushpop 將 item 放入 heap ，接著從 heap 取出並回傳最小的元素。這個組合函式比呼叫 heappush() 之後呼叫

// 堆積排序 Heapsort
// Heapsort（堆積排序）可以看作是 selection sort 的變形，同樣會將資料分為 sorted pile 與 unsorted pile，並在 unsorted pile 中尋找最大值（或最小值），加入 sorted pile 中。

// 和 selection sort 不同之處是，heapsort 利用堆積（heap）這種半排序（partially sorted）的資料結構輔助並加速排序。

// Heapsort 的特性如下：

// 使用 heap 資料結構輔助，通常使用 binary heap。
// 不穩定排序：排序後，相同鍵值的元素相對位置可能改變。
// 原地排序：不需額外花費儲存空間來排序。
// 較差的 CPU 快取：heap 不連續存取位址的特性，不利於 CPU 快取。
// 快取連續存取位址

// 效能
// Complexity
// Worst O(nlogn)
// Best	O(nlogn)
// Average	O(nlogn)
// Worst space	O(logn) or O(n) auxiliary

//Heapify
//⌊n/2⌋⋅O(logn)=O(nlogn)

//實際上，build heap 步驟的複雜度可達到 O(n)，可以看看 UMD 演算法課程 Lecture note 的分析。
