package main

import (
	"fmt"
	"math"
)

func reverseInteger(num int) int {
	reversed := 0
	sign := 1
	if num < 0 {
		sign = -1
		num = -num
	}
	for num != 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}
	reversed *= sign
	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}
	return reversed
}

func main() {
	fmt.Println(reverseInteger(12345))     // Output: 54321
	fmt.Println(reverseInteger(-12345))    // Output: -54321
	fmt.Println(reverseInteger(120))       // Output: 21
	fmt.Println(reverseInteger(0))         // Output: 0
}
