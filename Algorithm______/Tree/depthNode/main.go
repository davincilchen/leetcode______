package main



// Example 1:

// Input: root = [3,9,20,null,null,15,7] 
// deep 1 : 3
// deep 2 : 9,20
// deep 3 : 15,7
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}




func DepthNode(tree []int, depth int) []int {
	root := buildTree_slice(tree)
	return depthNode(root, depth)
}




func depthNode(root *TreeNode, depth int) []int {
	if root == nil {
		return []int{}
	}

	if depth == 0 {
		return []int{}
	}

	if depth == 1 {
		return []int{root.Val}
	}

	ret := []int{}

	queue := []*TreeNode{}
	queue = append(queue, root) 
	check := 0
	for len(queue) > 0 {
		check ++ 
		if check == depth {
			
			for len(queue) > 0 {
				ret = append(ret, queue[0].Val)
				queue = queue[1:]
			}
			break
		}
		size := len(queue)
		for size > 0 {
			size --
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}			
		}

	}
	

	return ret
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
