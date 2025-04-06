package main

import (
	"fmt"
	"sort"
)

type personagem struct {
	nome string
	mov  int
}

type queue struct {
	slice []personagem
}

type primeiroDaFila []personagem

func (a primeiroDaFila) Len() int           { return len(a) }
func (a primeiroDaFila) Less(i, j int) bool { return a[i].mov > a[j].mov }
func (a primeiroDaFila) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	equipe := []personagem{
		{"Saravasti", 4},
		{"Athena", 4},
		{"Argath", 5},
	}

	var fila *queue = new(queue)

	sort.Sort(primeiroDaFila(equipe))

	for i := range equipe {
		fila.enqueue(equipe[i])
	}

	fmt.Println(fila)
	fila.dequeue()
	fmt.Println(fila)

}

func (fila *queue) enqueue(nome personagem) {

	fila.slice = append(fila.slice, nome)
}

func (fila *queue) dequeue() personagem {

	removido := fila.slice[0]
	fila.slice = fila.slice[1:len(fila.slice)]

	return removido
}
