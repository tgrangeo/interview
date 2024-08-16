package main

import "fmt"

func contains(s []int, toSearch int) bool {
	for _, x := range s {
		if x == toSearch {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(contains([]int{11, 1, 1, 69, 22, 22, 11, 2, 3, 5, 4}, 4))
	fmt.Println(contains([]int{11, 1, 1, 69, 22, 22, 11, 2, 3, 5, 4}, 41))
}
