package main

import "fmt"

func singleNumber(s []int) int {
	result := 0
	for _, num := range s {
		result ^= num
	}
	return result
}

func main() {
	s := []int{1, 2, 1, 4, 4, -12, 12, 12, 2, 3, 3}
	fmt.Println(singleNumber(s))
}
