package main

//https://zxi.mytechroad.com/blog/string/leetcode-1541-minimum-insertions-to-balance-a-parentheses-string/
//https://www.geeksforgeeks.org/minimum-number-of-parentheses-to-be-added-to-make-it-valid/

//平衡是()
// Recommended: Please try your approach on {IDE} first, before moving on to the solution.
// Approach: We keep the track of balance of the string i:e the number of ‘(‘ minus the number of ‘)’. A string is valid if its balance is 0, and also every prefix has non-negative balance.
// Now, consider the balance of every prefix of S. If it is ever negative (say, -1), we must add a ‘(‘ bracket at the beginning. Also, if the balance of S is positive (say, +P), we must add P times ‘)’ brackets at the end.
// Below is the implementation of the above approach:

// class Solution {
// 	public:
// 		int minAddToMakeValid(string s) {
// 			 // maintain balance of string
// 			int bal = 0;
// 			int ans = 0;

// 			for (int i = 0; i < s.length(); ++i) {

// 				bal += s[i] == '(' ? 1 : -1;

// 				// It is guaranteed bal >= -1
// 				if (bal == -1) {
// 					ans += 1;
// 					bal += 1;
// 				}
// 			}

// 			return bal + ans;
// 		}
// 	};

//平衡是()
// Solution: Counting
// Count how many close parentheses we need.

// if s[i] is ‘)’, we decrease the counter.
// if counter becomes negative, means we need to insert ‘(‘
// increase ans by 1, increase the counter by 2, we need one more ‘)’
// ‘)’ -> ‘()’
// if s[i] is ‘(‘
// if we have an odd counter, means there is a unbalanced ‘)’ e.g. ‘(()(‘, counter is 3
// need to insert ‘)’, decrease counter, increase ans
// ‘(()(‘ -> ‘(())(‘, counter = 2
// increase counter by 2, each ‘(‘ needs two ‘)’s. ‘(())(‘ -> counter = 4
// Once done, if counter is greater than zero, we need insert that much ‘)s’
// counter = 5, ‘((()’ -> ‘((())))))’
// Time complexity: O(n)
// Space complexity: O(1)

// class Solution {
// 	public:
// 	  int minInsertions(string s) {
// 		int ans = 0;
// 		int close = 0; // # of ')' needed.
// 		for (char c : s) {
// 		  if (c == ')') {
// 			if (--close < 0) {
// 			  // need to insert one '('
// 			  // ')' -> '()'
// 			  ++ans;
// 			  close += 2;
// 			}
// 		  } else {
// 			if (close & 1) {
// 			  // need to insert one ')'
// 			  // case '(()(' -> '(())('
// 			  --close;
// 			  ++ans;
// 			}
// 			close += 2; // need two ')'s
// 		  }
// 		}
// 		return ans + close;
// 	  }
// 	};
