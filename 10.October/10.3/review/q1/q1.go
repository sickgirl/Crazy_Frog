package main

import "fmt"

//剑指 Offer II 042. 最近请求次数
//https://leetcode.cn/problems/H8086Q/
func main() {

	t := []int{
		1,
		100,
		3001,
		3002,
	}

	temp := Constructor()
	x := &temp
	res := []int{}
	for _, v := range t {
		res = append(res, x.Ping(v))
	}
	fmt.Printf("%v", res)

}

type RecentCounter []int

func Constructor() (_ RecentCounter) { return }

func (q *RecentCounter) Ping(t int) int {
	*q = append(*q, t)
	for (*q)[0] < t-3000 {
		*q = (*q)[1:]
	}
	return len(*q)
}

//用 * 替代外部存储
