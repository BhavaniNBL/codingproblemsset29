package main

import "fmt"

func validInputStr(s string) bool {
	bracketMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []rune{}

	for _, char := range s {
		if closing, ok := bracketMap[char]; ok {
			stack = append(stack, closing)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	s := "([])"
	res := validInputStr(s)
	fmt.Println(res)
}
