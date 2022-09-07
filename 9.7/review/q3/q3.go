package main

import (
	"fmt"
)

//求是否为2的幂
//https://leetcode.cn/problems/power-of-two/
func main() {
	n := 2
	r := isPowerOfTwo(n)
	fmt.Printf("%v", r)
}
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

//题解思路  类似于找某个值的时候 根据数据变换曲线 可以考虑 牛顿迭代方程 有时间再优化
//考虑进了误区  可以直接通过 n & (n-1) 是否等于0 来判断
//n 为 100000...    n-1 为  011111...  所以 n & (n-1) == 0
//位运算快捷   也可以 转为 2进制  判断是否为 1个1  后面都是0   也能判断  就是涉及类型转换比较繁琐
//思路3  -n  -n为 n的反码+1   所以  如果是2的倍数 n&(-n) == n
//思路4  取巧  最大值为 2 ^ 30 次方  即为 30个2的相乘    所以 如果 2^30 是 当前数的倍数  则说明 对应数 也是2的幂
