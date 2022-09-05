package main

import (
	"fmt"
	"strings"
)

//宝石与石头
//https://leetcode.cn/problems/jewels-and-stones/
func main() {
	r := numJewelsInStones("aA", "aAAbbbb")
	fmt.Printf("%v", r)
}

func numJewelsInStones(jewels string, stones string) int {
	r := 0
	for _, v := range stones {
		if strings.Contains(jewels, string(v)) {
			r++
		}
	}
	return r
}

//解题思考:
//遍历每一块石头 看看它在不在宝石行列里
