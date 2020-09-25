package main

import "fmt"

// Estacao struct = Node
type Estacao struct {
	// property
	nome string
	// nextNode
	estacaoSeguinte *Estacao
	// previousNode
	estacaoAnterior *Estacao
}

// Metro = Doubly linked list
type Metro struct {
	// head
	inicioDaLinha *Estacao
	// tail
	finalDaLinha *Estacao
}

func main() {
	var linhaVerde Metro
	linhaVerde = Metro{}

	estacoes := []string{"Tamanduateí", "Sacomã", "Alto do Ipiranga", "Santos-Imigrantes", "Chácara Klabin",
		"Ana Rosa", "Paraíso", "Brigadeiro", "Trianon-MASP", "Consolação", "Clínicas",
		"Sumaré", "Vila Madalena",
	}

	linhaVerde.AddinicioDaLinha("Vila Prudente")
	for i := range estacoes {
		linhaVerde.AddEstacaoNoFinalDaLinha(estacoes[i])
	}

	linhaVerde.ListarEstacoes()
}

//EntreDuasEstacoes method
func (linhaVerde *Metro) EntreDuasEstacoes(anterior, proxima string) *Estacao {

	var estacao *Estacao
	var estacaoAtual *Estacao

	for estacao = linhaVerde.inicioDaLinha; estacao != nil; estacao = estacao.estacaoSeguinte {

		if estacao.estacaoAnterior != nil && estacao.estacaoSeguinte != nil {

			if estacao.estacaoAnterior.nome == anterior && estacao.estacaoSeguinte.nome == proxima {
				estacaoAtual = estacao
				break
			}
		}
	}
	return estacaoAtual
}

//AddinicioDaLinha method
func (linhaVerde *Metro) AddinicioDaLinha(novaEstacao string) {

	var estacao = &Estacao{}
	estacao.nome = novaEstacao
	estacao.estacaoSeguinte = nil

	if linhaVerde.inicioDaLinha != nil {

		estacao.estacaoSeguinte = linhaVerde.inicioDaLinha
		linhaVerde.inicioDaLinha.estacaoAnterior = estacao
	}
	linhaVerde.inicioDaLinha = estacao
}

//ProcuraEstacao method returns Node given parameter property
func (linhaVerde *Metro) ProcuraEstacao(destino string) *Estacao {

	var estacao *Estacao
	var estacaoEncontrada *Estacao

	for estacao = linhaVerde.inicioDaLinha; estacao != nil; estacao = estacao.estacaoSeguinte {

		if estacao.nome == destino {
			estacaoEncontrada = estacao
			break
		}
	}
	return estacaoEncontrada
}

//AddEstacaoSeguinte method of LinkedList
func (linhaVerde *Metro) AddEstacaoSeguinte(destino string, novaEstacao string) {
	var estacao = &Estacao{}
	estacao.nome = novaEstacao
	estacao.estacaoSeguinte = nil

	var estacaoAtual *Estacao
	estacaoAtual = linhaVerde.ProcuraEstacao(destino)

	if estacaoAtual != nil {

		estacao.estacaoSeguinte = estacaoAtual.estacaoSeguinte
		estacao.estacaoAnterior = estacaoAtual
		estacaoAtual.estacaoSeguinte = estacao
	}
}

//UltimaEstacao method returns the last Node
func (linhaVerde *Metro) UltimaEstacao() *Estacao {

	var estacao *Estacao
	var fimDaLinha *Estacao

	for estacao = linhaVerde.inicioDaLinha; estacao != nil; estacao = estacao.estacaoSeguinte {

		if estacao.estacaoSeguinte == nil {
			fimDaLinha = estacao
		}
	}
	return fimDaLinha
}

//AddEstacaoNoFinalDaLinha method of LinkedList
func (linhaVerde *Metro) AddEstacaoNoFinalDaLinha(novaEstacao string) {
	var estacao = &Estacao{}
	estacao.nome = novaEstacao
	estacao.estacaoSeguinte = nil

	var fimDaLinha *Estacao
	fimDaLinha = linhaVerde.UltimaEstacao()

	if fimDaLinha != nil {

		fimDaLinha.estacaoSeguinte = estacao
		estacao.estacaoAnterior = fimDaLinha
	}
}

//ListarEstacoes method iterates over LinkedList
func (linhaVerde *Metro) ListarEstacoes() {

	var estacao *Estacao

	for estacao = linhaVerde.inicioDaLinha; estacao != nil; estacao = estacao.estacaoSeguinte {
		fmt.Println(estacao.nome)
	}
}

// RemoverEstacao method
func (linhaVerde *Metro) RemoverEstacao(estacaoRemovida string) bool {

	if linhaVerde.inicioDaLinha == nil {
		return false
	}

	if linhaVerde.inicioDaLinha.nome == estacaoRemovida {

		linhaVerde.inicioDaLinha = linhaVerde.inicioDaLinha.estacaoSeguinte
		linhaVerde.inicioDaLinha.estacaoAnterior = nil
		return true
	}

	if linhaVerde.finalDaLinha.nome == estacaoRemovida {

		linhaVerde.finalDaLinha = linhaVerde.finalDaLinha.estacaoAnterior
		linhaVerde.finalDaLinha.estacaoSeguinte = nil
		return true
	}

	estacaoAtual := linhaVerde.inicioDaLinha

	for estacaoAtual.estacaoSeguinte != nil {

		if estacaoAtual.estacaoSeguinte.nome == estacaoRemovida {

			if estacaoAtual.estacaoSeguinte.estacaoSeguinte != nil {
				estacaoAtual.estacaoSeguinte.estacaoSeguinte.estacaoAnterior = estacaoAtual
			}
			estacaoAtual.estacaoSeguinte = estacaoAtual.estacaoSeguinte.estacaoSeguinte
			return true
		}
		estacaoAtual = estacaoAtual.estacaoSeguinte
	}
	return false
}
