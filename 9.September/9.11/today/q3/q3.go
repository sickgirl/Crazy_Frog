package main

import (
	"fmt"
	"strings"
)

//1108. IP 地址无效化
//https://leetcode.cn/problems/defanging-an-ip-address/

func main() {
	ip := "1.1.1.1"
	ipv1 := defangIPaddr(ip)
	fmt.Printf("%v", ipv1)
}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
