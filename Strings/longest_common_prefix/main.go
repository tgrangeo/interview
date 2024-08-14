package main

import "fmt"

func longestCommonProfile(str1 string, str2 string) string{
	if len(str2) == 0 || len(str1) == 0 {
		return ""
	}
	i := 0
	for i < len(str1) && i < len(str2) && str1[i] == str2[i] {
		i++
	}
	return str1[:i]
}

func main(){
	fmt.Println(longestCommonProfile("hello","hell"))
	fmt.Println(longestCommonProfile("hello","bye"))
}