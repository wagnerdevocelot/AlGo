# STACK

Uma Stack é a estrutura de "último a entrar, primeiro a sair" (last in last out), na qual os itens são adicionados a partir do topo. As stacks são usadas em parsers para resolver algoritmos de labirinto. Push, pop, top e get size são as operações típicas permitidas nas estruturas de dados da stack. Parsing de sintaxe, backtracking e gerenciamento de memória de tempo de compilação são alguns cenários da vida real onde as stacks podem ser usadas.

Nesse exemplo usando a dinâmica de um web browser vamos implementar o que seria o seistema de avançar e voltar nas paginas.

Como o ultimo item a entrar é o primeiro a sair significa que vamos gurdar URL's em uma pilha e cada vez que quisermos voltar para a página anterior seria como uma função de pop para retirar da stack o ultimo item, que seria a URL atual.

Exemplo: 

Eu visiitei o "Google, Facebook, Twitter e Youtube". Então tenho uma lista onde a posição 0 é o Google e a ultima osição é o Youtube, o que seria como a URL atual e também o ultimo item que entrou na stack.

Então se eu quiser voltar na navegação do browser significa que essa operação contecerá:

Browser.pop

E o estado atual da lista seria "Google, Facebook, Twitter"


# Data

A struct de página contem os dados das páginas como o nome que aparece do header da aba e o caminho relativo.
Com a struct de stack definimos o que seria a estrutura onde as paginas são armazenadas e de onde podem ser feitas as operações.

```go
type pagina struct {
	nome string
	url  string
}

type stack struct {
	slice []pagina
}
```

# The Push method

O método avançar que é uma representação do push adiciona uma nova página no topo da stack. O exemplo é muito simples e é praticamente o mesmo padrão usado para queues quando usamos slices.

```go
func (navegador *stack) avancar(home pagina) {

	navegador.slice = append(navegador.slice, home)
}
```

# Empty

Precisamos verificar se a stack está vazia antes de dar pop nela então vamos criar um método para isso.

```go
func (navegador *stack) vazia() bool {

	return navegador.slice == nil
}
```

# The Pop method

O método voltar na implementação da stack remove a última página do slice, esse seria analogo ao pop. 

Para isso usamos a mesma operação de slice que também foi feita quando usamos queues, porém aqui retiramos os itens desde a posição 0 até o tamanho do slice - 1, mesmo falando que eu quero fatiar o slice da posição 0 a 4 ele me devolve "0,1,2,3".

Também fazemos uma checagem usando o método vazio para saber se a lista contem intens antes de utilizar o pop.


```go
func (navegador *stack) voltar() bool {

	if navegador.vazia() {
		return false
	}

	tamanho := len(navegador.slice) - 1
	navegador.slice = navegador.slice[:tamanho]

	return true
}
```


# The main method

Aqui instaciamos as páginas como o histórico acessado e a nossa stack como o nagevador chrome.
Depois tem um for iterando sobre os itens no historico fazendo push dos itens do historico no chrome. Como se estivessemos acessando as páginas uma depois da outra.
Vem então a chamada do pop que retira o ultimo item e fazemos o print do estado atual da stack.

```go
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
```

Output:

```text
&{[{GeeksforGeeks https://www.geeksforgeeks.org/} {Youtube https://www.youtube.com/} {Github https://github.com/} {Twitter https://twitter.com/home}]}

&{[{GeeksforGeeks https://www.geeksforgeeks.org/} {Youtube https://www.youtube.com/} {Github https://github.com/}]}

false
```

A primeira stack com todos os itens e a segunda depois da operação de pop e em seguida a chamada de vazio para ver o estado da stack. Que retorna false já que nossa stack ainda está populada.