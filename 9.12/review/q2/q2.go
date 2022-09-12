package main

import "fmt"

//欢乐数
//https://leetcode.cn/problems/happy-number/

func main() {
	x := isHappy(8)

	println(x)
}
func isHappy(n int) bool {
	m := map[int]bool{}
	i := 0
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
		//如果 这个n  曾经存在过   说明进入了死循环   这时候要看一下 n 是否等于1  如果不等于 则说明它不可能等于1了
		i++
		fmt.Printf("第 %d 轮结果: n : %d  , m[n]: %v", i, n, m[n])
		println("--")
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	println(sum)
	return sum
}
