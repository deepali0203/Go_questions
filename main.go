package main

import (
	"fmt"
)

func main() {
	s := "XX"
	fmt.Println(romanToIntStr(s))
}

func romanToIntStr(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	length := len(s)
	last_element := length - 1
	for i := 0; i < last_element; i++ {
		current := m[string(s[i])]
		next := m[string(s[i+1])]
		if current < next {
			result -= current
		} else {
			result += current
		}
	}
	result += m[string(s[last_element])]
	return result
}
