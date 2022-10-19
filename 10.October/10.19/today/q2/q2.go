package main

//面试题 01.02. 判定是否互为字符重排
//https://leetcode.cn/problems/check-permutation-lcci/
import (
	"fmt"
	"sort"
)

// 将[]string定义为MyStringList类型
type MyStringList []byte

// 实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// 实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// 实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func main() {

	a := "abc"
	b := "cbae"
	r := CheckPermutation(a, b)
	fmt.Printf("%v", r)
}
func CheckPermutation(s1 string, s2 string) bool {
	s11 := changeStr(s1)
	s22 := changeStr(s2)
	return s11 == s22
}
func changeStr(s string) string {
	temp := MyStringList{}
	for _, v := range s {
		temp = append(temp, byte(v))
	}
	sort.Sort(temp)
	return string(temp)
}
