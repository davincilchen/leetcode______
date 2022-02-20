package main

//https://ithelp.ithome.com.tw/articles/10213270
//https://medium.com/the-research-nest/solving-the-maximum-subarray-sum-a-super-simplified-explanation-34042ffd3fd7
//

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
