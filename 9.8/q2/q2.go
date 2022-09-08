package main

// 删除中间节点
//https://leetcode.cn/problems/delete-middle-node-lcci/
func main() {
	link := &ListNode{
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
	node1 := link.Next.Next
	deleteNode(node1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node != nil {
		if node.Next == nil {
			node = nil
		} else {
			node.Val = node.Next.Val
			node.Next = node.Next.Next
		}
	}
}

func deleteNode1(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//题解 : 多联系链表
