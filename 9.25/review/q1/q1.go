package main

import "fmt"

//1652. 拆炸弹
//https://leetcode.cn/problems/defuse-the-bomb/
func main() {
	code := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := -3

	r := decrypt(code, k)
	fmt.Printf("%v", r)

}
func decrypt(code []int, k int) []int {
	var res []int
	for i, _ := range code {
		if k == 0 {
			res = append(res, 0)
		} else if k > 0 {
			temp := code[i+1:]
			temp = append(temp, code...)
			temp = temp[:k]
			resTemp := summation(temp)
			res = append(res, resTemp)

		} else {
			temp := code[:i]
			temp = append(code, temp...)
			temp = temp[len(temp)+k:]
			resTemp := summation(temp)
			res = append(res, resTemp)
		}

	}
	return res
}

func summation(slice []int) int {
	var sum int
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum
}

//题解  当前题目忽略了  k > len(s) 的情况 所以容易多了
