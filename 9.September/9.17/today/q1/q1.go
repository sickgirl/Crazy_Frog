package main

//1624. 两个相同字符之间的最长子字符串
//https://leetcode.cn/problems/largest-substring-between-two-equal-characters/
import "fmt"

func main() {
	s := "cbzxy" //长度13  预计结果 11
	r := maxLengthBetweenEqualCharacters(s)
	fmt.Printf("%v", r)
}

func maxLengthBetweenEqualCharacters(s string) int {
	n := len(s)
	map2 := make(map[byte]int)
	max := -1
	never := true
	for i := 0; i < n; i++ {
		temp1 := s[i] //当前字符
		if _, ok := map2[s[i]]; ok {
			continue
		}
		for i1 := 0; i < n; i1++ {
			temp2 := s[n-1-i1]
			if temp1 == temp2 {
				map2[s[i]] = n - 1 - i1 - i - 1
				never = false
				break
			}
		}
	}
	if never {
		return -1
	}

	for _, v := range map2 {
		if v > max {
			max = v
		}
	}
	return max
}
