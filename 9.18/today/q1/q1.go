package main

import "fmt"

func main() {
	s := "abcdefgg"
	r := repeatedCharacter(s)
	fmt.Printf("%v", string(r))
}

func repeatedCharacter(s string) byte {
	map1 := make(map[byte]bool)
	for _, v := range s {
		if _, ok := map1[byte(v)]; ok {
			return byte(v)
		}
		map1[byte(v)] = true
	}
	return 'a'
}
