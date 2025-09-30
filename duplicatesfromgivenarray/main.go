// 3) Remove Duplicates from given Array

// Example 1:
// Input: 1,2,3,2,5,4,1,4,5
// Output: 1,2,3,5,4

// Example 2:
// Input: 1,5,8,6,5,1,2,4,2,4,8
// Output: 1,5,8,6,2,4

package main

import "fmt"

func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)

	result := []int{}

	for _, num := range arr {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 2, 5, 4, 1, 4, 5}

	res := removeDuplicates(arr)

	fmt.Println("res", res)
}
