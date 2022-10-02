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
		needPay := v - 5
		for 20 <= needPay && 0 < map1[20] {
			needPay -= 20
			map1[20]--
		}
		for 10 <= needPay && 0 < map1[10] {
			needPay -= 10
			map1[10]--
		}
		for 5 <= needPay && 0 < map1[5] {
			needPay -= 5
			map1[5]--
		}
		if needPay != 0 {
			return false
		}
		map1[v]++
	}
	return true
}

//优化了下题解:
//思路:优先花大票  小票灵活更容易找出去
//解题时候 判断 需要付的钱 是否大于对应面额  然后看对应面额是否有余额   循环花 直到 需要付的钱少 或者 柜台钱不够的时候
