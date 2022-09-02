package main

import "fmt"

//最长公共前缀
func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Printf("%v", res)
}
func longestCommonPrefix(strs []string) string {
	len1 := len(strs[0])
	r := ""
	for i := 0; i < len1; i++ {
		for _, v := range strs {
			if len(v) <= i {
				return r
			}
			if v[i] != strs[0][i] {
				return r
			}
		}
		r = r + string(strs[0][i])
	}
	return r
}
