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
	type person struct {
		name string
		h    int
	}
	var lstPerson []person
	for i, v := range names {
		p1 := person{
			name: v,
			h:    heights[i],
		}
		lstPerson = append(
			lstPerson,
			p1,
		)
	}
	sort.Slice(lstPerson, func(i, j int) bool {
		return lstPerson[i].h > lstPerson[j].h
	})
	var res []string
	for _, v := range lstPerson {
		res = append(res, v.name)
	}
	return res
}
