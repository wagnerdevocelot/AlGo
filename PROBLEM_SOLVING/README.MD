# Resolução de Problemas
Dicas e padões para problem solving

## Como funciona um algoritmo

A questão é que para entender como resolver algoritmos temos que primeiro entender o que são.
E a definição mais simples que se encontra em qualquer lugar é que Algoritmos são compostos de passos lógicos para resolver um problema.

Apesar de ser simples esquecemos disso as vezes! Por exemplo quando deixamos de pensar realmente nos passos que o algoritmo excuta deixamos de entender a sua propria lógica.

Em uma busca linear nós simplesmente pensamos, ok vamos fazer um for e procurar o parametro.
Se pensarmos na lógica é como se indice fosse um carteiro e ele tem que entregar uma carta na casa do pedro e as casas estão dentro do array. 
E ai o carteiro vai passando de casa em casa

0- "É aqui que mora o Pedro?"
"não"
1- "É aqui que mora o Pedro?"
"não"
2- "É aqui que mora o Pedro?"
"não"
3- "É aqui que mora o Pedro?"
"não"
4- "É aqui que mora o Pedro?"
"sim"
"Essa carta é pra você"

Bem engraçado, mas é isso que um loop ta fazendo e nós já sabemos disso. A questão é que em certo momento a gente deixa de se preocupar, porque é tão bobo, só que quando aparece um desafio e a gente ta preocupado com alto nível a nossa cabeça meio que da um tilt.

Isso é muito comum, e se você tem dificuldade para resolver problemas com algoritmos possivelmente é porque você nunca foi cobrado por isso, não era uma  questão entre conseguir ou não uma vaga. Também porque a maioria das empresas fazem seus testes baseados na premissa de que vão contratar pessoas já formadas então é pressuposto que esse tipo de conhecimento ja tenha sido aprendido.

E a diferença entre você e uma pessoa que resolve problemas com facilidade não é que ela tem um cerebro superior, é só que ela se propos a resolver mais problemas que você. Então a questão toda é prática. 
Da mesma forma a gente pode comparar alguem que já deu umas 10 palestras em meetups e outra que vai palestrar pela primeira vez, a primeira não é um ser superior inalcançavel, para de ficar endeusando qualquer um a troco de nada.

## Como Solucionar um problema que eu não conheço?

O primeiro problema sempre vai ser muito dificil, o segundo ainda mais. Só que existem algumas estratégias que ajudam a gente a resolver esses problema e podemos fazer uso delas para que essa trajetória não seja tão desgastante.

Então baseado nisso podemos:

- Criar um plano para resolução de problemas
- Dominar padrões para solução de problemas simples

Quando já temos uma ideia de abordagem para resolver problemas as coias se tornam um pouquinho mais simples. Não se trata de algo milagroso que vai resolver o problema pra você, a tarefa ainda é nossa, mas pelo menos ficamos mais relaxados e conseguimos lidar melhor com a ansiedade durantes um teste.

Primeiro:

- Entenda o problema
- Decompor
- Resolva
- Refatore

## Entenda o problema

Parece bobo, mas a maioria dos problemas que resolvemos são problemas que já conhecemos e se já conhecemos já sabemos como resolver.
Se alguém nos da um problema que nunca vimos e não fazemos nem ideia de por onde começar, o que faríamos?

Então nesse momento o que precisamos fazer é dar um passo para trás, soltar o teclado e pensar no problema. 
A maior parte do stress no trabalho vem das pessoas que se colocam no meio de vários problemas e ela fica tão desnorteada que  acaba ficando irritada e até agressiva com toda a frustração.

O melhor nesses casos é sair de onde está o problema, o ambiente é consimido pelo problema pois as pessoas se deixam consumir pelo problema. Então nos momentos de stress o ideal é nos levantarmos, beber um café uma água e pensar no problema fora do ambiente onde ele está sendo discutido.

