package main

import "sort"

//查询最富有的用户
//https://leetcode.cn/problems/richest-customer-wealth/
func main() {
	account := [][]int{{1, 2, 3}, {1, 2, 3}}
	r := getRich(account)
	println(r)
}

func getRich(account [][]int) int {
	rs := []int{}
	for _, v1 := range account {
		num := sSum(v1)
		rs = append(rs, num)
	}
	sort.Ints(rs)
	return rs[len(rs)-1]
}

func sSum(v []int) int {
	r := 0
	for _, v1 := range v {
		r += v1
	}
	return r
}
