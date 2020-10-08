package main

import (
	"strconv"
)

//Componenete struct
type Componenete struct {
	unidade int
}

// Pilha é um LIFO stack dinâmico
type Pilha struct {
	componentes  []*Componenete
	ContaUnidade int
}

// String conversão de componente para string
func (componente *Componenete) String() string {

	return strconv.Itoa(componente.unidade)
}

// New retorna uma nova stack
func (pilha *Pilha) New() {
	pilha.componentes = make([]*Componenete, 0)
}

// Push adiciona um novo componente a stack
func (pilha *Pilha) Push(componente *Componenete) {

	pilha.componentes = append(pilha.componentes[:pilha.ContaUnidade], componente)
	pilha.ContaUnidade = pilha.ContaUnidade + 1
}

// Pop remove e retorna o componente no top na stack
func (pilha *Pilha) Pop() *Componenete {

	if pilha.ContaUnidade == 0 {
		return nil
	}

	var tamanho int = len(pilha.componentes)
	var componente *Componenete = pilha.componentes[tamanho-1]

	if tamanho > 1 {

		pilha.componentes = pilha.componentes[:tamanho-1]
	} else {

		pilha.componentes = pilha.componentes[0:]
	}

	pilha.ContaUnidade = len(pilha.componentes)

	return componente
}
