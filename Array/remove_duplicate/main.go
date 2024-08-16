package main

import (
	"fmt"
	"slices"
)

func removeDuplicate(arr []int) []int {
	for i := range arr {
		if i+1 < len(arr) && arr[i] == arr[i+1] {
			arr = slices.Delete(arr, i, i+1)
		}
	}
	return arr
}

func main() {
	s := []int{11, 1,1,69,22,22, 11, 2, 3, 5, 4}
	slices.Sort(s)
	fmt.Println(removeDuplicate(s))
}
