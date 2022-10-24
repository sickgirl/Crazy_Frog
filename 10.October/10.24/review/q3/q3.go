package main

import (
	"fmt"
	"sort"
)

//1528. 重新排列字符串
//https://leetcode.cn/problems/shuffle-string/
func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	r := restoreString(s, indices)
	fmt.Printf("%v", r)
}
func restoreString(s string, indices []int) string {
	type demo struct {
		name  rune
		index int
	}
	var r []rune
	var demos []demo
	for i, v := range s {
		d1 := demo{
			name:  v,
			index: indices[i],
		}
		demos = append(demos, d1)
	}

	sort.Slice(demos, func(i, j int) bool {
		return demos[i].index < demos[j].index
	})
	for _, v := range demos {
		r = append(r, v.name)
	}
	return string(r)
}