Agora dando minha opnião, esse é um dos grandes motivos para reuniões serem desastres quando temos problemas sérios. Quando tudo está funcionando a reunião é completamente amigável todos contribuem com coisas que não precisamos no momento, e quando a coisa aperta e as reuniões tem uma atmosfera ruim dificilmente alguém vira com uma ideia que partiu dessa reunião.

Isso serve um pouco pra algoritmos mas também pra vida pessoal, não deixemos nos levar pelo clima que é gerado no ambiente se tem 10 pessoas usando o metodo Y pra resolver X e não ta funcionado porque 11 + Y resolveria X?

Pra saber se conseguimos entender o problema algumas perguntas podem ser feitas:

- Eu consigo explicar o problema com as minhas proprias palavras?
- Quais os inputs são feitos ao problema?
- Qual o output é esperado? Como ele se parece? Qual o seu tipo?
- O output é determinado pelo input?
- Como eu posso nomear os dados que fazem parte do problema?


## Decompor

Se um algoritmo é uma sequencia de passos lógicos finitos, uma boa ideia para resolver um problma que não conhecemos é decompor os passos lógicos como se já tivessemos o algoritimo pronto e fossemos tirando pedaço a pedaço e analisando um por um.

Geralmente em uma entrevista quando o entrevistador passa um challenge ele quer que expliquemos as suas ações para entender a nossa linha de raciocinio como programador, então é bom manter um nivel de comunicação enquanto se resolve o problema, explicando os passos lógicos que pretendemos tomar.

Isso nos ajuda a pensar no código antes mesmo de escrevê-lo e durante essa decomposição podemos acabar diferenciando quais são os passos que, "já sei como vou fazer isso" dos "não sei nem por onde começar".

## Resolva 

Agora que já temos um plano e uma linha de raciocinio já vale a pena começar a escrever código. Não dá pra sair escrevendo qualquer coisa e achar que vai resolver na sorte, isso já não funciona nem no dia a dia imagine em uma entrevista onde o tempo é curto.

#### Refatore

Com o problema resolvido já da pra fazer algumas perguntas, a complexidade ta lega? Eu escolhi bons nomes para as variáveis? Tá legível o negócio? Caso não esteja de uma revisada pra deixar a solução bacana.

# ABORDAGENS


## Brute Force

Um algoritmo de força bruta resolve um problema com base na sua própria declaração e definição como ordenação e busca. Um algoritmo de força bruta pode ser uma pesquisa exaustiva onde a solução está na propridade de um conjunto. Talvez você lembre do termo de brute force baseado em "brute force attack" que tem a ver com a quebra de senhas através da abordagem de tentativa e erro. Apesar das abordagens serem parecidas a usabilidade é diferente.

Ela acaba sendo exaustiva pois o numero de processamentos feitos para ser realizado é muito grande. Como o exemplo do carteiro. 

Algoritmos de força bruta são conhecidos por sua ampla aplicabilidade e simplicidade na resolução de problemas complexos. Busca, string matching e multiplicação de matrizes são alguns cenários onde eles são usados.
Apesar de resolverem diversos problemas e serem de fácil implementação podem ser muito pouco performáticos.

A representação de um algoritmo de força bruta é mostrada no seguinte código:

``` go 

package main

import (
	"fmt"
)

func encontraElemento(arr [10]int, elemento int) bool {
	for i := range arr {
		if arr[i] == elemento {
			return true
		}
	}
	return false
}

func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var tentativa bool = encontraElemento(arr, 10)
	fmt.Println(tentativa)
	var tentativa2 bool = encontraElemento(arr, 9)
	fmt.Println(tentativa2)
}
``` 

Essa é uma busca linear e como no exemplo do carteiro ela passa por todos os elementos do array por isso é considerada brute force pois ela exerce o número máximo de operações até chegar no objetivo o que tem bastante a ver com a questão de Big O que abordei anteriormente sobre o pior caso.

