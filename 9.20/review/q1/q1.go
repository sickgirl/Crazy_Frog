package main

import (
	"fmt"
	"sort"
)

type temp struct {
	k int
	v int
}
type temps []temp

//  实现sort包中Interface接口

func (t temps) Len() int {
	return len(t)
}

func (t temps) Less(i, j int) bool {
	return t[i].v < t[j].v
}

func (t temps) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func main() {
	nums := []int{1, 1, 2, 2, 2, 3}
	r := frequencySort(nums)
	fmt.Printf("%v", r)
}

func frequencySort(nums []int) []int {

	map1 := make(map[int]int)

	for _, v := range nums {
		map1[v]++
	}

	var ts temps

	for k, v := range map1 {
		ts = append(ts, temp{k: k, v: v})
	}

	var res []int
	sort.Sort(ts)
	for _, v := range ts {
		for i := 0; i < v.v; i++ {
			res = append(res, v.k)
		}
	}
	return res
}
