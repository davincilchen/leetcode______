package main

//https://leetcode.com/problems/binary-tree-postorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Pre-order 前序：root → left → right
// In-order 中序：left → root → right
// Post-order 後序：left → right → root
// Level-order 層序：一層一層從左至右

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := postorderTraversal(root.Left)

	r2 := postorderTraversal(root.Right)
	ret = append(ret, r2...)
	ret = append(ret, root.Val)
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
