package main

import "fmt"

func plusOne(s []int) []int {
    i := len(s) - 1
    for i >= 0 && s[i] == 9 {
        s[i] = 0
        i--
    }
    if i >= 0 {
        s[i] += 1
    } else {
        s = append([]int{1}, s...)
    }
    return s
}

func main() {
    s := []int{1, 2, 9, 9}
    fmt.Println(plusOne(s))
    
    ss := []int{9}
    fmt.Println(plusOne(ss))
}
