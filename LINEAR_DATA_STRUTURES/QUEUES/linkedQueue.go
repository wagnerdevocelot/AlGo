package main

import "fmt"

// Node define o node da linkedlist
type Node struct {
	valor   interface{}
	proximo *Node
}

// LinkedListQueue define a fila baseada em linkedlist
type LinkedListQueue struct {
	inicio  *Node
	final   *Node
	tamanho int
}

// NovaLinkedListQueue cria uma nova fila
func NovaLinkedListQueue() *LinkedListQueue {

	return &LinkedListQueue{
		inicio:  nil,
		final:   nil,
		tamanho: 0,
	}
}

// Enqueue insere um elemento no final da fila
func (fila *LinkedListQueue) Enqueue(v interface{}) {

	node := &Node{
		valor:   v,
		proximo: nil,
	}

	if fila.final == nil {

		fila.final = node
		fila.inicio = node

	} else {

		fila.final.proximo = node
		fila.final = node
	}

	fila.tamanho++
}

// Dequeue remove o primeiro item da fila
func (fila *LinkedListQueue) Dequeue() interface{} {

	if fila.inicio == nil {
		return nil
	}

	v := fila.inicio.valor
	fila.inicio = fila.inicio.proximo
	fila.tamanho--

	return v
}

// String retorna o estado da lista
func (fila *LinkedListQueue) String() string {

	if fila.inicio == nil {
		return "fila vazia"
	}
	resultado := "inicio"

	for atual := fila.inicio; atual != nil; atual = atual.proximo {

		resultado += fmt.Sprintf("<-%+v", atual.valor)
	}
	resultado += "<-final"

	return resultado
}
