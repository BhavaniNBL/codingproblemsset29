package main

import "fmt"

func validInputStr(s string) bool {
	bracketMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	fmt.Println(bracketMap, "bracketMap")

	stack := []rune{}

	for _, char := range s {
		if closing, ok := bracketMap[char]; ok {
			stack = append(stack, closing)
		} else {
			fmt.Println("stack", stack)
			fmt.Println("char ", char, "stack length", len(stack)-1)
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
