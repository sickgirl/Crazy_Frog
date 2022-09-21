package main

import "fmt"

//2413. 最小偶倍数
//https://leetcode.cn/problems/smallest-even-multiple/
func main() {
	n := 5
	r := smallestEvenMultiple(n)
	fmt.Printf("%v", r)
}
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return 2 * n
}

//题解分析: 如果是求两个数的最小公倍数   就需要 把两个数全部拆成  素数乘积     把素数集合 合并去重    最后素数之积 就为最小公倍数
//   难点在于 拆成素数之积    但是当前  固定一个是2   素数之积 就是2   只需要看 另一个值里 是否包含即可
