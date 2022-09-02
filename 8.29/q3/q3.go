package main

//最长公共前缀
//https://leetcode.cn/problems/longest-common-prefix/

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Printf("%v", res)
}
func longestCommonPrefix(strs []string) string {
	n := len(strs[0])
	r := ""
	for i := 0; i < n; i++ {
		taget := strs[0][i]
		for _, v := range strs {
			if i >= len(v) {

				return r
			}

			if v[i] != taget {

				return r
			}
		}
		r = r + string(taget)
	}
	return r
}
