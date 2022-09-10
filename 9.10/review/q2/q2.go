package main

import "fmt"

//文件夹深度
//https://leetcode.cn/problems/crawler-log-folder/
func main() {
	logs := []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}
	r := minOperations(logs)
	fmt.Printf("%v", r)
}
func minOperations(logs []string) int {
	num := 0
	for _, v := range logs {
		if v == "./" {
			continue
		}
		if v == "../" {
			if num >= 1 {
				num--
			}
		} else {
			num++
		}
	}
	return num
}
