package main

import "fmt"

//1684. 统计一致字符串的数目
//https://leetcode.cn/problems/count-the-number-of-consistent-strings/
func main() {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	r := countConsistentStrings(allowed, words)
	fmt.Printf("%v", r)
}

func countConsistentStrings(allowed string, words []string) int {
	ans := map[rune]bool{}
	r := 0
	for _, v := range allowed {
		ans[v] = true
	}
	for _, v := range words {
		error1 := false
		for _, v1 := range v {
			if !ans[v1] {
				error1 = true
				break
			}
		}
		if !error1 {
			r++
		}
	}
	return r
}
