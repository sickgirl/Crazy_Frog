package main

//面试题 02.02. 返回倒数第 k 个节点
//https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/
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
	k := 2
	r := kthToLast(&list, k)
	println(r)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
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
	return head.Val
}