Aqui tem um vídeo do programação dinâmica falando um pouco sobre com o exemplo do handshake

https://youtu.be/tnb2V7h4Fa4

#### Divide and conquer algorithms

Um algoritmo de divisão e conquista divide um problema complexo em problemas menores e
resolve esses problemas menores. O problema menor será subdividido até que seja resolvido.

Esta abordagem pode ser dividida nas seguintes partes:

- Divide: Envolve dividir o problema em algum subproblema.
- Conquiste: Subproblema chamando recursivamente até que o subproblema seja resolvido.
- Combine: O Sub problema foi resolvido para que possamos encontrar a solução do problema.

Se lembrarmos do exemplo de busca binária do tweet anterior sobre Big O Notation a ideia aqui é a mesma dividir e conquistar, utilizando a recursão.

Recursão, quick sort, binary search e merge sort são bons exemplos de algoritmos de dividir e conquistar.
A memória é usada de forma eficiente com estes algoritmos.

``` go

package main

import (
	"fmt"
)

```

Conforme mostrado no código a seguir, o método Fibonacci leva o parâmetro inteiro numero e
retorna o Fibonacci para numero usando recursão.

Caso não lembre a formula da sequencia fibonacci:   Fn = Fn-1 + Fn-2
O Fibonacci de 9 é 55 pois 144 e 89 são os próximos números na sequencia.

Usando um loop aqui gerando uma sequência de 10 números fibonacci.


```go

func fibonacci(numero int) int {
	if numero <= 1 {
		return 1
	}
	return fibonacci(numero-1) + fibonacci(numero-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}

```

### Backtracking algorithms

Backtracking é uma técnica algorítmica para resolver problemas recursivamente, tentando construir uma solução de forma incremental, uma peça de cada vez, removendo as soluções que falham em satisfazer as restrições do problema.

A seguir um exemplo de um algoritmo de backtracking. Onde problema é identificar combinações de elementos em uma matriz de 10 elementos cuja soma é igual a 18.
O método combinacaoDeSoma tenta recursivamente encontrar a combinação. Sempre que
soma vai além da meta chave, ele volta atrás:


```go

package main

import (
	"fmt"
)

func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var chave int = 18
	var combinacoes [19]int
	combinacaoDeSoma(arr, combinacoes, len(arr), chave, 0, 0, 0)
}

func combinacaoDeSoma(arr [10]int, combinacoes [19]int, tamanho int, chave int, addValor int, l int, m int) int {
	numero := 0

	if addValor > chave {
		return -1
	}

	if addValor == chave {
		numero = numero + 1
		for p := 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinacoes[p]])
		}
		fmt.Println(" ")
	}

	for i := l; i < tamanho; i++ {
		combinacoes[m] = l
		combinacaoDeSoma(arr, combinacoes, tamanho, chave, addValor+arr[i], l, m+1)
		l = l + 1
	}
	return numero
}

```

Ainda falaremos mais sobre backtracing.


# Quais habilidades o entrevistador procura?

Se pensarmos pelo ponto de vista do entrevistador em um code challenge quais seriam as habilidades que eles prestariam atenção?

Habilidade de Análise - Não querem que apenas resolvamos o problema e vamos embora, querem saber como pensamos no problema. Então é importante dizer os passos lógicos enquanto pensa na solução assim eles podem entender como é o nosso raciocinio.

Código - Essa já é uma questão mais próxima do que estamos acostumados, nosso código é bem escrito? Limpo e organizado ou legivel?

comunicação - Durante a analise como nos comunicamos? fazemos perguntas pertinentes? O perfil bate com a cultura da empresa?


O passo a passo através de um problema.


1. Quando o entrevistador explicar o desafio, anote os pontos chave para que você consiga interpretar  a sua forma e assim compreender melhor o problema. Mostre como você organiza seu pensamento.

