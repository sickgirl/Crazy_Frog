package main

import "fmt"

//1656. 设计有序流
//https://leetcode.cn/problems/design-an-ordered-stream/
func main() {

	arr1 := []string{"OrderedStream", "insert", "insert", "insert", "insert", "insert"}
	arr2 := []L{
		{3, "ccc"},
		{1, "aaa"},
		{2, "bbb"},
		{5, "eee"},
		{4, "ddd"},
	}

	res := [][]string{}
	OrderedStream1 := OrderedStream{}
	for i, v := range arr1 {
		if v == "OrderedStream" {
			OrderedStream1 = Constructor(len(arr2))
		}

		if v == "insert" {
			temp := OrderedStream1.Insert(arr2[i-1].id, arr2[i-1].value)
			res = append(res, temp)
		}
	}

	fmt.Printf("%v", res)
}

type OrderedStream struct {
	stream []string
	ptr    int
}

type L struct {
	id    int
	value string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n+1), 1}
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	s.stream[idKey] = value
	start := s.ptr
	for s.ptr < len(s.stream) && s.stream[s.ptr] != "" {
		s.ptr++
	}
	return s.stream[start:s.ptr]
}
