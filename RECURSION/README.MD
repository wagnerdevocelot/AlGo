### Recursions & Dragons

Recursão da forma mais simples que dá pra explicar

![](https://cdn-images-1.medium.com/max/800/1*RJ7-1FaFuVEVGHN3xTwFHw.jpeg)

## Intro

Sabe aquelas coisas que todo mundo diz, “isso é muito simples depois que você entende”.

E eu fico tipo??? Óbvio né kkkk

Esse é outro assunto que nós só entendemos depois que a ficha cai, **recursão**.

Junto com esse conteúdo uma outra parte já está no repositório sobre problem solving porém apenas esse será em formato de thread alguns padrões de problem solving envolvem o uso de recursão então é uma boa ideia ler os dois em sequência.

A recursão é antes de qualquer coisa uma forma diferente de pensar em soluções. A principio pode ser uma dor de cabeça para se acostumar, mas depois que vira uma segunda natureza você visualiza a programação de uma nova forma, meio que análogo aquele “click” que cheguei a comentar quando falei sobre Big O Notation.

Talvez seja desmotivante, mas se você ainda não sabe recursão, pode ser que leve até algumas semanas pra cair a ficha então é recomendado absorver muito conteúdo, não é um artigo que vai esclarecer todas as nossas dúvidas.

Bora de lado essa coisa de ficar comprando curso “Ultimate”, “Premium”, “Advanced”.

Parte do processo de estudar é aprender a procurar recursos interessantes e aprender com mais de um se possível ou viável.  
Ter muito mais do que apenas um ponto de vista sobre o mesmo assunto nos ajuda a ter o nosso próprio. Essa é uma dificuldade muito grande pra maioria visto que desde o nosso ensino fundamental todo o conteúdo é pré definido, e por isso existe essa ânsia de comprar o “melhor” curso ou o mais completo, porque a gente precisa do material pronto na nossa mão ninguém ensinou a gente a fazer pesquisa.

Pensando nesse caso tenho aqui um repositório interessante com boas dicas de como pesquisar melhor e tirar dúvidas sobre programação.

[Joga No Google](https://github.com/Caaddss/joga_no_google)

Se puder contribuir também vai ser muito bom.

Mesmo que termine esse artigo sem entender recursão e eu espero que não, não desista, corre atrás que tem muita coisa interessante na web sobre isso e quem sabe exista um conteúdo na medida.

#### Definindo recursão em três partes

  

A - **Repetição**, a recursão é uma repetição infinita. Como um loop, mas por debaixo dos panos a implementação é diferente.

B - **Endpoint**, a recursão precisa de uma resolução chamada geralmente de _caso base_ para que o algoritmo se resolva e a repetição termine.

C -**Operações**, aplicando a recursão sem nenhuma operação ele é inútil então é necessário algum tipo de lógica que defina a recursão para que durante a repetição das chamadas essa lógica se aplique e o algoritmo chegue ao caso base.

  

## Caverna do Dragão

  

Hank, Erick, Diana, Sheila, Presto e Bobby conversam com o mestre dos magos.

![](https://cdn-images-1.medium.com/max/800/1*_wpUTKcovPMSNnp5t2saTw.jpeg)

“Para achar o caminho de casa vocês terão que descobrir qual destas cinco chaves é a chave encantada que abre o portal para o seu mundo, se usarem a chave errada o portal nunca mais se abrirá” — diz o Mestre dos Magos

“Ta legal e como é que a gente vai saber qual é a chave que abre o portal?” — Pergunta Sheila.

“Para isso vocês devem ir ao cemitério dos dragões onde vive a deusa Tiamat e ela os ensinará recursividade” — Explica o Mestre dos Magos

“Eu é que não vou, até o vingador tem medo desse dragão!” — Reclama Erick

“Mas o que é recursividade e como isso ajuda a descobrir qual das chaves é a chave do portal?” — Pergunta Diana

O mestre dos magos não está mais lá.

Os jovens então se encaminham até o cemitério dos dragões para convencer Tiamat a ajudá-los. Logo na entrada a visão do dragão de cinco cabeças amedronta o grupo, mas curiosamente o dragão fala na linguagem humana e pergunta. — “Quem ousa invadir meu território?”

![](https://cdn-images-1.medium.com/max/800/1*-yCW5v2FVvD4TccD2KcyUg.jpeg)

Então Tiamat se aproxima e desce uma de suas cabeças a altura dos jovens, agora impressionados com o tamanho da criatura.

“Tiamat, o mestre dos magos nos enviou e nos disse que só você pode dizer qual das chaves é a chave do portal para o nosso mundo” — diz Hank apresentando uma caixa de 5 chaves enfileiradas.

```text
[]string{“chave1”, “chave2”, “chave3”, “chave4”, “chave5”}
```

“Entendo, eu geralmente não aceito pedidos, mas como o mestre dos magos os enviou direi apenas se a primeira chave dessa caixa é a chave do portal” — diz o Dragão.

O grupo reage assustado.

![](https://cdn-images-1.medium.com/max/800/1*L4ngkIofRW9O2Z2l68K-kA.png)

  

“Mesmo se houver a mínima chance ainda vale a pena, sinto saudade de casa” — diz Sheila triste.

“Mas se errarmos nunca mais sairemos daqui!” —responde Bobby nervoso.

“Acho que tive uma ideia” — diz Presto

Ele pega a caixa das mãos de Hank e pergunta a Tiamat, “A primeira chave é a chave do portal para casa?”.

![](https://cdn-images-1.medium.com/max/800/1*LF4E_N54a7Gfhtt-5RH2hw.jpeg)

“Não” — Responde Tiamat

Presto joga fora a primeira chave e reagrupa as chaves restantes.

```text
[]string{“chave2”, “chave3”, “chave4”, “chave5”}
```

“A primeira chave é a chave do portal para casa?” — pergunta Presto novamente.

“Não” — Responde Tiamat

Presto joga mais uma chave fora e começa a reagrupar novamente as chaves restantes.

“Entendi, Tiamat disse que responderia apenas se a primeira chave da sequencia é a chave do portal, então ele vai responder a todas as chaves que estiverem na primeira posição!” — diz Diana.

![](https://cdn-images-1.medium.com/max/800/1*f41sm6VhZgnCHJ9kvT6z_g.jpeg)

“Ainda bem que ele não precisou usar o chapéu mágico” — diz Erick agora rindo.

Então Presto repetiu por mais três vezes o processo até que não restasse mais nenhuma chave

```text
[]string{“chave3”, “chave4”, “chave5”}
[]string{“chave4”, “chave5”}
[]string{“chave5”}
[]string{}
```

“Parece que vocês conseguiram descobrir a recursão” — diz Tiamat

“Sim, mas não funcionou porque nenhuma das chaves é a chave do portal” — diz Hank decepcionado.

“Realmente, todas as chaves eram encantadas e todas elas levavam a uma armadilha do Vingador, ele se disfarçou para enganar vocês” — responde Tiamat.

E mais uma vez os jovens não voltam para a sua dimensão, mas pelo menos agora eles sabem recursão.

  

Esse exemplo inclui bem as três definições usadas anteriormente, temos uma repetição, o caso base onde seria encontrar a chave e ai encontrada ou terminado o numero de itens a repetição termina, e a operação que nesse caso poderia ser um slicing no Golang, (chaves[1:]) ou a função (pop) que retira os itens de um array no Javascript.

Ainda faltam mais exemplos, mas aguenta que a gente chega lá.

## Call Stack

  

O que acontece por detrás das cortinas quando funções são chamadas?

Existe uma estrutura de dados que armazena as chamadas de funções, elas não são chamadas aleatoriamente. Assim como nosso código é lido de cima para baixo em uma certa ordem as funções também são organizadas nessa estrutura.

![](https://cdn-images-1.medium.com/max/800/1*nPueCX1vHg-miJ51Q75qJQ.gif)

Podemos ver como a Stack recebe itens com a operação de Push e todos eles entram empilhados como pratos, e quando usamos a operação Pop para remover itens da estrutura, os itens são retirados do último para o primeiro assim como faríamos se fossem pratos, pra que a pilha não quebre.

Vamos implementar ela futuramente então não vamos mergulhar assim tão fundo pois não é o assunto principal.

```go
package main

import  "fmt"

func  main() {
	acorda()
}

func  tomarBanho() string {
	return  "xuaxua"
}

func  tomarCafeDaManha() string {
	refeicao  :=  fazComida()
	return refeicao
}

func  fazComida() string {
	return  "ovo frito e café"
}

func  acorda() {
	tomarBanho()
	tomarCafeDaManha()
	fmt.Println("Ok bora trabalhar")
}
```

Aqui existem alguns exemplos de funções e algumas funções compostas de outras funções. Como seria a pilha das chamadas dessas funções?

- push acorda (Primeira call na stack)  
- push acorda > tomarBanho (Agora temos duas calls na stack)  
- log(“xuaxua”)  
- pop tomarBanho (Depois de executada tomarBanho sai da stack)  
- push acorda > tomarCafeDaManha (Temos duas call na stack de novo )  
- push acorda > tomarCafeDaManha > fazComida (Terceira call)  
- log(“ovo frito e café”)  
- pop fazComida (Terceira call retirada)  
- pop tomarCafeDaManha (Segunda call retirada)  
- log(“Ok bora trabalhar”)  
- pop acorda() (A primeira call é a última a sair da stack)

Essa são calls de funções comuns, agora entender como se comporta a chamada de uma função recursiva parece mais simples.  
Partindo do ponto de que elas funcionam como repetições uma função recursiva em uma stack se parece com isso:

push funcaoRecursiva()  
push funcaoRecursiva()  
push funcaoRecursiva()  
push funcaoRecursiva()  
push funcaoRecursiva()  
push funcaoRecursiva()  
…

E ela continua assim até que o caso base se resolva, caso não haja um caso base podemos ter um problema com estouro de pilha que é quando o numero de chamadas é muito alto e excede a quantidade suportada pela pilha. Também conhecido como stackOverFlow.

  

## Exemplos

  

```go
func  sumRange(num int) int {
	if num == 1 {
	return 1
}
	return num + sumRange(num-1)
}
```

O caso base de sumRange é “se num for igual a 1”, simples. Como temos certeza de que essa recursão não vai ficar rodando infinitamente?

A partir da operação sumRange(num -1), em toda chamada num será subtraído até que em algum momento num será igual a 1 é onde a repetição termina. Temos o caso base, operação e repetição.

Se a chamada for sumRange(4)

Será algo como isso:

```text
4  +  sumRange(3) push
3  +  sumRange(2) push
2  +  sumRange(1) push
1 (caso base) push

na ida
na volta

return  1 pop
return  2  +  1 pop
return  3  +  3 pop
return  4  +  6 pop

valor final de num  =  10
```

Vamos a mais um exemplo.

```go
func  fact(n int) int {
	if n == 1 {
	return 1
}
	return n * fact(n-1)
}
```

Esse é clássico, o exemplo do fatorial. O caso base fica óbvio “quando o input for 0” a repetição termina. Temos duas operações, multiplicação e subtração.

Se o nosso input for 3 fact(3)

Teremos 3 repetições onde a primeira chmada de fact(n) é 3 que é o valor do input: push —  fact(3)  
A próxima é a primeira chamada recursiva onde fact(n-1), isso transforam n em 2: push — fact(2)  
Ultima chamada onde n vale 2 passando pela chamada recursiva fact(n-1) temos 1: push — fact(1)

Agora que n é igual a 1 ele atende ao caso base e terminamos as repetições, a stack começa a esvaziar e fazemos as operações de fora da recursividade.

3 * fact(2) = 6 pop  
6 * fact(1) = 6 pop  
fact(6) pop

Essa ultima a sair da stack ela não tem uma operação de multiplicação pois ela não é uma chamada recursiva ela é a chamada principal.

São exemplos bem básicos mas que servem pra vc ter uma boa ideia do comportamento de uma recursão e como criar sua própria função recursiva.

## Exemplos de Recursão

  

**Recursão Direta:** Exemplo que você acabou de ver. A recursão direta chama a a si mesma sem assistência de outras funções.

**Recursão Indireta**: O tipo de recursão em que a função **A** chama outra função **B** e esta função, por sua vez, chama a função **A**. Este tipo de recursão requer o auxílio de outra função. A função chama a si mesma, mas indiretamente, ou seja, por meio de outra função. O exemplo a seguir explica o conceito de recursão indireta:

```go
package main

import (
"fmt"
)

func  printaUm(n int) {
	if n >=  0 {
	fmt.Println("Na primeira func:", n)
	printaDois(n -  1)
	}
}

func  printaDois(n int) {
	if n >=  0 {
	fmt.Println("Na segunda func:", n)
	printaUm(n -  1)
	}
}

func  main() {
	printaUm(10)
	printaUm(-1)
}
```

No caso essa função printa a contagem regressiva de 10 a 0 alternativamente. O output:

```text
Na primeira func: 10
Na segunda func: 9
Na primeira func: 8
Na segunda func: 7
Na primeira func: 6
Na segunda func: 5
Na primeira func: 4
Na segunda func: 3
Na primeira func: 2
Na segunda func: 1
Na primeira func: 0
```

#### Tail Call Recursion

  

Uma recursão em calda é uma chamada de sub-rotina que é a última chamada da função. Aqui, a chamada recursiva é a última coisa executada pela função.

```go
package main

import (
"fmt"
)

func  printaInt(n int) {
	if n >  0 {
	fmt.Println(n)
	printaInt(n -  1)
	}
}

func  main() {
	printaInt(5)
}
```

Output:

```text
5
4
3
2
1
```

#### Head Recursion

  

Em uma head recursion, a chamada recursiva é a primeira instrução na função. Não há nenhuma outra instrução ou operação antes da chamada. A função não precisa processar nada no momento da chamada e todas as operações são feitas no momento do retorno.

```go
package main

import (
"fmt"
)

func  printaNum(n int) {
	if n >  0 {
	printaNum(n -  1)
	fmt.Println(n)
	}
}

func  main() {
	printaNum(5)
}
```

Output:

```text
1
2
3
4
5
```

O exemplo inverso contando de 1 a N.

Além desses ainda existe a recursão de função anonima e a recursão infinita que é uma recursão que não tem caso base atendido e fica se repetindo até o seu notebook fica assim:  
  

![](https://cdn-images-1.medium.com/max/800/1*ztbqe0v6ZHq-9oH1lMGnIQ.jpeg)

## Porque usar recursão?

  

Recursão é muito utilizada apesar de ser frequentemente ignorada ou abordada superficialmente em cursos introdutórios de programação.  
Quase todo algoritmo usado em estrutura de dados tem uma abordagem recursiva além da iterativa (for, while) e em alguns casos pode ser que a complexidade entre as duas abordagens seja completamente diferente.

É bastante provável que se você for iniciante você ainda não sabe resolver problemas recursivamente. Esse tópico se faz necessário devido as estruturas como arvores binárias que virão, esse é um dos motivos.

A recursão aparece em diversos lugares não apenas em algoritmos de estruturas de dados. A linguagem que você utiliza pode ter funções recursivas na própria stdlib.

No final mesmo que você não use, vocẽ vai acabar lendo então é necessário pelo menos entender para ler e pensar talvez em uma mesma solução iterativa caso prefira.

## Recursivo x Iterativo

  

Tudo que se faz recursivamente pode ser feito iterativamente, então qual o ponto? A questão é que em alguns casos é muito mais fácil resolver os problemas recursivamente, ou a recursividade trará uma melhoria em performance em relação a uma solução iterativa para o mesmo problema.

Como a maior parte dos problemas em computação, depende. Porém saber diferenciar caso a caso qual é a melhor opção é o que faz de você um programador melhor. Você agora tem um segundo modo de resolver problemas repetitivos em programação e isso é ótimo.

Algumas das vantagens que que a recursividade pode trazer ao seu código é a legibilidade, a recursividade diminui a repetitividade no código, fazendo com que ele atenda o principio DRY “dont repeat yourself” encunhado por Andy Hunt e Dave Thomas no livro “The Pragmatic Programmer”.

> “DRY é um princípio de desenvolvimento de software que visa reduzir a repetição de padrões de software substituindo-o por abstrações ou usando a normalização de dados para evitar redundância.” — Wikipedia

E aí entra um outro “depende”, a recursão vista como um meio de se resolver problemas na programação é usada por uma certa parcela de desenvolvedores não são todos os devs que compreendem ou preferem e isso depende mais da sua equipe. Impor isso a uma equipe que está acostumada a resolver as coisas iterativamente só vai atrapalhar a produtividade.

Devs que utilizam o paradigma funcional sempre preferem recursividade ou simplesmente não tem escolha devido a linguagens onde o paradigma é imposto. Recursão é um pattern empregado para contornar o uso de loops. Como os loops sempre mantêm um estado interno para saber em qual rodada eles estão, não podemos usá-los sob o paradigma de programação funcional.

## Mais tarde…

  

![](https://cdn-images-1.medium.com/max/800/1*Fw2GyzAkXHPynSJEPPRw3w.png)

  

![](https://cdn-images-1.medium.com/max/800/1*RcyzQS5mNpc1bXnV7ClwpQ.png)

  

![](https://cdn-images-1.medium.com/max/800/1*d19b7ybmxfp3fPSahJiNKQ.png)

  

![](https://cdn-images-1.medium.com/max/800/1*yQZnuUBPwUAeR6O_03h0hw.png)


Próximo:

[Resolução de Problemas](https://github.com/wagnerdevocelot/DSA/tree/master/PROBLEM_SOLVING)