package main

//https://leetcode.com/problems/binary-tree-inorder-traversal/
// Input: root = [1,null,2,3]
// Output: [1,3,2]
// Example 2:

// Input: root = []
// Output: []
// Example 3:

// Input: root = [1]
// Output: [1]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Pre-order 前序：root → left → right
// In-order 中序：left → root → right
// Post-order 後序：left → right → root
// Level-order 層序：一層一層從左至右

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	r1 := preorderTraversal(root.Left)
	ret := append([]int{root.Val}, r1...)
	r2 := preorderTraversal(root.Right)
	ret = append(ret, r2...)
	return ret

}

//其他題目
// Validate Binary Search Tree
// Binary Tree Preorder Traversal
// Binary Tree Postorder Traversal
// Binary Search Tree Iterator
// Kth Smallest Element in a BST
// Closest Binary Search Tree Value II
// Inorder Successor in BST
// Convert Binary Search Tree to Sorted Doubly Linked List
// Minimum Distance Between BST Nodes

func main() {

}
