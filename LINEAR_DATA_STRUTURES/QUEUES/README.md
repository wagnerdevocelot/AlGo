# QUEUES

Uma *queue* consiste em elementos a serem processados em uma ordem específica ou com base na prioridade. 

Operações como _enqueue_ (entra na fila), _dequeue_  (sai da fila) e _peek_ (observar fila) podem ser executadas na *queue*. 
A *queue* é um estrutura de dados linear e uma coleção sequencial. 

Os elementos são adicionados ao final e são removido do início da coleção. As *queues* são comuns para armazenar tasks, ou requisições HTTP que precisam ser processadas por um servidor.
Tratamento de interrupções em real-time systems, call handling e scheduling de tasks de CPU
são bons exemplos de uso de *queues*.

Queues podem ser implementadas de formas diferentes, isso vai depender do propósito.


## ROLL THE DICE

Um bom exemplo de como funciona uma *queue* é pensar num jogo de rpg, geralmente quando vamos iniciar uma batalha precisamos jogar os dados e definir através do resultado quem tem mais pontos de iniciativa, e baseado na iniciativa sabemos a ordem de cada personagem atuar.
Faz um bom tempo que não jogo então não lembro tão bem, mas acho que a iniciativa é baseada em quem tirar o menor número no dado, mas talvez isso possa variar.

Então o primeiro na *fila* é o primeiro a atuar e também o primeiro a sair da *fila* dando lugar para o próximo. O primeiro da *fila* é o primeiro a sair dela. First In First Out (_FIFO_).

Já em rpgs eletrônicos com batalhas de turno se define o primeiro da *fila* como aquele que tem o maior numero no stat de agility, speed ou move.

### Personagens

Vou dar um exemplo de uma *queue* bem simples, aqui vou definir os dados de personagens do rpg.

```go
type personagem struct {
	nome string
	mov  int
}
```
### Queue

Essa aqui é a *queue*, vou usar um slice mas queues são implementadas de diferentes formas, essa é para se ter uma introdução.

```go
type queue struct {
	slice []personagem
}
```
### Ordenando o Turno

Esse é um *sort* customizado que é possível de ser feito em Go, ele não é necessário para se fazer uma *queue*, mas nesse exemplo eu quero utilizar o *sort* para ordenar um slice de personagens baseado no atributo de movimento, fazendo com que o personagem que tenha o maior nível nesse atributo possa entrar na fila primeiro.

Geralmente eu não gosto de usar parâmetros como (i, j, a, k) mas como esse sort vem da documentação padrão então vou deixá-lo como na documentação, pois ele não é o foco nesse caso.
 
```go
type primeiroDaFila []personagem

func (a primeiroDaFila) Len() int           { return len(a) }
func (a primeiroDaFila) Less(i, j int) bool { return a[i].mov > a[j].mov }
func (a primeiroDaFila) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
```

### FIFO enqueue/dequeue

Esses aqui são os métodos que mencionei no inicio *enqueue* insere um personagem na queue, *dequeue* retira o primeiro personagem da *queue*. 

Caso não tenha familiaridade com Go talvez você tenha estranhado essa parte (fila.slice[1:len(fila.slice)]) essa é uma operação de _slicing_, [1:] significa: "eu quero tudo que está do index 1 até, o limite do slice = len(fila.slice)" Excluindo o index 0.

O *dequeue* não leva parametros justamente porque a ideia da queue não é escolher quem retirar da fila então só retornamos o personagem removido.

```go
func (fila *queue) enqueue(nome personagem) {

	fila.slice = append(fila.slice, nome)
}

func (fila *queue) dequeue() personagem {

	removido := fila.slice[0]
    fila.slice = fila.slice[1:len(fila.slice)]
    
	return removido
}
```

### Iniciando o Turno

Agora na func main é hora de fazer isso tudo funcionar, criamos uma instância de personagens e uma fila, em seguida usamos a função de ordenação para que essa lista fique na ordem em que o personagem com maior nível de movimento esteja na primeira posição.

Fazemos uma iteração para inserir na fila todos os personagens que foram ordenados.


```go
func main() {

	equipe := []personagem{
		{"Cloud Strife", 8},
		{"Barret Wallace", 4},
		{"Tifa Lockhart", 9},
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
```

Aqui tem a primeira saída mostrando a fila ordenada. 

```text
&{[{Argath 5} {Saravasti 4} {Athena 4}]}
```

Uma chamada para o método dequeue, para que o primeiro personagem após realizar sua ação dê seu lugar ao próximo personagem com maior atributo de movimento, e a saída mostrando a fila depois do dequeue.

```text
&{[{Saravasti 4} {Athena 4}]}
```
