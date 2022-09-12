package main

import "fmt"

//有效的括号
//https://leetcode.cn/problems/valid-parentheses/
func main() {

	str := "(){}[]({})({)}"
	r := isValid(str)
	fmt.Printf("%v", r)
}

func isValid(str string) bool {
	var l []string

	for _, v := range str {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			l = append(l, string(v))
		} else {
			n := len(l)
			if n == 0 {
				return false
			}
			if string(v) == ")" {
				if l[n-1] == "(" {
					l = l[:n-1]
				} else {
					return false
				}
			}
			if string(v) == "]" {
				if l[n-1] == "[" {
					l = l[:n-1]
				} else {
					return false
				}
			}
			if string(v) == "}" {
				if l[n-1] == "{" {
					l = l[:n-1]
				} else {
					return false
				}
			}
		}
	}
	n := len(l)
	if n == 0 {
		return true
	} else {
		return false
	}
}
