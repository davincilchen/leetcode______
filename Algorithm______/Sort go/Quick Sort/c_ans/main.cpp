
#include <iostream>



//http://alrightchiu.github.io/SecondRound/comparison-sort-quick-sortkuai-su-pai-xu-fa.html


//https://rust-algo.club/sorting/quicksort/index.html

//簡單的最佳化實作下，Quicksort 僅需 O(logn) 的額外儲存空間，
//比它的競爭對手 mergesort 來得節省。非常適合運用在真實世界中的排序法。


//======
// Space complexity
// Quicksort 的空間複雜度取決於實作細節，由於分割序列步驟需 O(1) 的空間複雜度，因此僅需分析遞迴式會在 call stack 產生多少 stack frame 即可。

// 前面提及，最 naïve 的 Lomuto partition 最糟糕的情形下，會產生 n−1 個嵌套遞迴，也就是需額外使用 O(n) 的空間儲存 call stack frame，但只要 compiler 有支援尾端呼叫最佳化（tail-call optimization，TCO），Quicksort 很容易最佳化至 O(logn)。

//===============
// int front為數列的「最前端」index。
// 此例，front為index(0)。
// int end為數列的「最尾端」index。
// 此例，end為index(8)。
// int i為所有小於pivot的數所形成的數列的「最後位置」。
// 一開始將index(i)初始化為front-1，因為有可能數列中，所有數都比pivot大。
// 一旦發現有數比pivot小，index(i)便往後移動(i++)，表示「所有小於pivot的數所形成的數列」又增加了一名成員。
// 當演算法結束時，所有在index(i)左邊的數，都比pivot小，所有在index(i)右邊的數，都比pivot大。
// int j是讓pivot與其餘數值逐一比較的index，從front檢查到end-1(因為end是pivot自己)。
// int pivot=array[end]，以數列最後一個數做為pivot。
// 此例，pivot為5。
// pivot的「數值」可以是任意的，挑選矩陣的最後一個位置是為了方便index(j)的移動，也可以挑選任意位置。


void swap(int *a, int *b){
    int temp = *a;
    *a = *b;
    *b = temp;
}



//比較大的不動等人來換
int Partition1(int *arr, int front, int end){
    std::cout <<  "Partition : " << front << " - "<< end << std::endl;
    int pivot = arr[end];
    int i = front -1;
    for (int j = front; j < end; j++) {
        if (arr[j] < pivot) {
            i++;
            swap(&arr[i], &arr[j]);
        }
    }
    i++;
    swap(&arr[i], &arr[end]); //switch pivot
    return i;
}

//error
int Partition2(int *arr, int front, int end){
    std::cout <<  "Partition : " << front << " - "<< end << std::endl;
    int pivot = arr[end];
    int i = front;//front -1;
    for (int j = front; j < end; j++) {
        if (arr[j] < pivot) {
            //i++;
            swap(&arr[i], &arr[j]);
            i++;
        }
    }
    //i++;
    swap(&arr[i], &arr[end]); //switch pivot
    //i++;
    return i;
}

int Partition(int *arr, int front, int end) {
    //return Partition1(arr,  front,  end);
    return Partition2(arr,  front,  end);
}

void QuickSort(int *arr, int front, int end){
    std::cout <<  front << " - "<< end << std::endl;
    if (front < end) {        
        int pivot = Partition(arr, front, end);
        QuickSort(arr, front, pivot - 1);
        QuickSort(arr, pivot + 1, end);
    }
}
void PrintArray(int *arr, int size){
    for (int i = 0; i < size; i++) {
        std::cout << arr[i] << " ";
    }
    std::cout << std::endl;
}
int main() {

    int n = 9;
    int arr[] = {9, 4, 1, 6, 7, 3, 8, 2, 5};
    std::cout << "original:\n";
    PrintArray(arr, n);

    QuickSort(arr, 0, n-1);

    std::cout << "sorted:\n";
    PrintArray(arr, n);
    return 0;
}