2. Tenha total compreensão de quais são os inputs e outputs, se não tiver certeza pergunte.

3. Qual é o valor mais importante do problema? Você tem tempo, espaço e memória, etc. Qual é o objetivo principal?

4. Não faça perguntas demais, pode ser um pouco irritante.

5. Comece com a abordagem ingênua / força bruta. Geralmente a primeira solução que vem na cabeça, e fale sobre ela, não precisa necessáriamente escrever a implementação. 

6. Diga por que essa abordagem não é a melhor (ou seja, O(n^ 2) ou superior, não legível, etc ...)

7. Percorra sua abordagem, comente coisas e veja onde você pode quebrar coisas. Qualquer repetição, gargalos como O (N^ 2) ou trabalho desnecessário? Você usou todas as informações disponíveis? Gargalo é a parte do código com o maior Big O.

8. Antes de começar a codificar, analise seu código e anote as etapas que você pretende fazer.

9. Modularize seu código desde o início. Divida seu código em pedaços pequenos e adicione apenas comentários, se necessário.

10. Comece realmente a escrever seu código agora. Tenha em mente que quanto mais você se prepara e entende o que você precisa codificar, melhor será o whiteboard. Portanto, nunca comece a codar sem ter certeza de como as coisas vão funcionar. Essa é uma receita para o desastre.

11. Pense nas verificações de erro e como você pode quebrar esse código. Nunca faça suposições sobre o input. Suponha que as pessoas estão tentando quebrar seu código. Como você o protegerá? Dica: Comente no código, as verificações que deseja fazer ... escreva a função e diga ao entrevistador que você escreveria testes agora para fazer sua função falhar (mas você não precisará realmente escrever os testes caso tenha pouco tempo).

12. Não use nomes ruins/confusos como i e j. Escreva um código legivél.

13. Teste seu código: verifique se há parâmetros, 0, indefinido, nulo, matrizes massivas, código assíncrono, etc ... Pergunte ao entrevistador se pudermos fazer suposições sobre o código. Você pode fazer a resposta retornar um erro?

14. Se o seu entrevistador estiver satisfeito com a solução, a entrevista geralmente termina aqui.


Verificação de código:
[✅] Funciona
[✅] Bom uso de estruturas de dados
[✅] Reutilizar código/não se repetir
[✅] Modular - torna o código mais legível, sustentável e testável
[✅] Menor que O (N^ 2). Queremos evitar loops aninhados, se possível, pois eles são caros.
[✅] Baixa complexidade de espaço -> Recursão pode causar stackoverflow, a cópia de matrizes grandes pode
exceder a memória da máquina.

Essas dicas são mais valiosas para entrevistas em empresas como o Google, Amazon, Facebook e etc, pois é mais comum que eles utilizem estrutura de dados e algoritmos como foco da entrevista, no restante das empresas a solução varia de acordo com o segmento. Mas se temos dominio nesse nível é muito provável que consigamos lidar com testes de outras empresas.

Essas são dicas interessantes e que são possíveis de praticar em diversas plataformas de code Challenge como o HackerRank por exemplo. A ideia toda aqui é nos fazer passar por problemas de algoritmos que não sabemos a resposta mas que temos as ferramentas necessárias para ler o enunciado, interpretar e pensar na nossa própria solução antes de procurar uma explicação no google. 

Se olharmos pra um desafio, sentirmos que não sabeamos como resolver mas não despendemos de um tempo pra pelo menos pensar em uma solução e procuramos  tutoriais no google pra que expliquem como chegaram naquela resposta é meio que terceirizar nosso cerebro pro google e da forma que eu to falando parece até que eu sou o mestre dos code challenges, mas não eu me sinto tão desconfortável pensando em uma resolução quanto você. 

Você não é uma fraude por não conseguir resolver X problema, isso tudo é um processo pra mim e pra você e logo mais vamos estar ajudando outras pessoas a resolverem esses problemas.
