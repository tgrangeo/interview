package main

import "fmt"

func twoSum(s []int, target int) (int, int) {
	for i := range s {
		y := 0
		for y = i; y < len(s); y++ {
			if s[i]+s[y] == target {
				return s[i], s[y]
			}
		}
	}
	return 0, 0
}

func main() {
	s := []int{1, 2, 3, 4, 41}
	fmt.Println(twoSum(s, 42))
}
