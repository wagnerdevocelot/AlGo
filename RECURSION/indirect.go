package main

import (
	"fmt"
)

func printaUm(n int) {
	if n >= 0 {
		fmt.Println("Na primeira func:", n)
		printaDois(n - 1)
	}
}

func printaDois(n int) {
	if n >= 0 {
		fmt.Println("Na segunda func:", n)
		printaUm(n - 1)
	}
}
