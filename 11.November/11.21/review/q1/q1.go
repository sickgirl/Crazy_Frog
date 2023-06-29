package main

//2180. 统计各位数字之和为偶数的整数个数
//https://leetcode.cn/problems/count-integers-with-even-digit-sum/
func main() {
	num := 30

	res := countEven(num)
	println(res)
}

func countEven(num int) int {
	res := 0
	for {
		temp := getEven(num)
		if temp {
			println(num)
			res++
		}

		num = num - 1
		if num == 0 {
			break
		}
	}
	return res

}

func getEven(num int) bool {
	sum := 0
	for num > 0 {
		temp := num % 10
		sum += temp
		num = num / 10
	}

	if sum%2 == 0 {
		return true
	} else {
		return false
	}
}
