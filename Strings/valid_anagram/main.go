package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func isAnagram(s1, s2 string) bool {
	return sortString(s1) == sortString(s2)
}

func main() {
	fmt.Println(isAnagram("Hello, 世界", "世界, Hello"))
	fmt.Println(isAnagram("yeyeyeey", "nononon"))
	fmt.Println(isAnagram("ieeench", "chieeen"))
}
