package main

import (
	"fmt"
)

func printaInt(n int) {
	if n > 0 {
		fmt.Println(n)
		printaInt(n - 1)
	}
}
