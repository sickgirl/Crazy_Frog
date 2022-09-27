package main

import "fmt"

//1700. 无法吃午餐的学生数量
//https://leetcode.cn/problems/number-of-students-unable-to-eat-lunch/
func main() {
	//[1,1,1,0,0,1]
	//[1,0,0,0,1,1]
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}
	res := countStudents(students, sandwiches)
	fmt.Printf("%v", res)
}

func countStudents(students []int, sandwiches []int) int {
	end := false
	for _, v := range sandwiches {
		n := len(students)
		if end {
			break
		}
		for {
			if students[0] == v {
				students = students[1:]
				break
			} else {
				if n == 0 {
					end = true
					break
				}
				temp := students[0]
				students = append(students[1:], temp)
				n--
			}
		}
	}
	return len(students)
}
