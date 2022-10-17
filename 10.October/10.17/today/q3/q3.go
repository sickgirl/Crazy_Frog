package main

import (
	"fmt"
	"strings"
)

//2315. 统计星号
//https://leetcode.cn/problems/count-asterisks/
func main() {
	s := "l|*e*et|c**o|*de|"
	r := countAsterisks2(s)
	fmt.Printf("%v", r)
}
func countAsterisks(s string) int {
	num1 := 0 //|的序号
	num2 := 0 //|外*的数量
	for _, v := range s {
		if v == '|' {
			num1++
		}
		if num1%2 == 0 && v == '*' {
			num2++
		}
	}
	return num2
}

func countAsterisks2(s string) (ans int) {
	sp := strings.Split(s, "|")
	fmt.Printf("%v", sp)
	for i := 0; i < len(sp); i += 2 {
		ans += strings.Count(sp[i], "*")
	}
	return
}
