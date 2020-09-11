package main

import (
	"fmt"
)

func encontraElemento(arr [10]int, elemento int) bool {
	for i := range arr {
		if arr[i] == elemento {
			return true
		}
	}
	return false
}

func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var tentativa bool = encontraElemento(arr, 10)
	fmt.Println(tentativa)
	var tentativa2 bool = encontraElemento(arr, 9)
	fmt.Println(tentativa2)
}
