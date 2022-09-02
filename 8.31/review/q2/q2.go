package main

//字符串转数字
//https://leetcode.cn/problems/string-to-integer-atoi/
import (
	"fmt"
	"math"
)

func main() {
	r := myAtoi(" -4193  2343fds")
	fmt.Print(r)
}
func myAtoi(str string) int {
	res, flag := 0, 1
	firstNumber := false
	end := false
	for _, s := range str {
		switch s {
		case ' ':
			if firstNumber {
				end = true
				break
			}
			continue
		case '+':
			if firstNumber {
				end = true
				break
			}
			firstNumber = true

		case '-':
			if firstNumber {
				end = true
				break
			}
			firstNumber = true
			flag = -1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			firstNumber = true
			res = res*10 + int((s - '0'))
			if res*flag > math.MaxInt32 {
				return math.MaxInt32
			}
			if res*flag < math.MinInt32 {
				return math.MinInt32
			}
		default:
			end = true
			break
		}
		if end {
			break
		}
	}
	return res * flag
}
