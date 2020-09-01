package main

import (
	"fmt"
)

func linear(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total++
	}
	return total
}

func main() {
	fmt.Println(linear(10))
}
