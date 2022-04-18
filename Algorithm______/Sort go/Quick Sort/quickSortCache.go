package main

//search:	 quick sort 局部

// 如果要对适合缓存的数组进行排序,那么quicksort将需要更少的内存访问,
// 因为mergesort需要分配第二个数组.

// Quicksort会将数组加载到缓存中,然后在不等待内存的情况下继续运行.
// Mergesort将支付访问第二个阵列的额外费用.

// 如果你选的是不适合高速缓存的数组,
// 然后快速排序仍从一个地方点获胜,
// 因为他们改乘更小的部分进行排序,这两种算法很快就会到的部分不放入高速缓存,

// 以及通过上述论点,这些快速排序更快.在不适合缓存的排序的较高级别上,
// quicksort和mergesort从缓存局部性的角度来看几乎等效.

//=======

//https://www.jianshu.com/p/250cec3ba09a

// Quick Sort & Merge Sort

// 40巨盗
// 0.089
// 2018.08.11 22:01:48
// 字数 289
// 阅读 355
// Merge Sort & Quick Sort，这两个排序算法都是利用Divide & Conquer最经典的例子。Merge Sort是先局部有序再整体有序，而Quick Sort是先整体有序再局部有序。由于Merge Sort需要一个拷贝数组的过程，所以速度不及Quick Sort。但两种排序算法中的思想都是非常重要的，在很多题中都会用到，所以在此提及。
// Merge Sort: 由于是先局部有序再整体有序，所以要先调用两次mergeSort()之后再调用merge()将已排序的两个子数组合并。还需要注意需要一个辅助数组aux[]以及在merge时，对一个数组已经结束时的处理。
// Quick Sort: 由于是先整体有序再局部有序，所以要先调用partision()根据pivot将原数组化为两个字数组，在调用两次quickSort()对子数组进行排序。我们默认left指针所对应的元素即为pivot元素，注意下标的处理。

// Merge Sort代码如下：

// class Solution:
//     """
//     @param A: an integer array
//     @return: nothing
//     """
//     def sortIntegers2(self, A):
//         aux = self.mergeSort(A)
//         for i in range(len(A)):
//             A[i] = aux[i]

//     def mergeSort(self, A):
//         result = []
//         if len(A) < 2:
//             return A
//         mid = int(len(A) / 2)
//         left = self.mergeSort(A[:mid])
//         right = self.mergeSort(A[mid:])
//         i, j = 0, 0
//         while i < len(left) and j < len(right):
//             if left[i] < right[j]:
//                 result.append(left[i])
//                 i += 1
//             else:
//                 result.append(right[j])
//                 j += 1
//         result += left[i:]
//         result += right[j:]
//         return result
// Quick Sort代码如下：

// class Solution:
//     """
//     @param A: an integer array
//     @return: nothing
//     """
//     def sortIntegers2(self, A):
//         self.quickSort(A, 0, len(A) - 1)

//     def quickSort(self, A, left, right):
//         if left >= right:
//             return
//         low = left
//         high = right
//         key = A[low]
//         while left < right:
//             while left < right and A[right] > key:
//                 right -= 1
//             while left < right and A[left] <= key:
//                 left += 1
//             A[left], A[right] = A[right], A[left]
//         A[right], A[low] = A[low], A[right]
//         self.quickSort(A, low, left - 1)
//         self.quickSort(A, left + 1, high)

//==================
//https://www.yikanggao.com/blog/2017/04/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F%E7%9A%84%E7%BB%86%E8%8A%82.html
// 快速排序的细节
//  2017-04-12 | 阅读：次
// 声明：本博客由 FrancisGeek原创，如需使用请在开头引用或者添加转载字样，谢谢配合。同时也仅代表个人观点。

// 目录

// 快速排序的直观理解
// Code:
// 快速排序用到的技巧
// 细节漫谈
// 真正理解了快速排序的思想了吗
// 快速排序的直观理解
// 快速排序是非常基础并且有用的排序算法，其parttion思想被众多算法借鉴．同时也是分治思想的体系，先整体有序再局部有序,是一个in-place排序算法，但是不是稳定排序算法.时间复杂度O(nlgn),空间复杂度O(1).. 但是要写出bug-free的快速排序，还是有很多细节需要注意．

// Code:
// 	public class Solution {
// 	    /**
// 	     * @param A an integer array
// 	     * @return void
// 	     */
// 	    public void sortIntegers2(int[] A) {
// 		// Write your code here
// 		if (A == null || A.length == 0){
// 		    return;
// 		}
// 		quickSort(A, 0, A.length -1);
// 	    }
// 	    private void quickSort(int[] nums, int start, int end){
// 		if (start >= end){
// 		    return;
// 		}
// 		int pivot = nums[(start + end) / 2];
// 		int left = start;
// 		int right = end;

// 		while (left <= right){
// 		    while (left <= right && nums[left] < pivot){
// 		        left++;
// 		    }
// 		    while (left <= right && nums[right] > pivot){
// 		        right--;
// 		    }
// 		    if (left <= right){
// 		        int tmp = nums[left];
// 		        nums[left] = nums[right];
// 		        nums[right] = tmp;
// 		        left++;
// 		        right--;
// 		    }
// 		}
// 		quickSort(nums, start, right);
// 		quickSort(nums, left, end);
// 	    }
// 	}

// 快速排序用到的技巧
// 1.递归 2.双指针(反向) 3.分治

// 细节漫谈
// 快速排序是先整体有序再局部有序，所以先parttion再分治递归．左右指针的循条件为left <= right,目的是让left和right出一次循环的时候没有交集，以免StackOverflow…在与pivot比较的时候，用</>,实现左右元素的均分，对付像[1, 1, 1, 1, 1,1]这种情况． 由于切分操作完毕后，left和right指针已经交错了，所以左边是right,右边是left.

// 真正理解了快速排序的思想了吗
// 有一个和快速排序的思想的算法,quick select(find the Kth largest(smallest) element),如果能够bug-free写出来，基本也就掌握了快速排序的思想．
