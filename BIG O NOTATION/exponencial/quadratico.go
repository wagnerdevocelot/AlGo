package main

import (
	"fmt"
)

func printaTodosOsPares(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println(i, j)
		}
	}
}
