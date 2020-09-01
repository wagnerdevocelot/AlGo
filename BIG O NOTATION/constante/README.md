## Constantes O(1)

Fazemos isso contando as operações, exemplo:

```go
package main

func main() {

}

func constante(n int) int {
	return n * (n + 1) / 2
}
```


Aqui temos, uma multiplicação, uma adição e uma divisão. três operações.

E ai não importa se **N** é 2 ou 1 bilhão, o numero de operações é o mesmo, três! Essa é uma complexidade de _O(3)_ é um tipo de complexidade constante pois o numero de operações não muda mesmo que o input seja diferente.  
Só que, em Big O a notação tem regras onde você não precisa ficar somando cada operação, _O(3)_ ou _O(200)_ no final tempo constante é sempre _O(1)_.  
Geralmente se ignora as constantes em um algoritmo pois o tempo de execução em _O(1)_ em relação as outras complexidades é totalmente irrelevante.

![](https://cdn-images-1.medium.com/max/800/1*TVLSKfjYYMpMLCdps-z2GA.png)

Nesse gráfico fica bem explicito o como uma complexidade constante se comporta, no eixo **vertical** temos a relação de tempo e **horizontal** a de **N** (input) temos um ultimo input de **10tb** sobre esse algoritmo do exemplo e nessa simulação ele não leva nem um milésimo de execução. Mais a frente olharemos o mesmo gráfico em perspectiva a diferentes complexidades.

Próximo:

![O(n) Linear](https://github.com/wagnerdevocelot/DSA/tree/master/BIG%20O%20NOTATION/linear)