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
			tempLen := 0
			for {
				tempLen = len(temp)
				if tempLen >= k {
					break
				}
				temp = append(temp, code...)
			}
			temp = temp[:k]
			tempRes := summation(temp)
			res = append(res, tempRes)

		} else {
			temp := code[:i]
			tempLen := 0
			for {
				tempLen = len(temp)
				if tempLen >= -k {
					break
				}
				temp = append(code, temp...)
			}
			temp = temp[tempLen+k:]
			tempRes := summation(temp)
			res = append(res, tempRes)
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
