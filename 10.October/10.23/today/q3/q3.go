package main

import (
	"fmt"
	"sort"
)

//2418. 按身高排序
//https://leetcode.cn/problems/sort-the-people/
func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	res := sortPeople(names, heights)
	fmt.Printf("%v", res)

}

func sortPeople(names []string, heights []int) []string {
	type peroson struct {
		Name   string
		height int
	}
	var lstPerson []peroson
	for k, v := range names {
		h := heights[k]
		lstPerson = append(lstPerson, peroson{v, h})
	}
	sort.Slice(lstPerson, func(i, j int) bool {
		return lstPerson[i].height > lstPerson[j].height // 升序
	})
	var res []string
	for _, v := range lstPerson {
		res = append(res, v.Name)
	}
	return res
}
