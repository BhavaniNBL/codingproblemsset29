package main

import "fmt"

func fib(n int) []int {
	if n <= 0 {
		return []int{}
	}

	if n == 1 {
		return []int{1}
	}

	fibSeq := make([]int, n)
	fibSeq[0] = 0
	fibSeq[1] = 1
	for i := 2; i < n; i++ {
		fibSeq[i] = fibSeq[i-1] + fibSeq[i-2]
	}

	return fibSeq

}

func main() {
	length := 10

	res := fib(length)

	fmt.Println("res", res)
}
