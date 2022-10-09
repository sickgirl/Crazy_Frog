package main

import "fmt"

//1725. 可以形成最大正方形的矩形数目
//https://leetcode.cn/problems/number-of-rectangles-that-can-form-the-largest-square/
func main() {
	rectangles := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	r := countGoodRectangles(rectangles)
	fmt.Printf("%v", r)
}

func countGoodRectangles(rectangles [][]int) int {
	ans := map[int]int{}
	max := 0
	for _, v := range rectangles {
		min := min(v[0], v[1])
		ans[min]++
		if min > max {
			max = min
		}
	}
	return ans[max]
}

func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
