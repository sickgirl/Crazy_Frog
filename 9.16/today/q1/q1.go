package main

import "fmt"

//剑指 Offer 15. 二进制中1的个数
//https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func main() {
	var n uint32 = 11
	r := hammingWeight(n)
	fmt.Printf("%v", r)

}

func hammingWeight(num uint32) int {

	n := 0
	for {
		if num == 0 {
			return n
		}
		if num%2 == 1 {
			n++
		}
		num = num / 2
	}
}
