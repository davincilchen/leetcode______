package main

//Topic : Two Pointers, String
//Running Length

// Encode and Decode Strings
// Medium
// How many valid binary substrings exist in "000111", and how many in "11100"? What about "00011100"?

//https://www.cnblogs.com/grandyang/p/7716150.html
//https://zxi.mytechroad.com/blog/string/leetcode-696-count-binary-substrings/

// [LeetCode] Count Binary Substrings

// Give a string s, count the number of non-empty (contiguous) substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively.

// Substrings that occur multiple times are counted the number of times they occur.

// Example 1:

// Input: "00110011"
// Output: 6
// Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".

// Notice that some of these substrings repeat and are counted the number of times they occur.

// Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.

// Example 2:

// Input: "10101"
// Output: 4
// Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.

// Note:

// s.length will be between 1 and 50,000.
// s will only consist of "0" or "1" characters.

// 这道题给了我们一个二进制字符串，然后我们统计具有相同0和1的个数，
//且0和1各自都群组在一起(即0和1不能交替出现)的子字符串的个数，
//题目中的两个例子也很能说明问题。那么我们来分析题目中的第一个例子00110011，
//符合要求的子字符串要求0和1同时出现，那么当第一个1出现的时候，前面由于前面有两个0，
//所以肯定能组成01，
//再遇到下一个1时，此时1有2个，0有2个，能组成0011，
//下一个遇到0时，此时0的个数重置为1，而1的个数有两个，所以一定有10，
//同理，下一个还为0，就会有1100存在，之后的也是这样分析。

//那么我们可以发现我们只要分别统计0和1的个数，而且如果当前遇到的是1，
//那么只要之前统计的0的个数大于当前1的个数，就一定有一个对应的子字符串，
//而一旦前一个数字和当前的数字不一样的时候，那么当前数字的计数要重置为1。
//所以我们遍历元数组，如果是第一个数字，那么对应的ones或zeros自增1。
//然后进行分情况讨论，如果当前数字是1，然后判断如果前面的数字也是1，则ones自增1，
//否则ones重置为1。如果此时zeros大于ones，res自增1。反之同理，如果当前数字是0，
//然后判断如果前面的数字也是0，则zeros自增1，否则zeros重置为1。
//如果此时ones大于zeros，res自增1。参见代码如下：

// 解法一：

// 复制代码
// class Solution {
// public:
//     int countBinarySubstrings(string s) {
//         int zeros = 0, ones = 0, res = 0;
//         for (int i = 0; i < s.size(); ++i) {
//             if (i == 0) {
//                 (s[i] == '1') ? ++ones : ++zeros;
//             } else {
//                 if (s[i] == '1') {
//                     ones = (s[i - 1] == '1') ? ones + 1 : 1;
//                     if (zeros >= ones) ++res;
//                 } else if (s[i] == '0') {
//                     zeros = (s[i - 1] == '0') ? zeros + 1 : 1;
//                     if (ones >= zeros) ++res;
//                 }
//             }
//         }
//         return res;
//     }
// };
// 复制代码

// 下面这种方法更加简洁了，不用具体的分0和1的情况来讨论了，而是直接用了pre和cur两个变量，其中pre初始化为0，cur初始化为1，然后从第二个数字开始遍历，如果当前数字和前面的数字相同，则cur自增1，否则pre赋值为cur，cur重置1。然后判断如果pre大于等于cur，res自增1。其实核心思想跟上面的方法一样，只不过pre和cur可以在0和1之间切换，参见代码如下：

// 解法二：

// 复制代码
// class Solution {
// public:
//     int countBinarySubstrings(string s) {
//         int res = 0, pre = 0, cur = 1, n = s.size();
//         for (int i = 1; i < n; ++i) {
//             if (s[i] == s[i - 1]) ++cur;
//             else {
//                 pre = cur;
//                 cur = 1;
//             }
//             if (pre >= cur) ++res;
//         }
//         return res;
//     }
// };

// ===== //
// Give a string s, count the number of non-empty (contiguous) substrings that have the same number of 0’s and 1’s, and all the 0’s and all the 1’s in these substrings are grouped consecutively.

// Substrings that occur multiple times are counted the number of times they occur.

// Example 1:

// Input: "00110011"
// Output: 6
// Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".

// Notice that some of these substrings repeat and are counted the number of times they occur.

// Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.
// Example 2:

// Input: "10101"
// Output: 4
// Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.
// Note:

// s.length will be between 1 and 50,000.
// s will only consist of “0” or “1” characters.
// Solution 0: Search

// Try all possible substrings and check whether it’s a valid one or not.

// Time complexity O(2^n * n) TLE

// Space complexity: O(n)

// Solution 1: Running Length

// For S = “000110”, there are two virtual blocks “[00011]0” and “000[110]” which contains consecutive 0s and 1s or (1s and 0s)

// Keep tracking of the running length of 0, 1 for each block

// “00011” => ‘0’: 3, ‘1’:2

// ans += min(3, 2) => ans = 2 (“[0011]”, “0[01]1”)

// “110” => ‘0’: 1, ‘1’: 2

// ans += min(1, 2) => ans = 3 (“[0011]”, “0[01]1”, “1[10]”)

// We can reuse part of the running length from last block.

// Time complexity: O(n)

// Space complexity: O(1)

// Author: Huahua
// // Running time: 43 ms
// class Solution {
// public:
//   int countBinarySubstrings(string s) {
//     vector<int> lens(128, 0);
//     int i = 0;
//     int l = 0;
//     int ans = 0;
//     while (true) {
//       while (i < s.length() && s[i] == s[l]) ++i;
//       lens[s[l]] = i - l;
//       ans += min(lens['0'], lens['1']);
//       if (i == s.length()) break;
//       lens[s[i]] = 0;
//       l = i;
//     }
//     return ans;
//   }
// };
