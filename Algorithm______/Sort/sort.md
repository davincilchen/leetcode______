
<!-- https://clu.gitbook.io/data-structure-note/merge-sort-v.s.-quick-sort -->
<!-- https://www.geeksforgeeks.org/quick-sort-vs-merge-sort/ -->







* **2.1.6 Merge Sort v.s. Quick Sort**


<br>
1. 兩者其實非常相似, 都是把資料分成兩邊, 直到不能再分了, 才把資料合起來. 不過quick sort最大的特色就是會有partition的這個動作, 講白了就是把數字分好大小後再繼續往下分, 所以這樣循環下去分到最小後, 也表示完成了排序的動作了.


<br>
2. 在大部分的worst case下, merge sort是優於quick sort的, 再加上merge sort的worst case跟quick sort的best case之時間複雜度是一樣的, 這樣看來似乎是merge sort比較快(理論上).

<br>

3.可是實際上來說, 如果兩者都用遞迴的方式去實作的話, quick sort的method call為N, 那merge sort就會是2N-1(註1), 即merge sort多花了一倍的method call. 如果用迴圈來做, merge sort會花比較多時間在記憶體上面, 因為它不是in-place sorting.

<br>
4.不過merge sort有一個很好的特性: 它是穩定的(stable sorting), 穩定的意思是說, 排序前與排序後, 擁有相同key值的兩個資料, 彼此之間的順序是一樣的.

<br>
5.最後, 這兩者都是Divide and Conquer的做法, 只是quick sort為先苦後樂(遞迴之前的partition比較麻煩, 遞迴完後就沒事了); 而merge sort是先樂後苦(進入遞迴之前都沒事, 但是遞迴之後的合併動作就累了).
註1: 參閱原始碼, 注意merge function的step1~step2, 這是第一次的N, 即排序當前分割後的結果, 而step3, 把排序好的結果放回原本的array, 也會有一次N.



<br>
原地演算法 (In-place algorithm):
在電腦科學中，一個原地演算法是基本上不需要藉助額外的資料結構就能對輸入的資料進行變換的演算法。不過，分配少量空間給部分輔助變數是被允許的。演算法執行過程中，輸入的資料往往會被輸出結果覆蓋。原地演算法只能通過替換或交換元素的方式來修改原始的輸入。不滿足「原地」原則的演算法也被稱為非原地演算法或異地演算法