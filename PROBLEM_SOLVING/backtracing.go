package main

import (
	"fmt"
)

func combinacaoDeSoma(arr [10]int, combinacoes [19]int, tamanho int, chave int, addValor int, l int, m int) int {
	numero := 0

	if addValor > chave {
		return -1
	}

	if addValor == chave {
		numero = numero + 1
		for p := 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinacoes[p]])
		}
		fmt.Println(" ")
	}

	for i := l; i < tamanho; i++ {
		combinacoes[m] = l
		combinacaoDeSoma(arr, combinacoes, tamanho, chave, addValor+arr[i], l, m+1)
		l = l + 1
	}
	return numero
}
