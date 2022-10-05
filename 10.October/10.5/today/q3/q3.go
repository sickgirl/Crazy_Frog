package main

import (
	"fmt"
	"sort"
)

//2418. 按身高排序
//https://leetcode.cn/problems/sort-the-people/
func main() {
	//mapInfo := map[string]int32{
	//	"roy":   18,
	//	"kitty": 16,
	//	"hugo":  21,
	//	"tina":  35,
	//	"jason": 23,
	//}
	//
	//type peroson struct {
	//	Name string
	//	Age  int32
	//}
	//
	//var lstPerson []peroson
	//for k, v := range mapInfo {
	//	lstPerson = append(lstPerson, peroson{k, v})
	//}
	//
	//sort.Slice(lstPerson, func(i, j int) bool {
	//	return lstPerson[i].Age < lstPerson[j].Age // 升序
	//})
	//fmt.Println(lstPerson)

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
