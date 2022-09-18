package main

//1624. 两个相同字符之间的最长子字符串
//https://leetcode.cn/problems/largest-substring-between-two-equal-characters/
import "fmt"

func main() {
	s := "aabbccddeeffa" //长度13  预计结果 11
	r := maxLengthBetweenEqualCharacters(s)
	fmt.Printf("%v", r)
}

func maxLengthBetweenEqualCharacters(s string) int {
	ans := -1
	firstIndex := [26]int{}

	fmt.Printf("%v", firstIndex)
	for i := range firstIndex {
		fmt.Println(i)
		firstIndex[i] = -1
	}
	for i, c := range s {
		c -= 'a'
		if firstIndex[c] < 0 {
			firstIndex[c] = i //首次出现位置
		} else {
			//再次出现时候   用当前位置 减去初始位置   计算为长度   并且求最大值
			ans = max(ans, i-firstIndex[c]-1)
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
