package main

import (
	"log"
)

func palindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	valid := "kayak"
	invalid := "qwerty"
	log.Println(palindrome(valid))   // true
	log.Println(palindrome(invalid)) // false
}
