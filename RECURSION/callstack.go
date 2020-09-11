package main

import "fmt"

func main() {
	acorda()
}

func tomarBanho() string {
	return "xuaxua"
}

func tomarCafeDaManha() string {
	refeicao := fazComida()
	return refeicao
}

func fazComida() string {
	return "ovo frito e café"
}

func acorda() {
	tomarBanho()
	tomarCafeDaManha()
	fmt.Println("Ok bora trabalhar")
}
