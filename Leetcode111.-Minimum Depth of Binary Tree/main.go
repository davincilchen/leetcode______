package main

import "fmt"

//https://leetcode.com/problems/minimum-depth-of-binary-tree/
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

func Pow(x, n int) int {
	ret := 1
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}

func findParent(idx int) int { //輸入陣列的意思不同
	return (idx+1)/2 - 1
}

func TransToNormalTree(s []int) []int {
	if len(s) == 0 {
		return []int{}
	}
	padding := make([]bool, len(s))
	ret := make([]int, 0, 1)

	level := 0
	lastTotalNodes := 0
	levelTotalNodes := 1
	levelCnt := 0

	for i, v := range s {
		if i <= 2 {
			ret = append(ret, v)
			if i == 2 || i == 0 {
				level++
				levelCnt = 0
				levelTotalNodes = Pow(2, level)
				lastTotalNodes = levelTotalNodes - 1
			}

			continue
		}

		levelCnt++
		curIdx := levelCnt + lastTotalNodes - 1
		parent := findParent(curIdx)
		for j := parent; ret[j] == -1 && j < levelTotalNodes; j++ {
			if len(padding) <= j {
				slice := make([]bool, j+1)
				copy(slice, padding)
				padding = slice
			}
			if padding[j] {
				continue
			}
			padding[j] = true
			ret = append(ret, []int{-1, -1}...)
			levelCnt += 2
		}
		ret = append(ret, v)

		if levelCnt == levelTotalNodes {
			level++
			levelCnt = 0
			levelTotalNodes = Pow(2, level)
			lastTotalNodes = levelTotalNodes - 1

		}
	}
	return ret
}

func buildTree(in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	tree := TransToNormalTree(in)
	tmp := make([]*TreeNode, len(tree))
	for i, v := range tree {
		if i == 0 {
			tmp[0] = &TreeNode{Val: v}
			continue
		}
		if v == -1 {
			continue
		}
		p := findParent(i)
		c := &TreeNode{Val: v}
		tmp[i] = c
		if i%2 == 0 {
			tmp[p].Right = c
		} else {
			tmp[p].Left = c
		}

	}
	DFS_Print(tmp[0])
	return tmp[0]

}

// func findParent(grandfather *TreeNode, newLevel bool, index int) (bool, *TreeNode) {

// 	if grandfather.Left == nil {
// 		return true, grandfather.Right
// 	}
// 	if index%2 == 0 {
// 		if grandfather.Left.Left == nil {
// 			return false, grandfather.Left
// 		}
// 	} else {
// 		if grandfather.Left.Right == nil {
// 			return false, grandfather.Left
// 		}
// 	}

// }

// func buildTree(in []int) *TreeNode {
// 	if len(in) == 0 {
// 		return nil
// 	}

// 	t := &TreeNode{}
// 	var left *TreeNode
// 	var right *TreeNode
// 	var parent *TreeNode
// 	//t = &TreeNode{Val: in[0]}
// 	//left := t.Left
// 	//right := t.Right

// 	for i, v := range in {
// 		if i == 0 {
// 			t = &TreeNode{Val: v}
// 			//left = t.Left
// 			right = t
// 			continue
// 		}
// 		if v == -1 {
// 			continue
// 		}

// 		c := &TreeNode{Val: v}
// 		if left == nil {
// 			parent = right
// 		} else {
// 			parent = left
// 		}
// 		if i%2 == 0 {
// 			parent.Right = c
// 		} else {
// 			parent.left = c
// 		}
// 		left = c.Left
// 		right = c.Right

// 	}
// 	DFS_Print(t)
// 	return t

// }

func DFS_Print(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println("DFS_Print:")
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			fmt.Print(node.Val, " --> ")
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}

	}

	fmt.Println()
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/solution/er-cha-shu-de-zui-da-shen-du-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// func maxDepth(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     queue := []*TreeNode{}
//     queue = append(queue, root)
//     ans := 0
//     for len(queue) > 0 {
//         sz := len(queue)
//         for sz > 0 {
//             node := queue[0]
//             queue = queue[1:]
//             if node.Left != nil {
//                 queue = append(queue, node.Left)
//             }
//             if node.Right != nil {
//                 queue = append(queue, node.Right)
//             }
//             sz--
//         }
//         ans++
//     }
//     return ans
// }

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

func main() {

}
