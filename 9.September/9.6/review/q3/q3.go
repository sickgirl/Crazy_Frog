package main

import (
	"fmt"
	"sort"
	"strings"
)

//分割字符串的最大得分
//https://leetcode.cn/problems/maximum-score-after-splitting-a-string/
func main() {
	s := "10011"
	r := maxScore1(s)
	fmt.Printf("%v", r)
}

func maxScore(s string) int {
	full := len(s)
	max := 0
	for i := 1; i <= full-1; i++ {
		left := s[:i]
		right := s[i:]
		no1 := strings.Count(left, "0")
		no2 := strings.Count(right, "1")
		if no1+no2 > max {
			max = no1 + no2
		}
	}
	return max
}
func maxScore1(s string) int {
	full := len(s)
	max := 0
	left := s[:1]
	right := s[1:]
	no1 := strings.Count(left, "0")
	no2 := strings.Count(right, "1")
	max = no1 + no2
	var score []int
	score = append(score, max)

	for i := 1; i <= full-1; i++ {
		if string(s[i]) == "0" {
			max++
		} else {
			max--
		}
		score = append(score, max)
		fmt.Printf("%v", score)
		fmt.Println("")
	}
	sort.Ints(score)
	return score[full-1]
}

//题解思路  本想找规律 找到正确分割点 然后再计算价值
//事实上没有找到 正确分割点  直接遍历暴力破解   优化点   每次字符串计算比较耗内存  可以每次存储下来  下一次根据移动的元素  进行 +1 -1
//分割点 右移 跨过 1 时候可以忽略计算  因为总值一定是减一的
//分析后 得出新的题解思路  直接遍历所有元素 首先计算初始状态值   值为0 +1  值为1 -1 求最大值即可   算了懒得写了
//还是写了吧   写完觉得也节省不了啥计算量
