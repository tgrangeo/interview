package main

import "fmt"

func rotateArray(s []int)[]int{
	res := []int{}
	for i := len(s) - 1; i >= 0 ; i--{
		res = append(res, s[i])
	}
	return res
}

func main(){
	s := []int{1,2,3,4,5}
	fmt.Println(rotateArray(s))
}