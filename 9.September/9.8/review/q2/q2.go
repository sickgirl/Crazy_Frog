package main

import "fmt"

//旅行终点站
//https://leetcode.cn/problems/destination-city/
func main() {
	paths := [][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}
	r := destCity(paths)
	fmt.Printf("%v", r)
}

func destCity(paths [][]string) string {
	starts := map[string]bool{}
	for _, v := range paths {
		starts[v[0]] = true
	}
	for _, v := range paths {
		if !starts[v[1]] {
			return v[1]
		}
	}
	return ""
}
