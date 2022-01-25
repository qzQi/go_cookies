/*
codeTop里面的一些题，用go来写。这些都是很熟练的题目
*/
package leetcode

// import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList01(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, nil}
	ans := dummy
	for head != nil {
		dummy = head
		head = head.Next
		dummy.Next = ans.Next
		ans.Next = dummy
	}
	return ans.Next
}

// 前面那个是类似于插入的方法；这个是改变指针的指向。
// 就需要多看看这个改变图
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
