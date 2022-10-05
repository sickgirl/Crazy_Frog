package main

import "fmt"

//1832. 判断句子是否为全字母句
//https://leetcode.cn/problems/check-if-the-sentence-is-pangram/
func main() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	r := checkIfPangram(sentence)
	fmt.Printf("%v", r)
}
func checkIfPangram(sentence string) bool {
	ans := make(map[rune]bool)
	for _, v := range sentence {
		ans[v] = true
	}
	return len(ans) == 26
}
