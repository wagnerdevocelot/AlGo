package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 10, 40}
	item := 9
	busca := buscaBinaria(arr, 0, len(arr), item)
	fmt.Println(busca)
}

func buscaBinaria(arr []int, esquerda, direita, item int) bool {
	for esquerda <= direita {
		meio := esquerda + (direita-esquerda)/2

		if arr[meio] == item {
			return true
		}

		if arr[meio] < item {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}
	return false
}
