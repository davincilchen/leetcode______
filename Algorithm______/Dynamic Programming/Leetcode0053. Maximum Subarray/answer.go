package main

//https://ithelp.ithome.com.tw/articles/10213270
//https://medium.com/the-research-nest/solving-the-maximum-subarray-sum-a-super-simplified-explanation-34042ffd3fd7
//https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/

// 「假設我知道nums中1到i-1的最大子陣列和，
// 有沒有辦法知道1到i的最大子陣列和呢？」
// 我們討論1到i的最大子陣列和，應該分兩種情況：
// a. 這個子陣列包括index i
// b. 這個子陣列不包括index i
// 如果是b的話，那答案其實就是1到i-1的最大子陣列和；
// 但如果是a的話，我們好像還需要一些條件？
// 什麼條件呢？就是1到i-1中，包含i-1元素的最大子陣列和。
// 因為如果是前面所提到的a狀況的話，
// 那麼1到i的最大連續子陣列，就必須從i->i-1連回去才行(或乾脆是只有i)。

// 所以這麼一來，條件就比較清楚了，我們會需要的有：

// 到現在為止包含當前這個元素的最大子陣列和，我們命名為curr(current)
// 到目前為止的最大子陣列和，我們命名為res(result)
// 那麼我們就以範例的[-2,1,-3,4,-1,2,1,-5,4]來操作看看。
// 該怎麼決定curr跟res的初始值呢？因為一開始只有一個值，
// 所以curr和res應該都等於nums[0]=-2才對。
// 所以一開始我們將curr和res都設為-2，接下來由i=1開始，
// 一路往下到最後一個結束。

// i=1: [-2,1,-3,4,-1,2,1,-5,4]
// 按我們的定義，curr應該是-2+1和1中較大的值，
// 我們可以先將curr加上nums[1]，
// 接下來檢查curr小於0，或者curr<nums[1]，
// 就表示其實我們只要取nums[1]就好。
// (註1: 表示兩者的正負號為(-, -), (-, +)，不論如何前面那部分貢獻均是負的)
// (註2: 也可以拆成別的形式或調換順序，但先判斷<0的計算速度會比較快)
// (註3: 後面程式碼Java和Python的判斷方式就不同，可自行依喜好決定使用)
// 所以新的curr為1, 接著再比較得知curr<res，故res也要改為1。
// 接下來的流程，讀者可以嘗試自己操作一遍再對照一下。

// i=2: [-2,1,-3,4,-1,2,1,-5,4]
// -> curr = 1 + -3 = -2 < 0
// -> curr = -3
// -> res = 1

// i=3: [-2,1,-3,4,-1,2,1,-5,4]
// -> curr = -3 + 4 = 1 < 4
// -> curr = 4
// -> res = 4

// i=4: [-2,1,-3,4,-1,2,1,-5,4]
// -> curr = 4 + -1 = 3 > -1

// i=5: [-2,1,-3,4,-1,2,1,-5,4]
// -> curr = 3 + 2 = 5 > 2
// -> res = 5

// i=6: [-2,1,-3,4,-1,2,1,-5,4]
// -> curr = 5 + 1 = 6 > 1
// -> res = 6

// i=7: [-2,1,-3,4,-1,2,1,-5,4]
// -> curr = 6 + -5 = 1 > -5

// i=8: [-2,1,-3,4,-1,2,1,-5,4]
// -> curr = 1 + 4 = 5 > 4
// 所以最後我們得到res = 6，再回傳即可。

// 這個解法第一次並不是很容易上手，而且並非相當直觀的方式，
// 但它常常可以在特別的地方有效地解決問題。
// 適用於這類型的解法的問題會有一個特性：
// 「N=i+1時的解，通常和N=i時的解以及第i+1的值有關連性」
// (也有可能跟更前面幾個值同時有關連性)
// 而我們通常很難直觀地直接用簡單的方式去算出N=i+1的解，
// 所以會依靠其本身和前面的連動關係，
// 讓答案像堆積木一樣，從i=0開始堆到最後面得到最終的答案。

// 有些問題可能未必是一層一層疊上去，
// 也有可能是把一個大問題拆成數個小問題，
// 最後將其組合起來得到大問題的答案。
// 但通常我們最常見的形式還是如上面那樣子按順序堆疊的。

// 上面這樣子形式的解題方法，
// 被稱之為動態規劃(Dynamic Programming)，或者也常簡稱為DP。
// 每個動態規劃的問題都會有些不一樣，
// 但只要抓到其中的連動性，找到從i -> i+1中間的關聯，
// 就可以順利解開題目。

// // Javascript solution 2
// // Input is an array of integers of length n
// var maxSubArray = function(array) {
//     let currentSum = array[0];
//     let maxSum = currentSum;
//     const n = array.length;
//     for (let i = 1; i < n; i++) {
//         currentSum = Math.max(array[i], currentSum + array[i]);
//         maxSum = Math.max(maxSum, currentSum);
//     }
//     return maxSum;
// };

