package main

//https://www.hackerrank.com/challenges/string-function-calculation/topics/two-pointer-technique

// Two Pointer Technique
// The 2 pointer technique is mostly applicable in sorted arrays where we try to perform a search in O(N).

// Here's an example.
// Given two arrays ( and ) sorted in ascending order and an integer , we need to find  and , such that  is equal to .

//  and  are our pointers, at first step  is pointing to the first element of , and  points to the last element of .

// i = 0; j = b.size() - 1;
// Move the first pointer one by one through the array , and correct position of the second pointer if needed

// while (i < a.size()) {
//     while(a[i] + b[j] > X && j > 0) j--;
//     if (a[i] + b[j] == X) writeAnswer(i, j);
//     i++;
// }
