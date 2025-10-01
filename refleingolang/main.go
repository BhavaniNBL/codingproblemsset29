package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime", runtime.NumGoroutine())
}
