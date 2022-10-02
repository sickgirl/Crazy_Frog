package main

import "fmt"

//判断是否为回文

func main() {
	r := isPalindrome(11211)
	fmt.Printf("%v", r)
}

func isPalindrome(int2 int) bool {
	if int2 < 0 {
		return false //负值不会为回文
	}
	int3 := getHuiWen(int2)
	return int2 == int3
}

func getHuiWen(int2 int) int {
	int3 := 0
	for {
		if int2 == 0 {
			return int3
		}
		yu := int2 % 10
		int3 = int3*10 + yu
		int2 = int2 / 10
	}
}
