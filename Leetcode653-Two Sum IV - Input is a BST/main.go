package main

import "fmt"

//https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
// Given the root of a Binary Search Tree and a target number k,
//return true if there exist two elements in the BST
//such that their sum is equal to the given target.

// Example 1:

// Input: root = [5,3,6,2,4,null,7], k = 9
// Output: true
// Example 2:

// Input: root = [5,3,6,2,4,null,7], k = 28
// Output: false

// Constraints:

// The number of nodes in the tree is in the range [1, 104].
// -104 <= Node.val <= 104
// root is guaranteed to be a valid binary search tree.
// -105 <= k <= 105

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//BST
//https://pjchender.dev/dsa/dsa-bst/

// func printPreOrder(n *Node) {
// 	if n == nil {
// 		return
// 	} else {
// 		fmt.Printf("%d ", n.data)
// 		printPreOrder(n.left)
// 		printPreOrder(n.right)
// 	}
// }

// func printInOrder(n *Node) {
// 	if n == nil {
// 		return
// 	} else {
// 		printPostOrder(n.left)
// 		fmt.Printf("%c ", n.data)
// 		printPostOrder(n.right)
// 	}
// }

// func printPostOrder(n *Node) {
// 	if n == nil {
// 		return
// 	} else {
// 		printPostOrder(n.left)
// 		printPostOrder(n.right)
// 		fmt.Printf("%c ", n.data)
// 	}
// }

//https://magiclen.org/arithmetic/
// 前序：運算子 運算元 運算元
// 中序：運算元 運算子 運算元
// 後序：運算元 運算元 運算子

// Runtime: 24 ms, faster than 78.03% of Go online submissions for Two Sum IV - Input is a BST.
//Memory Usage: 7.9 MB, less than 28.03% of Go online submissions for Two Sum IV - Input is a BST.

//因為是BST 有些狀況還可以加快
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]int)
	return findTargetInTree(root, m, k)
}

func findTargetInTree(root *TreeNode, m map[int]int, k int) bool {

	if root == nil {
		return false
	}

	_, ok := m[k-root.Val]
	if ok {
		return true
	}
	m[root.Val] = root.Val

	if root.Left != nil {
		ret := findTargetInTree(root.Left, m, k)
		if ret {
			return ret
		}
	}

	if root.Right != nil {
		ret := findTargetInTree(root.Right, m, k)
		if ret {
			return ret
		}
	}

	return false
}

func main() {
	t1 := &TreeNode{
		Val: 5,
	}
	t1.Left = &TreeNode{
		Val: 3,
	}
	t1.Right = &TreeNode{
		Val: 6,
	}
	tt := t1.Left
	tt.Left = &TreeNode{
		Val: 2,
	}
	tt.Right = &TreeNode{
		Val: 4,
	}

	tt = t1.Right
	tt.Right = &TreeNode{
		Val: 7,
	}
	target := 9
	ret := findTarget(t1, target)
	fmt.Println("findTarget = ", target, ret)

	target = 28
	ret = findTarget(t1, target)
	fmt.Println("findTarget =", target, ret)

}
