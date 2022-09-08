package main

//83. 删除排序链表中的重复元素
//https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
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
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	node := deleteDuplicates(link)
	fmt.Printf("%v", node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	map1 := make(map[int]bool, 0)
	test := head
	if test != nil {
		map1[test.Val] = true
	}

	for {
		if test != nil && test.Next != nil {
			if _, ok := map1[test.Next.Val]; ok {
				test.Next = test.Next.Next
			} else {
				map1[test.Next.Val] = true
				test = test.Next
			}
		} else {
			break
		}

	}
	return head
}
