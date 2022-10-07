package main

import "fmt"

//1313. 解压缩编码列表
//https://leetcode.cn/problems/decompress-run-length-encoded-list/
func main() {
	nums := []int{1, 1, 2, 3}
	r := decompressRLElist(nums)
	fmt.Printf("%v", r)
}
func decompressRLElist(nums []int) []int {
	n := len(nums)
	var res []int
	for i := 0; i < n; i += 2 {
		for i1 := 0; i1 < nums[i]; i1++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}
