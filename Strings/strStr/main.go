package main

import "fmt"

func strStr(str string, sub string) int {
	if len(sub) == 0 {
		return 0
	}
	for i := 0; i <= len(str)-len(sub); i++ {
		j := 0
		for j < len(sub) && str[i+j] == sub[j] {
			j++
		}
		if j == len(sub) {
			return i
		}
	}
	return -1 
}

func main() {
	fmt.Println(strStr("hello", "lo"))
	fmt.Println(strStr("hello", "loca"))
}
