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
		if find := strings.Contains(jewels, string(v)); find {
			r++
		}
	}
	return r
}
