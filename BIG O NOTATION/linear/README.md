## Complexidade Linear O(n)

A complexidade linear se da em relação ao numero de itens que a sua função recebe, então quanto maior o input maior o tempo de execução do algoritmo.

Em Big O Notation, complexidade linear é apresentada como _O(n)_. Algoritmos de String matching, como Boyer-Moore e Ukkonen tem complexidade linear.

Mas a questão toda não é pegar os algoritmos conhecidos e procurar no Google a complexidade deles e sim enxergar os padrões nesses algoritmos e entender a ponto de flagrar a complexidade no seu código e de outras pessoas. A pergunta a ser feita não é “_esse algoritmo é O(n)?_” e sim “_O porquê desse algoritmo ser O(n)?_”.

Se chegar ao final e conseguir responder essa ultima pergunta, você entendeu Big O Notation.

A complexidade linear,_O(n)_, é demonstrada em um algoritmo da seguinte forma:

```go
package main

import (
	"fmt"
)

func linear(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total++
	}
	return total
}

func main() {
	fmt.Println(linear(10))
}
```

Aqui temos uma operação apenas e não três como anteriormente, nesse caso é **total ++** que é o mesmo que: total é igual a total + 1, porém a complexidade é totalmente diferente pois o **input** também define quantas vezes **i** será incrementado no **for** statement. Então quanto maior for o input_(n)_ maior o numero de operações serão feitas em total.

Se no lugar de _(i <= n)_ eu colocasse _(i <= 10)_ seria um tempo constante _O(1)_ mas como o **input** interfere no for essa complexidade é _O(n)_.

Se formos olhar cada pedaço e dizer o número de operações:

-   _total := 0 (Uma operação)_
-   _i := 0 (Uma operação)_
-   _i++(N operações)_
-   _total ++( N operações)_

Quando for analisar a complexidade do algoritmo não é necessário contar todas as operações, na maioria das vezes você vai acabar percebendo. Nesse caso como temos operações que são relacionadas a **N** que é o numero de inputs não tem porque nos preocuparmos com as que são constantes pois o Big O é sobre o pior caso, então se a gente deixar as constantes de lado sobram apenas as _O(n)_.

E se você tiver dois for statement? Tipo assim:

```go
package main

import "fmt"

func elevador(n) {
    for i := 0 ; i < n; i++{
        fmt.Println(i)
    }

    for j := n - 1 ; j >= 0; j--{
        fmt.Println(j)
    }
}
```

Ainda assim a complexidade será _O(n)_ e não _O(2n)_ porque Big O é sobre a perspectiva geral.

Pode ser que eu esteja batendo na mesma tecla, porém recomendo que leia mesmo as partes que você considera que já sabe. O maior perigo de aprender é achar que já entendeu e não precisa estudar mais.

![](https://cdn-images-1.medium.com/max/800/1*4bXJwVpMhbdiff-FshcGBQ.png)

Agora temos em perspectiva dois tipos de complexidade a constante em **cinza** e a linear em **azul** e agora fica muito óbvio o porque do nome. Repare que dessa vez o **input** é de _10gb_ e não _10tb_ se eu tentasse fazer com um input de _10tb_ em um algoritmo de complexidade _O(n)_ levaria muito tempo. A linha de _O(1)_ permanece plana não leva nem 1 segundo de runtime, a azul com um input de _10gb_ já leva 18 segundos.

Eu não to falando pra você ficar contando segundo, isso vai depender do seu cenário, mas se você olhar o quanto as linhas já se distanciaram é bastante e como diz o ditado “de grão em grão a galinha enche o papo”.

Próximo:

![O(n²) - O(n³) Exponencial](https://github.com/wagnerdevocelot/DSA/tree/master/BIG%20O%20NOTATION/exponencial)