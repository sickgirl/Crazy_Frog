package main

import (
	"fmt"
	"math"
)

//判断欢乐数
func main() {
	r := isHappyNo(12)
	fmt.Printf("%v", r)
}

func isHappyNo(n int) bool {
	//在处理数字过程中 往往是三种情况   1.归1   2.循环  3.逐步变大
	big := math.Pow(2, 31)
	i := 0
	res := make(map[int]bool)
	for {
		i++
		n = makeHappy(n)
		println("第", i, "轮循环: n值为:", n)
		if n == 1 {
			return true
		}
		if _, ok := res[n]; ok {
			println("命中循环n值为", n)
			return false
		}
		if n >= (int(big) - 1) {
			println("命中超大数n值为", n)
			return false
		}

		res[n] = true
		if i >= 10000 {
			println("命中循环次数过多 n值为", n)
			return false
		}
	}
}
func makeHappy(n int) int {
	r := 0
	for n >= 10 {
		r += (n % 10) * (n % 10)
		n = n / 10
	}
	return r + n*n
}
