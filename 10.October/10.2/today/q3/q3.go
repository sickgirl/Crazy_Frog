package main

import "fmt"

//2373. 矩阵中的局部最大值
//https://leetcode.cn/problems/largest-local-values-in-a-matrix/
func main() {
	grid := [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}
	r := largestLocal(grid)
	fmt.Printf("%v", r)
}
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	for i, row := range grid[:n-2] {
		for j := 0; j < n-2; j++ {
			mx := 0
			for _, r := range grid[i : i+3] {
				for _, v := range r[j : j+3] {
					mx = max(mx, v)
				}
			}
			row[j] = mx
		}
		grid[i] = row[:n-2]
	}
	return grid[:n-2]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//题解: 没有什么思路而言   因为结果是确定的   相当于原有矩形框缩小两个维度
//而每个坐标的数据  都是以原数据为左上顶点的 3*3 矩形框中的最大值  即可得出答案
