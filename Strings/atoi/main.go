package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func checkEntry(str string) error {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return errors.New("input is empty")
	}
	if str[0] == '-' {
		str = str[1:]
	}
	if len(str) == 0 {
		return errors.New("input only contains a '-'")
	}
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return errors.New("input contains non-digit characters")
		}
	}
	return nil
}

func Atoi(str string) (int, error) {
	str = strings.TrimSpace(str)
	err := checkEntry(str)
	if err != nil {
		return 0, err
	}
	sign := 1
	if str[0] == '-' {
		sign = -1
		str = str[1:]
	}
	res := 0
	for _, c := range str {
		res = res*10 + int(c-'0')
	}
	return res * sign, nil
}

func main() {
	if result, err := Atoi("12"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	if result, err := Atoi("-122"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	if result, err := Atoi("12a3"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
