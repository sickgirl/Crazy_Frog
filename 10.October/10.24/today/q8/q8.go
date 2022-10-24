package main

import "fmt"

//剑指 Offer 22. 链表中倒数第k个节点
//https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
func main() {
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	r := getKthFromEnd(&list, 2)
	fmt.Printf("%v %v", r.Val, r.Next.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	temp := head
	n := 0
	for {
		if temp == nil {
			break
		}
		n++
		temp = temp.Next
	}
	l := n - k
	for i := 0; i < l; i++ {
		head = head.Next
	}
	return head
}
