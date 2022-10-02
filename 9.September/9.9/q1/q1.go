package main

//判断链表是否是回文
//https://leetcode.cn/problems/palindrome-linked-list/
import "fmt"

func main() {
	link := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	}
	node := isPalindrome(link)
	fmt.Printf("%v", node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vals := []int{}

	for {
		if head == nil {
			break
		}
		vals = append(vals, head.Val)
		head = head.Next
	}
	n := len(vals)
	for i, v := range vals {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true

}
