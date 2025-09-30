package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) (string, int) {
	if len(s) == 0 {
		return "", 0
	}

	maxLen := 0
	longest := ""

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			sub := s[i : j+1]
			if isPalindrome(sub) && len(sub) > maxLen {
				longest = sub
				maxLen = len(sub)
			}
		}
	}

	return longest, maxLen
}

func main() {
	str := "abageafd"
	substr, length := longestPalindrome(str)
	fmt.Println(substr, length)

}
