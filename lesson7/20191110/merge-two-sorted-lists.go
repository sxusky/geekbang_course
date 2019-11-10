package main

import "fmt"

// 链-数据结构
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 使用递归的方式合并
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}

func main() {
	// https://leetcode-cn.com/problems/merge-two-sorted-lists
	ListNode1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	ListNode1.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	ListNode1.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	ListNode2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	ListNode2.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	ListNode2.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	ListNode12 := mergeTwoLists(ListNode1, ListNode2)
	tmp := ListNode12
	for tmp != nil {
		fmt.Println(tmp.Val, tmp, tmp.Next)
		tmp = tmp.Next
	}

}