// # Python solution 2
// # Input is an arrya of integers of length n
// def maxSubArrayV2(arr):
//   currentSum = arr[0]     #Starting with the first element
//   maxSum = currentSum
//   for i in range(1, len(arr)):
//     currentSum = max(arr[i], currentSum + arr[i])     #Updating current sum if the current element adds more sum value
//     maxSum = max(maxSum, currentSum)                  #Updating maxSum
//   return maxSum

// C++ program to print largest contiguous array sum
// #include<iostream>
// #include<climits>
// using namespace std;

// int maxSubArraySum(int a[], int size)
// {
// 	int max_so_far = INT_MIN, max_ending_here = 0,
// 	start =0, end = 0, s=0;

// 	for (int i=0; i< size; i++ )
// 	{
// 		max_ending_here += a[i];

// 		if (max_so_far < max_ending_here)
// 		{
// 			max_so_far = max_ending_here;
// 			start = s;
// 			end = i;
// 		}

// 		if (max_ending_here < 0)
// 		{
// 			max_ending_here = 0;
// 			s = i + 1;
// 		}
// 	}
// 	cout << "Maximum contiguous sum is "
// 		<< max_so_far << endl;
// 	cout << "Starting index "<< start
// 		<< endl << "Ending index "<< end << endl;
// }

// /*Driver program to test maxSubArraySum*/
// int main()
// {
// 	int a[] = {-2, -3, 4, -1, -2, 1, 5, -3};
// 	int n = sizeof(a)/sizeof(a[0]);
// 	int max_sum = maxSubArraySum(a, n);
// 	return 0;
// }

// 一定有一些我們可以使用的技巧來做到這一點。讓我們看看我們是否可以構建一個思維過程。
// 如果我們只循環一次數組，這意味著我們訪問數組的每個元素一次。
// 現在讓我們將這些點向後連接。為了找到我們的答案，我們首先需要找到總和最大的子數組。
// 我們肯定需要訪問每個元素一次。你能告訴那個元素到最大和的子數組嗎？它要么屬於那個子數組，要么不屬於。正確的？
// 如果它確實屬於具有最大總和的子數組，您可以將其值添加到您存儲的臨時變量中並不斷更新潛在的最大總和。如果不是，您可以簡單地忽略它並移至下一個元素。
// 如果你在整個陣列上這樣做，你會得到什麼？屬於具有最大和的子數組的所有元素的總和 = 我們正在尋找的答案，我們只需運行一個循環就可以得到它！
// 現在您將遇到的問題是，我們如何知道給定元素屬於該子數組？
// 讓我們假設一個包含n 個元素的數組。我們在這個數組的第 k 個元素。
// 如果這個第 k 個元素是我們最大子數組的一部分，它會有什麼特別之處？它要么大於元素的總和，要么直到+元素k-1的最大總和更大。k-1kth
// 你能注意到我們是如何用我們的邏輯創建一個子問題的嗎？
// 此數組在第 k 個元素處結束的最大子數組和實際上是數組的最大子數組和，直到第 k-1 個元素 + 第 k 個元素（如果第 k 個元素為正）。
// 在任何時刻，我們都在找到一個數組的最大子數組和，直到第 k 個元素。當我們到達第 n 個元素時，我們會找到完整數組的答案。
// 聽起來很混亂？沒關係。如果這是您第一次遇到這樣的問題，可能會有點複雜。讓我們用一個真實的數組來試運行。
// 讓數組[-1, 2, -5, 7]與我們在解決方案 1 中使用的示例一樣。因此，我們需要將屬於具有最大總和的子數組的元素總和（讓我們稱之為目標子數組）。
// 我們只運行一個循環。所以，首先，我們去，-1它本身就是一個帶有 sum = 的子數組-1。讓我們假設-1它是我們的目標子數組的一部分。-1現在是我們暫時的最大金額。
// 接下來，我們去2。現在，它要么是我們目標子數組的一部分，要么不是。我們怎麼知道？如果它是子數組的一部分，則它應該大於當前最大和或添加到最大和。2 > -1 + 2 = 1，這大於我們當前的臨時最大總和。所以，讓我們更新我們的臨時最大總和 = 2。這只是證實了我們之前作為目標數組一部分的假設-1是錯誤的。現在，我們假設2是目標數組的一部分。
// 讓我們轉到下一個元素，即-5。-5顯然小於我們當前的最大總和。因此，我們可以忽略它並轉到下一個元素。
// 現在，讓我們轉到下一個元素，7它大於當前的 max sum = 2。7 + 2 = 9這也大於我們當前的最大總和。這意味著7是我們目標子數組的一部分，我們可以更新我們的最大 sum = 9。
// 我們已經到達數組的末尾，我們得到的解決方案是9，這是所需的答案，我們只使用了一個循環！
// 簡而言之，我們檢查了數組的每個元素，並檢查它是否會成為具有最大和的子數組的一部分。如果是這樣，我們只是將它添加到我們的 maxSum 變量並繼續到下一個元素。
