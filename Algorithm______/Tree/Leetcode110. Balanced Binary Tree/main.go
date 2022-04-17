package main

//https ://https://leetcode.com/problems/balanced-binary-tree/
// Given a binary tree, determine if it is height-balanced.

// For this problem, a height-balanced binary tree is defined as:

// a binary tree in which the left and right subtrees of every node differ in height by no more than 1.

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: true
// Example 2:

// Input: root = [1,2,2,3,3,null,null,4,4]
// Output: false
// Example 3:

// Input: root = []
// Output: true

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  isBalanced2 比較慢
//  isBalanced 比較快

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(tree []int) bool {
	root := buildTree_slice(tree)
	return isBalanced(root)
}

func isBalanced(root *TreeNode) bool {
	ok, _ := isBalancedDeep(root)
	return ok
}

func isBalancedDeep(root *TreeNode) (bool, int) {

	if root == nil {
		return true, 0
	}

	bL, left := isBalancedDeep(root.Left)
	if !bL {
		return false, -1 //-1 for don't care
	}
	bR, right := isBalancedDeep(root.Right)
	if !bR {
		return false, -1 //-1 for don't care
	}

	check := left - right
	if check <= -2 || check >= 2 {
		return false, -1 //-1 for don't care
	}
	return true, max(left, right) + 1
}

func IsBalanced2(tree []int) bool {
	root := buildTree_slice(tree)
	return isBalanced2(root)
}

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := depth(root.Left)
	right := depth(root.Right)
	check := left - right
	if check <= -2 || check >= 2 {
		return false
	}

	if root.Left != nil {
		ok := isBalanced(root.Left)
		if !ok {
			return false
		}
	}

	if root.Right != nil {
		ok := isBalanced(root.Right)
		if !ok {
			return false
		}
	}

	return true
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := depth(root.Left)
	right := depth(root.Right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {

}

func buildTree_slice(tree []int) *TreeNode {
	if len(tree) == 0 {
		return nil
	}

	cur := 0
	root := &TreeNode{Val: tree[cur]}
	cur++

	tlist := []*TreeNode{}
	tlist = append(tlist, root)
	for cur < len(tree) {
		v1 := tree[cur]
		cur++
		v2 := -1
		if cur < len(tree) {
			v2 = tree[cur]
		}
		cur++
		if len(tlist) <= 0 {
			break
		}

		node := tlist[0]

		tlist = tlist[1:]

		if v1 != -1 {
			node.Left = &TreeNode{
				Val: v1,
			}
			tlist = append(tlist, node.Left)
		}
		if v2 != -1 {
			node.Right = &TreeNode{
				Val: v2,
			}
			tlist = append(tlist, node.Right)
		}
	}

	return root

}
