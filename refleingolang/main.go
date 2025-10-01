package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Type", runtime.NumGoroutine())
	s := "hii"
	fmt.Println(s)
}
