package main

//唯一摩尔斯密码词
//https://leetcode.cn/problems/unique-morse-code-words/
import "fmt"

func main() {
	s := []string{"aa", "bb", "cc", "dd"}

	var r = uniqueMorseRepresentations(s)
	fmt.Printf("%v ", r)
}
func uniqueMorseRepresentations(words []string) int {
	types := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	res := map[string]bool{}
	for _, v := range words {
		temp := ""
		for _, v1 := range v {
			temp += types[v1-97]
		}
		res[temp] = true
	}

	return len(res)
	//types := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
}
