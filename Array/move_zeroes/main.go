package main

import (
	"fmt"
)

func moveZeroes(s []int) []int {
	lastNonZeroIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] != 0 {
			s[lastNonZeroIndex], s[i] = s[i], s[lastNonZeroIndex]
			lastNonZeroIndex++
		}
	}
	return s
}

func main() {
	s := []int{0,0,67, 0, 1, 0, 2, 0, 3}
	fmt.Println(moveZeroes(s))
}
