package main

//1640. 能否连接形成数组
//https://leetcode.cn/problems/check-array-formation-through-concatenation/
func main() {
	var arr []int = []int{49, 18, 16}
	var pieces [][]int = [][]int{{16, 18, 49}}
	r := canFormArray(arr, pieces)
	println(r)
}
func canFormArray(arr []int, pieces [][]int) bool {
	map1 := map[int][]int{}
	n := len(arr)
	rn := 0
	for _, v := range pieces {
		rn += len(v)
		map1[v[0]] = v
	}
	if rn != n {
		return false
	}

	for {
		temp := arr[0]
		if _, ok := map1[temp]; !ok {
			return false
		}
		for i, v := range map1[temp] {
			if arr[i] != v {
				return false
			}
		}
		arr = arr[len(map1[temp]):]
		if len(arr) == 0 {
			return true
		}

	}

}

//题解 1.输入两个参数  a,一个完整链条 b,几个链条碎片   问:是否能能通过组装b 得到a
