package main

import "fmt"

//657. 机器人能否返回原点
//https://leetcode.cn/problems/robot-return-to-origin/
func main() {
	moves := "UD"
	r := judgeCircle(moves)
	fmt.Printf("%v", r)
}

func judgeCircle(moves string) bool {
	u := 0
	l := 0
	for _, v := range moves {
		if v == 'U' {
			u++
		}
		if v == 'D' {
			u--
		}
		if v == 'L' {
			l++
		}
		if v == 'R' {
			l--
		}
	}
	return (u == 0) && (l == 0)
}
