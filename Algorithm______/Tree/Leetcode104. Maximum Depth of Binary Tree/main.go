package main

//https ://leetcode.com/problems/minimum-depth-of-binary-tree/
// Given a binary tree, find its minimum depth.

// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

// Note: A leaf is a node with no children.

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: 2
// Example 2:

// Input: root = [2,null,3,null,4,null,5,null,6]
// Output: 5

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 
// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/solution/er-cha-shu-de-zui-da-shen-du-by-leetcode-solution/

func MaxDepth2(tree []int) int {
	root := buildTree_slice(tree)
	return maxDepth2(root)
}

func maxDepth2(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue := []*TreeNode{}
    queue = append(queue, root)
    ans := 0
    for len(queue) > 0 {
        sz := len(queue)
        for sz > 0 {
            node := queue[0]
            queue = queue[1:]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            sz--
        }
        ans++
    }
    return ans
}

//Max deep
//DFS
// func maxDepth(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 迭代實現，廣度優先遍歷
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode       // 查詢佇列
	queue = append(queue, root) // 將起點加入佇列
	depth := 1                  // root 本身就是一層，depth初始化為1
	for {
		if len(queue) == 0 { // 佇列為空時，退出迴圈
			break
		}
		size := len(queue)
		// 將當前佇列中的所有結點向四周擴散
		for i := 0; i < size; i++ {
			cur := queue[0]
			// 判斷是否到終點
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			// 將 cur的相鄰節點加入佇列
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			// 去掉當前節點
			queue = queue[1:]
		}
		// 這裡增加步數
		depth++
	}
	return depth
}

func MaxDepth(tree []int) int {
	root := buildTree_slice(tree)
	return maxDepth(root)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	val := max(left, right)

	return val + 1
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
