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
	mp := ['z' + 1]byte{' ': ' '}
	cur := byte('a')
	for _, c := range key {
		if mp[c] == 0 {
			mp[c] = cur
			cur++
		}
	}
	s := []byte(message)
	for i, c := range s {
		s[i] = mp[c]
	}
	return string(s)
}
