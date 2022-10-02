package main

import (
	"fmt"
	"strings"
)

//1694. 重新格式化电话号码
//https://leetcode.cn/problems/reformat-phone-number/
func main() {
	number := "123 4-5678"
	r := reformatNumber(number)
	fmt.Printf("%v", r)
}

func reformatNumber(number string) string {
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")
	new := ""
	for {
		if len(number) == 0 {
			return new
		}
		if len(number) == 4 {
			if new != "" {
				new = new + "-" + number[:2] + "-" + number[2:]
			} else {
				new = number[:2] + "-" + number[2:]
			}
			return new
		}
		n := 3
		if len(number) < 3 {
			n = len(number)
		}

		if new != "" {
			new = new + "-" + number[:n]
		} else {
			new = number[:n]
		}
		number = number[n:]
	}
}
