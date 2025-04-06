# Doubly Linked List

![](https://cdn-images-1.medium.com/max/800/1*jN7zAXqMC91AzCOgkD17Hg.png)

Em uma Doubly Linked List, (que a partir de agora vou chamar de DLL) todos os nodes têm um ponteiro para o node ao qual estão conectados.
Isso significa que cada node está conectado a dois nodes, e podemos avançar para o próximo node ou retroceder até o node anterior.

![](https://cdn-images-1.medium.com/max/800/1*kOYbVWYeEnxGz6JQdlxWHg.png)

DLL permitem operações de insert, deleting e, obviamente, de traversing.

E Para fins de manter o exemplo de linhas temos agora o que seria uma representação de estações de trem.
Enquanto que a single linked list representa um monotrilho aqui temos duas ligações.

Quem nunca voltou da paulista mais loco e que o coringa chegou na consolação e se perguntou "o meu é sentido vila madalena ou é sentido vila prudente? E o sentido da vida?"

![](https://cdn-images-1.medium.com/max/800/1*30q1_bNCMOhlEAR5odi4Ng.png)

Listas duplamente ligadas são exatamente iguais a estações de metrô pois através de um Node você pode seguir para o próximo ou para o anterior  pois temos ponteiros nas duas direções.

```go
// Estacao struct = Node
type Estacao struct {
	// property
	nome string
	// nextNode
	estacaoSeguinte *Estacao
	// previousNode
	estacaoAnterior *Estacao
}
```

Assim como na single LinkedList temos o mesmo conceito de head e tail então a struct é exatamente igual.

``` go
// Metro = Doubly linked list
type Metro struct {
	// head
	inicioDaLinha *Estacao
	// tail
	finalDaLinha *Estacao
}
```

## Node Between Values

O método EntreDuasEstacoes do Metro retorna a Estacao que tem um nome situado entre os valores anterior e proxima. O método percorre a lista para descobrir se as strings anterior e proxima correspondem em Nodes consecutivos.

```go
func (linhaVerde *Metro) EntreDuasEstacoes(anterior string, proxima string) *Estacao {

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
```

Primeiro instanciamos o Node fazemos um for para percorrer os itens do head ao tail enquanto estacao for diferente de nil. Em seguida uma comparação com os parâmetros anterior e proxima feitos com o tail e head para reconhecer o node entre eles, sendo que a comparação é feita somente se head e tail forem diferentes de nil.

## The AddToHead method

O método AddInicioDaLinha define o head para que possamos seguir construindo mais estacões assim como fizemos com o monotrilho.

``` go
func (linhaVerde *Metro) AddInicioDaLinha(novaEstacao string) {

	var estacao = &Estacao{}
	estacao.nome = novaEstacao
	estacao.estacaoSeguinte = nil

	if linhaVerde.inicioDaLinha != nil {

		estacao.estacaoSeguinte = linhaVerde.inicioDaLinha
		linhaVerde.inicioDaLinha.estacaoAnterior = estacao
	}
	linhaVerde.inicioDaLinha = estacao
}
```

## AddAfter method

Aqui fazemos um insert de um node após outro node que é passado como parâmetro presente na lista, e para saber se o Node está presente reutilizamos o método ProcuraEstacao() que havíamos feito para a Single Linkedlist.

```go
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
```
Fazemos a instância do Node atribuímos o nome usado como parâmetro e setamos o node seguinte como nil para que se mantenha o conceito de tail.

Fazemos a busca e recuperamos a localização do Node de referência como estacaoAtual.

Então a estacao seguinte do node atual para a ser a estacao seguinte do node recuperado (mindfuck)
Parece um malabarismo de valores e é na verdade isso mesmo, talvez por isso pareça complicado mas basta trackear onde estão os valores e prestar atenção por onde eles estão passando.

## The AddToEnd method

AddEstacaoNoFinalDaLinha cujo o nome é bastante descritivo, faz uma nova instância de um node e reutiliza o método UltimaEstacao() para recuperar o ultimo node e passar o valor do Node instanciado como referencia através de ponteiro.

```go
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
```

## Main Func

Aqui vou fazer o mesmo processo que usei com o monotrilho pra popular a DLL, e além desses métodos no arquivo doubly.go no repositório tem mais exemplos de métodos para DLL e os métodos reaproveitados da Single Linked List.


```go
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
```

Output:

```text
Vila Prudente
Tamanduateí
Sacomã
Alto do Ipiranga
Santos-Imigrantes
Chácara Klabin
Ana Rosa
Paraíso
Brigadeiro
Trianon-MASP
Consolação
Clínicas
Sumaré
Vila Madalena
```

## Doubly vs Single

**Vantagens sobre a single linked list**

- Uma DLL pode ser percorrida de trás para frente e vice versa.

- A operação de delete na DLL é mais eficiente se o ponteiro para o node excluído for fornecido.

- Podemos usar insert mais rapidamente em referencia a um item tanto a frente quanto trás.

- Na SLL (single linked list), para excluir um node, é necessário um ponteiro para o node anterior. Para obter esse node anterior, às vezes a lista precisa ser percorrida. Na DLL, podemos obter o node anterior usando o ponteiro anterior.

**Desvantagens sobre a singly linked list**

- Cada node da DLL requer espaço extra para um ponteiro anterior. (ainda assim é possível criar uma DLL sem um segundo ponteiro.)

- Todas as operações requerem um ponteiro extra para ser mantida. Por exemplo, no insert, precisamos modificar os ponteiros anteriores junto com os próximos ponteiros.

