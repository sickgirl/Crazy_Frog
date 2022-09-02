package main

//欢乐数
//https://leetcode.cn/problems/happy-number/

func main() {
	//x := isHappy(100)
	var x int
	var y int
	x, y = 1, 2
	println(x, y)
}
func isHappy(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
