package main

// 拆分数位后四位数字的最小和
// https://leetcode.cn/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
import (
	"fmt"
	"sort"
)

func main() {
	num := 1234 //结果是37
	r := minimumSum(num)
	fmt.Printf("%v", r)
}

/*/**
* minimumSum
*  @Description: 求拆解最小和
*  @param num
*  @return int
*
*  @Author: hill
 */
func minimumSum(num int) int {
	var m []int
	for {
		if num == 0 {
			break
		}
		m = append(m, num%10)
		num = num / 10
	}
	sort.Ints(m)
	return m[0]*10 + m[1]*10 + m[2] + m[3]
}

//解题分析: 把一个四位数拆成四个数 然后任意拼接 求最小的值
//其实主要看 最终拼接时候 十位数是哪两个  取到最小的两个值当十位数 即可
