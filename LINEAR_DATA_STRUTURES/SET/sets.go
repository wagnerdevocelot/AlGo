package main

import (
	"fmt"
)

//Set class
type Set struct {
	estrutura map[int]bool
}

//Criar faz um map de inteiros e boleanos
func (set *Set) Criar() {
	set.estrutura = make(map[int]bool)
}

// AddValor adiciona um elemento no set
func (set *Set) AddValor(valor int) {

	if !set.PossuiValor(valor) {
		set.estrutura[valor] = true
	}
}

//DeletaValor do set
func (set *Set) DeletaValor(valor int) {

	delete(set.estrutura, valor)
}

//PossuiValor verifica se o set possui um valor especifico
func (set *Set) PossuiValor(valor int) bool {

	var localizado bool

	_, localizado = set.estrutura[valor]

	return localizado
}

//CruzarCom metodo retorna set com a interseção feita em um setDiferente
func (set *Set) CruzarCom(setDiferente *Set) *Set {

	var setCruzado = &Set{}
	setCruzado.Criar()
	var valor int

	for valor, _ = range set.estrutura {

		if setDiferente.PossuiValor(valor) {

			setCruzado.AddValor(valor)
		}
	}

	return setCruzado
}

//UnirCom retorna a junção de dois sets
func (set *Set) UnirCom(setDiferente *Set) *Set {

	var setReunido = &Set{}
	setReunido.Criar()
	var valor int

	for valor, _ = range set.estrutura {
		setReunido.AddValor(valor)
	}

	for valor, _ = range setDiferente.estrutura {
		setReunido.AddValor(valor)
	}

	return setReunido
}

func main() {

	var set *Set
	set = &Set{}

	set.Criar()

	set.AddValor(1)
	set.AddValor(2)

	fmt.Println("set inicial", set)
	fmt.Println(set.PossuiValor(1))

	var setDiferente *Set
	setDiferente = &Set{}

	setDiferente.Criar()
	setDiferente.AddValor(2)
	setDiferente.AddValor(4)
	setDiferente.AddValor(5)

	fmt.Println("outro set", set)

	fmt.Println("cruzamento de sets", set.CruzarCom(setDiferente))

	fmt.Println("união de sets", set.UnirCom(setDiferente))

}
