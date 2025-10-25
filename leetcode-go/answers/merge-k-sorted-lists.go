package main

import (
	"container/heap"
)

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val // 小さい順
}

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// ---- main ----

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// 各リストの先頭を heap に入れる
	for _, node := range lists {
		if node != nil {
			heap.Push(pq, node)
		}
	}

	dummy := &ListNode{}
	tail := dummy

	// heap が空になるまで最小ノードを取り出して結合
	for pq.Len() > 0 {
		minNode := heap.Pop(pq).(*ListNode)
		tail.Next = minNode
		tail = tail.Next

		// 次ノードがあれば heap に追加
		if minNode.Next != nil {
			heap.Push(pq, minNode.Next)
		}
	}

	return dummy.Next
}
