package main

import (
	"fmt"
)

func printaNum(n int) {
	if n > 0 {
		printaNum(n - 1)
		fmt.Println(n)
	}
}
