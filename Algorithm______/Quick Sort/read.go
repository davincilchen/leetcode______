package main

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
