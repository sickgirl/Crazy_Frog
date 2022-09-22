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

	if n != rn {
		return false
	}

	for {
		temp := arr[0]
		if _, ok := map1[temp]; !ok {
			return false
		}
		tempList := map1[temp]
		for i, v := range tempList {
			if arr[i] != v {
				return false
			}
		}

		arr = arr[len(tempList):]
		if len(arr) == 0 {
			return true
		}

	}
	return true

}
