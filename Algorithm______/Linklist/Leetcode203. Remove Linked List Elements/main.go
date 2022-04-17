package main

// 203. Remove Linked List Elements
// Easy

// 4799

// 162

// Add to List

// Share
// Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

// Example 1:

// Input: head = [1,2,6,3,4,5,6], val = 6
// Output: [1,2,3,4,5]
// Example 2:

// Input: head = [], val = 1
// Output: []
// Example 3:

// Input: head = [7,7,7,7], val = 7
// Output: []

// Constraints:

// The number of nodes in the list is in the range [0, 104].
// 1 <= Node.val <= 50
// 0 <= val <= 50

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime: 10 ms, faster than 48.04% of Go online submissions for Remove Linked List Elements.
// Memory Usage: 4.6 MB, less than 100.00% of Go online submissions for Remove Linked List Elements.
func removeElements1(head *ListNode, val int) *ListNode {

	cur := head
	var pre *ListNode
	for cur != nil {
		if cur.Val != val {
			pre = cur
			cur = cur.Next
			continue
		}
		if cur == head { //remove head
			cur = cur.Next
			head = cur
			continue
		}
		next := cur.Next
		cur.Next = nil
		cur = next
		if pre == nil {
			//pre = head
			head.Next = next
		} else {
			pre.Next = next
		}

	}

	return head
}

// Runtime: 10 ms, faster than 48.04% of Go online submissions for Remove Linked List Elements.
// Memory Usage: 4.6 MB, less than 100.00% of Go online submissions for Remove Linked List Elements.
func removeElements2(head *ListNode, val int) *ListNode {

	cur := head
	var pre *ListNode
	for cur != nil {
		if cur.Val != val {
			pre = cur
			cur = cur.Next
			continue
		}
		if cur == head { //remove head
			cur = cur.Next
			head = cur
			continue
		}
		// next := cur.Next
		// cur.Next = nil
		// cur = next
		if pre == nil {
			pre = head
			//head.Next = next
		} else {
			pre.Next = cur.Next
		}
		cur = pre.Next
	}

	return head
}

func main() {

}
