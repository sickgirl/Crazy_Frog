package main

import "fmt"

//1720. 解码异或后的数组
//https://leetcode.cn/problems/decode-xored-array/
func main() {
	encoded := []int{1, 2, 3}
	first := 1
	r := decode(encoded, first)
	fmt.Printf("%v", r)
}
func decode(encoded []int, first int) []int {
	r := []int{
		first,
	}
	for _, v := range encoded {
		r = append(r, first^v)
		first = first ^ v
	}
	return r
}
