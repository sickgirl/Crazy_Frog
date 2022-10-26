package main

import "fmt"

//832. 翻转图像
//https://leetcode.cn/problems/flipping-an-image/
func main() {
	image := [][]int{
		{1, 1, 0}, {1, 0, 1}, {0, 0, 0},
	}
	res := flipAndInvertImage(image)
	fmt.Printf("%v", res)
}
func flipAndInvertImage(image [][]int) [][]int {
	var res [][]int
	for _, v1 := range image {
		var r []int
		for _, v12 := range v1 {
			v12 ^= 1
			r = append([]int{v12}, r...)
		}
		res = append(res, r)
	}
	return res
}
