package main

//分割平衡字符串
//https://leetcode.cn/problems/split-a-string-in-balanced-strings/
import "fmt"

func main() {
	s := "RLLR"
	r := balancedStringSplit2(s)
	fmt.Printf("%v", r)
}
func balancedStringSplit(s string) int {
	temp := 0   //储存当前字节类型
	tempNo := 0 //储存当前字节数量
	n := 0      // 总数量

	for _, v := range s {
		if temp == 0 {
			temp = int(v)
			tempNo = 1
		} else {
			if temp == int(v) {
				tempNo++
			} else {
				tempNo--
				if tempNo == 0 {
					temp = 0
					tempNo = 0
					n++
				}
			}
		}
	}
	return n
}
func balancedStringSplit1(s string) int {
	l := 0
	res := 0
	for _, v := range s {
		if v == 'L' {
			l++
		} else {
			l--
		}
		if l == 0 {
			res++
		}
	}
	return res
}
func balancedStringSplit2(s string) int {
	res := 0
	for i, v := range s {
		if i != 0 {
			if s[i-1] != byte(v) {
				res++
			}
		}
	}
	return res
}
