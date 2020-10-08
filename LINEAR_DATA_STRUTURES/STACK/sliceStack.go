package main

import "fmt"

type pagina struct {
	nome string
	url  string
}

type stack struct {
	slice []pagina
}

func main() {

	historico := []pagina{
		{"GeeksforGeeks", "https://www.geeksforgeeks.org/"},
		{"Youtube", "https://www.youtube.com/"},
		{"Github", "https://github.com/"},
		{"Twitter", "https://twitter.com/home"},
	}

	var chrome *stack = new(stack)

	for i := range historico {
		chrome.avancar(historico[i])
	}

	fmt.Println(chrome)
	chrome.voltar()
	fmt.Println(chrome)

	fmt.Println(chrome.vazia())

}

func (navegador *stack) avancar(home pagina) {

	navegador.slice = append(navegador.slice, home)
}

func (navegador *stack) voltar() bool {

	if navegador.vazia() {
		return false
	}

	tamanho := len(navegador.slice) - 1
	navegador.slice = navegador.slice[:tamanho]

	return true
}

func (navegador *stack) vazia() bool {
	return navegador.slice == nil
}
