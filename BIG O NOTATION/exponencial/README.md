## Complexidade Quadrática O(N²)

  
```go
package main

import (
	"fmt"
)

func printaTodosOsPares(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println(i, j)
		}
	}
}
``` 

Acabamos de falar sobre um caso onde temos dois **for** e a complexidade era _O(n)_ mas esse caso é diferente pois o **for** é aninhado e não é como se fosse um caso de _N * 2_ é **N ao quadrado**, se **N** fosse 10 o retorno seria 100 a proporção em que a complexidade cresce é muito maior.

Enquanto a complexidade linear sobe em uma linha linear “_duh”_, a complexidade quadrática sobre em uma curva exponencial em relação ao eixo de tempo, para cada vez que **N** aumenta.

Nesse caso a regra de pensar no pior caso também é válida, se vc tiver dois for aninhado e um solto dentro da função vc não tem um _O(n + n²)_ porque o **N** linear é insignificante perto do quanto **N²** cresce. Então no final conte apenas com o Worst case, _O(n²)_.

Output:

``` text
00
01
10
11
```

Com um parâmetro de 2 esse seria o output, pois os dois loops percorrem os índices **i** e **j** até que atinjam o parâmetro **N**, não temos mais uma linear são dois índices em paralelo.

Existem também a **Complexidade Cúbica O(n³)** seriam três loops aninhados a lógica é a mesma, não pretendo abordar sobre ele por aqui mas aqui vai um exemplo:

``` go
package main

import (
	"fmt"
)

func main() {
	var k, l, m int

	var arr [10][10][10]int

	for k = 0; k < 10; k++ {
		for l = 0; l < 10; l++ {
			for m = 0; m < 10; m++ {
				arr[k][l][m] = 1
				fmt.Println("Valor do elemento", k, l, m, " é", arr[k][l][m])
			}
		}
	}
}
```

Temos três loops percorrendo um array multidimensional a complexidade aumenta tanto que pra percorrer esse array com três loops são precisas mil operações pois 10 x 10 x 10 = 1000.

![](https://cdn-images-1.medium.com/max/800/1*14SAgGBfb7QhSF9NNsQ4zg.png)

Agora dá pra ter uma ideia da perspectiva, veja como a linha **vermelha** de uma complexidade **quadrática** sobre em relação as outras e também como a complexidade **linear** parece plana em relação a **quadrática**. De repente aqueles 18 segundos não são nada comparados a mais de 1 minuto isso com um input de _10kb_. Então se eu tivesse no mesmo algoritmo essas três complexidades pouco importa se o _O(n)_ vai receber _1tb_ de **input** o que ta te ferrando é essa _O(n²)_, por isso nos preocupamos apenas com o pior caso em relação a Big O Notation.

**Algumas ideias pra você calcular as operações com Big O**

-   Operações aritméticas são sempre **constant time**
-   Atribuições de variáveis são **constant time**
-   Acessar elementos de array por index ou objetos por chave **constant time**
-   Em um loop a complexidade é o tamanho do loop * a complexidade do que tiver dentro do loop

Não precisa decorar mas é bom saber.

[O(log n Logaritmo)](https://github.com/wagnerdevocelot/DSA/tree/master/BIG%20O%20NOTATION/logaritmo)
