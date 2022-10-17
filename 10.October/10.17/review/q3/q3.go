package main

import (
	"fmt"
	"math/bits"
)

//剑指 Offer 15. 二进制中1的个数
//https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func main() {
	var n uint32 = 11
	r := hammingWeight(n)
	fmt.Printf("%v", r)

}

func hammingWeight(num uint32) int {

	return bits.OnesCount32(num)
}

//为啥会有这种函数?
