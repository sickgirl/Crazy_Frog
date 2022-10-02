package main

import "fmt"

//解密字符串
//https://leetcode.cn/problems/decode-the-message/
//中文翻译确实是有问题
//key 意思就是对照表  "fffdsa..."  翻译成对照表其实是
//{
//f:a, d:b, s:c, a:d ...
//}
func main() {
	key := "the quick brown fox jumps over the lazy dog"
	msg := "vkbs bs t suepuv"
	r := decodeMessage(key, msg)
	fmt.Printf("%v", r)
}

func decodeMessage(key string, message string) string {
	s := "abcdefghijklmnopqrstvuwxyz"
	i := 0
	map1 := map[string]string{
		" ": " ",
	}
	for _, v := range key {
		oneString := string(v)
		if _, ok := map1[oneString]; !ok {
			map1[oneString] = string(s[i])
			i++
		}
	}
	r := ""
	for _, v := range message {
		r += map1[string(v)]
	}

	return r
}

//解题思路   byte 本质是 数字   含义是 当前字符在 字符表中是第几位
//所以没必要 操作完整字符表   当前对byte操作还是不熟练
