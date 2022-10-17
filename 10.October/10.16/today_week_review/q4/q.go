package main

//二进制链表转整数
//https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/
import "fmt"

func main() {
	link := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  0,
						Next: nil,
					},
				},
			},
		},
	}
	node := getDecimalValue(link)
	fmt.Printf("%v", node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}
	r := 0
	for {
		if head == nil {
			return r
		}
		r = r*2 + head.Val
		head = head.Next
	}

}
