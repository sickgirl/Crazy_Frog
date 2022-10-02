package main

//判断是否为回文
import (
	"fmt"
)

func main() {
	r := isPalindrome(11211)
	fmt.Printf("%v", r)
}

func isPalindrome(int2 int) bool {
	if int2 < 0 {
		return false
	}
	int3 := gethui(int2)
	return int2 == int3
}

func gethui(int2 int) int {
	r := 0
	for {
		if int2 == 0 {
			return r
		}
		y := int2 % 10
		r = r*10 + y
		int2 = int2 / 10
	}
}
