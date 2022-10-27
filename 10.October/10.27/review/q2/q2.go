package main

import "fmt"

//832. 翻转图像
//https://leetcode.cn/problems/flipping-an-image/
func main() {
	image := [][]int{
		{1, 1, 0}, {1, 0, 1}, {0, 0, 0},
	}
	res := flipAndInvertImage(image)
	res2 := flipAndInvertImage2(image)
	fmt.Printf("%v  %v", res, res2)
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

func flipAndInvertImage2(image [][]int) [][]int {
	for _, row := range image {
		left, right := 0, len(row)-1
		for left < right {
			if row[left] == row[right] {
				row[left] ^= 1
				row[right] ^= 1
			}
			left++
			right--
		}
		if left == right {
			row[left] ^= 1
		}
	}
	return image
}

//
