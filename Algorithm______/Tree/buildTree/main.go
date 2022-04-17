package main

import (
	"container/list"
	"fmt"
	"time"
)

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

func buildTree1(in []int) *TreeNode {
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
	return tmp[0]

}

func BFS_Print_String(root *TreeNode) string {
	if root == nil {
		return ""
	}
	str := "BFS_Print:"
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			ss := fmt.Sprintf("%d", node.Val)
			str += ss
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

	return str
}

func BFS_Print(root *TreeNode) {
	s := BFS_Print_String(root)
	fmt.Println(s)
}

func buildTree_linklist(tree []int) *TreeNode {
	if len(tree) == 0 {
		return nil
	}

	cur := 0
	root := &TreeNode{Val: tree[cur]}
	cur++

	tlist := list.New()
	tlist.PushBack(root)
	for cur < len(tree) {
		v1 := tree[cur]
		cur++
		v2 := -1
		if cur < len(tree) {
			v2 = tree[cur]
		}
		cur++
		e := tlist.Front()
		if e == nil {
			break
		}
		tlist.Remove(e)
		node := e.Value.(*TreeNode)
		if v1 != -1 {
			node.Left = &TreeNode{
				Val: v1,
			}
			tlist.PushBack(node.Left)
		}
		if v2 != -1 {
			node.Right = &TreeNode{
				Val: v2,
			}
			tlist.PushBack(node.Right)
		}
	}

	return root

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

// 200 node
// buildTree1  2.5434089s
// buildTree_linklist  0s
// buildTree_slice  0s

//if len(tlist) > 5000 { slice 有到23472

// 25600 node
//buildTree_linklist  3.0001ms
//buildTree_slice  2.9994ms

// 15200 node
// buildTree_linklist  6.9996ms
// buildTree_slice  5.4364ms
func main() {

	now := time.Now()
	t1 := buildTree1(testTree)
	fmt.Println("buildTree1 ", time.Since(now))
	//BFS_Print(t1)
	now = time.Now()
	t2 := buildTree_linklist(testTree)
	fmt.Println("buildTree_linklist ", time.Since(now))
	//BFS_Print(t2)
	now = time.Now()
	t3 := buildTree_slice(testTree)
	fmt.Println("buildTree_slice ", time.Since(now))
	//BFS_Print(t3)

	in := append(testTree, testTree2...)
	in = append(in, testTree2...)
	in = append(in, testTree2...)
	in = append(in, testTree2...)
	in = append(in, testTree2...)
	in = append(in, testTree2...)
	in = append(in, testTree2...)

	in = append(in, in...)
	in = append(in, in...)
	in = append(in, in...)
	in = append(in, in...)
	in = append(in, in...)

	now = time.Now()
	t4 := buildTree_linklist(in)
	fmt.Println("buildTree_linklist ", time.Since(now))

	now = time.Now()
	t5 := buildTree_slice(in)
	fmt.Println("buildTree_slice ", time.Since(now))
	//BFS_Print(t5)
	fmt.Println(t1, t2, t3, t4, t5)
}

var testTree = []int{2, 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8, //194 node
	//3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	//3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	// 3, 4, -1, 5, -1, 6, -1, 7, -1, 8,
	3, 4, -1, 5, -1, 6, -1, 7, -1, 8}

var testTree2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
