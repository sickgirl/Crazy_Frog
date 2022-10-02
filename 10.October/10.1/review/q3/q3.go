package main

import "fmt"

//1742. 盒子中小球的最大数量
//https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/
func main() {
	lowLimit := 1
	highLimit := 10
	r := countBalls(lowLimit, highLimit)
	fmt.Printf("%v", r)
}

func countBalls(lowLimit int, highLimit int) int {

	map1 := map[int]int{}
	max := 0
	for i := lowLimit; i <= highLimit; i++ {
		n := getNum(i)
		map1[n]++
		if map1[n] > max {
			max = map1[n]
		}
	}
	return max
}

func getNum(num int) int {
	n := 0
	for {
		if num == 0 {
			return n
		}
		n += num % 10
		num = num / 10
	}
}
