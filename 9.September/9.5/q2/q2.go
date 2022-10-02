package main

import "fmt"

//找零钱
//https://leetcode.cn/problems/lemonade-change/
func main() {
	bills := []int{5, 5, 5, 10, 20}
	r := lemonadeChange(bills)
	fmt.Printf("%v", r)
}

func lemonadeChange(bills []int) bool {
	map1 := map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}

	for _, v := range bills {
		if v == 5 {
			map1[5]++
		}
		if v == 10 {
			if map1[5] == 0 {
				return false
			} else {
				map1[5]--
				map1[10]++
			}
		}
		if v == 20 {
			if map1[10] > 0 {
				if map1[5] == 0 {
					return false
				} else {
					map1[5]--
					map1[10]--
				}
			} else {
				if map1[5] >= 3 {
					map1[5] = map1[5] - 3
				} else {
					return false
				}
			}
		}
	}
	return true
}
