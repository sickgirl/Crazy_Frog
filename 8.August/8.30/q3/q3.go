package main

//合并链表
//

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		},
	}

	l3 := mergeTwoLists(l1, l2)
	fmt.Printf("%v", l3)
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	result := prehead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}
	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}
