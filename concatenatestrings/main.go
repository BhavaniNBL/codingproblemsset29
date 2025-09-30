// 4) Concatenate any 2 strings to give smallest string  from given array of srings

// Example:
// Input: S=[“aab”,”bcddbc”,”aa”,”aazef”]
// Output: “aabaa”

package main

import "fmt"

func concatenateStrings(arr []string) string {

	minStr := arr[0] + arr[1]
	minLen := len(minStr)

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if i != j {
				concat := arr[i] + arr[j]
				fmt.Println(concat)
				if len(concat) < minLen {
					minLen = len(concat)
					minStr = concat
				}
			}
		}
	}
	return minStr
}

func main() {
	arr := []string{"aab", "bcddbc", "aa", "aazef"}

	res := concatenateStrings(arr)

	fmt.Println(res)
}
