package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		// 入れ替え
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		// 次のペアへ進む
		prev = first
		head = first.Next
	}

	return dummy.Next
}
