package main

// https://medium.com/@houzier.saurav/dfs-and-bfs-golang-d5818ec690d3

// https://www.geeksforgeeks.org/breadth-first-traversal-bfs-on-a-2d-array/
//BFS to print the tree in breadth first fashion
type Tree struct {
	Value  int
	Left   *Tree
	Right  *Tree
	Repeat int
	Parent *Tree
}

func BFS(tree *Tree) []int { //給bfs 的link list就會好做iteration
	queue := []*Tree{}
	queue = append(queue, tree)
	result := []int{}
	return BFSUtil(queue, result)
}

func BFSUtil(queue []*Tree, res []int) []int {
	//This appends the value of the root of subtree or tree to the
	//result
	if len(queue) == 0 {
		return res
	}
	res = append(res, queue[0].Value)
	if queue[0].Right != nil {
		queue = append(queue, queue[0].Right)
	}
	if queue[0].Left != nil {
		queue = append(queue, queue[0].Left)
	}
	return BFSUtil(queue[1:], res)
}

// ============ //
// // iterate over list elements
// for e := alist.Front(); e != nil; e = e.Next() {
// 	fmt.Println(e.Value.(string))
// }
//

type Iteration interface {
	GetValue() int
	Next() Iteration
}

func newNode(val int, queue []int) *Node {

	return &Node{
		queue: queue,
	}
}

type Node struct {
	queue []int
	val   int
	next  int
}

func (t *Node) GetValue() int {
	return t.val
}

func (t *Node) Next() Iteration { //原本就是link list會很方便
	if t.next > 0 && t.next < len(t.queue) {

		return nil //??
	}

	return nil
}
