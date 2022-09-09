package main

// 罗马数字转整型
//https://leetcode.cn/problems/roman-to-integer/
import (
	"fmt"
	"strings"
)

func main() {
	s := "III"
	node := romanToInt(s)
	fmt.Printf("%v", node)
}

func romanToInt(s string) int {
	t4 := strings.Count(s, "IV")
	s = strings.ReplaceAll(s, "IV", "")

	t9 := strings.Count(s, "IX")
	s = strings.ReplaceAll(s, "IX", "")

	t40 := strings.Count(s, "XL")
	s = strings.ReplaceAll(s, "XL", "")

	t90 := strings.Count(s, "XC")
	s = strings.ReplaceAll(s, "XC", "")

	t400 := strings.Count(s, "CD")
	s = strings.ReplaceAll(s, "CD", "")

	t900 := strings.Count(s, "CM")
	s = strings.ReplaceAll(s, "CM", "")

	t1 := strings.Count(s, "I")
	t5 := strings.Count(s, "V")
	t10 := strings.Count(s, "X")
	t50 := strings.Count(s, "L")
	t100 := strings.Count(s, "C")
	t500 := strings.Count(s, "D")
	t1000 := strings.Count(s, "M")
	return t1 + 5*t5 + 10*t10 + 50*t50 + 100*t100 + 500*t500 + 1000*t1000 + 4*t4 + 40*t40 + 400*t400 + 9*t9 + 90*t90 + 900*t900

}
