package main

import "fmt"

//有效的括号
func main() {
	str := "(){}[]({})(){}"
	r := isValid(str)
	fmt.Printf("%v", r)
}
func isValid(str string) bool {
	var temp string
	for _, v := range str {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			temp += string(v)
		} else {
			if len(temp) == 0 {
				return false
			}
		}

		if string(v) == ")" {
			if string(temp[len(temp)-1]) == "(" {
				temp = temp[0 : len(temp)-1]
			} else {
				return false
			}
		}
		if string(v) == "}" {
			if string(temp[len(temp)-1]) == "{" {
				temp = temp[0 : len(temp)-1]
			} else {
				return false
			}
		}
		if string(v) == "]" {
			if string(temp[len(temp)-1]) == "[" {
				temp = temp[0 : len(temp)-1]
			} else {
				return false
			}
		}
	}

	if len(temp) == 0 {
		return true
	} else {
		return false
	}
}
