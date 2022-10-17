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
		temp := v[0]
		starts[temp] = true
	}
	for _, v := range paths {
		temp1 := v[1]
		if _, ok := starts[temp1]; !ok {
			return v[1]
		}
	}
	return ""
}
