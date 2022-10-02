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
					Val: 2,
					Next: &ListNode{
						Val:  1,
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
	s1 := make([]int, 0, 0)
	for {
		if head == nil {
			lens := len(s1)
			for i, v := range s1 {
				if v != s1[lens-1-i] {
					return false
				}
			}
			return true
		}
		s1 = append(s1, head.Val)
		head = head.Next
	}
	return true

}

//题解思考: 如果是php  最简单直白方式 就是 遍历 链表  同时生成两个字符串
//a = a+v   b = v+b    一个前缀 一个后缀  最终比较两个是否相等 即可
//我用 go 按这个思路做的时候 发现一个问题
//由于go是强类型  而 node.Val 是int  拼装比较 只能是string
//类型转换对性能影响非常大  甚至会导致超时
//所以 遍历链表时候 生成切片  遍历切片 头尾比较  比较合适
//进一步优化  可以不用遍历全程  遍历一半就够了
