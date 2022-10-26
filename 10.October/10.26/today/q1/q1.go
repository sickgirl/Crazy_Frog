package main

import "fmt"

//1768. 交替合并字符串
//https://leetcode.cn/problems/merge-strings-alternately/

func main() {
	s1 := "abc"
	s2 := "dfgt"
	s3 := mergeAlternately(s1, s2)
	fmt.Printf("%v", s3)
}

func mergeAlternately(word1 string, word2 string) string {
	n1 := len(word1)
	n2 := len(word2)
	n := 0
	if n1 >= n2 {
		n = n1
	} else {
		n = n2
	}
	var res []byte
	for i := 0; i < n; i++ {
		if i < n1 {
			res = append(res, word1[i])
		}
		if i < n2 {
			res = append(res, word2[i])
		}

	}
	return string(res)
}
