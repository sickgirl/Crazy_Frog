package main

import (
	"fmt"
	"strings"
)

//1678. 设计 Goal 解析器
//https://leetcode.cn/problems/goal-parser-interpretation/
func main() {
	command := "G()(al)"
	r := interpret(command)
	fmt.Printf("%v", r)
}
func interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	command = strings.ReplaceAll(command, "(al)", "al")
	return command
}
