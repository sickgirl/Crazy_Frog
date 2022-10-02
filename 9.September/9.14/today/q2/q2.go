package main

import (
	"fmt"
	"sort"
)

//删除某些元素后的数组均值
//https://leetcode.cn/problems/mean-of-array-after-removing-some-elements/
func main() {
	s := []int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0}
	ans := trimMean(s)
	fmt.Printf("%v", ans)
}

func trimMean(arr []int) float64 {
	n := len(arr)
	b := n / 20
	sort.Ints(arr)
	res := arr[b : n-b]
	sum1 := 0
	for _, v := range res {
		sum1 += v
	}
	return float64(sum1) / float64(len(res))
}
