package main

import "fmt"

//1773. 统计匹配检索规则的物品数量
//https://leetcode.cn/problems/count-items-matching-a-rule/
func main() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	ruleKey := "color"
	ruleValue := "silver"
	r := countMatches(items, ruleKey, ruleValue)
	fmt.Printf("%v", r)

}

func countMatches(items [][]string, ruleKey string, ruleValue string) (n int) {
	map1 := map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}
	t1 := map1[ruleKey]
	for _, v := range items {
		if v[t1] == ruleValue {
			n++
		}
	}
	return
}
