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
	studentMap := map[int]int{}
	for _, v := range students {
		studentMap[v]++
	}
	eat := 0
	for _, v := range sandwiches {
		if studentMap[v] > 0 {
			studentMap[v]--
			eat++
		} else {
			break
		}
	}
	return len(students) - eat
}

//思考: 原以为跟顺序无关 只和总数有关    事实上跟面包顺序还是有关的
//面包顺序影响结果    比如  面包是  101111   学生是 11111 第二个面包没人吃   后面的都吃不到了
//所以 直观写的队列轮转的方式 还是比较差的
