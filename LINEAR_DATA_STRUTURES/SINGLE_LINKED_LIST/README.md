### LINKED LISTS

Single

![](https://cdn-images-1.medium.com/max/800/1*pC4OofsOSxQj3K_3cDM3dg.png)

Você pode chamar de _lista ligada, lista encadeada_, eu vou chamar de _linked list_.

A grande diferença _linked list_ e _arrays_ é que _linked list_ não são indexadas, não tem como buscar o sexto elemento, ou o segundo e etc. A linked list funciona como estações do monotrilho, onde cada estação é um valor e existe uma conexão entre uma estação e outra, formando uma linha.

E por isso podemos declarar a _Linked list_ como uma estrutura de dados linear.

O exemplo com o monotrilho é bem específico exatamente porque vou falar de _doubly linked list_ também e essa tem mais a ver com linhas de metrô.

N_a linked list_ cada **Node** ou nó irá representar uma estação, cada node tem uma propriedade que pode ser um _int_, _string_ ou qualquer outro tipo. Nesse caso será uma **string** e a propriedade são as **estações** do **monotrilho**.

A segunda parte de um **Node** é um **ponteiro** que faz a ligação apontando para o próximo node, no nosso exemplo, a próxima **estação**, criando assim uma ligação entre todos os nós.

![](https://cdn-images-1.medium.com/max/800/1*UDg8GhUQD8X-_2DxVK-KAw.png)

No nosso exemplo enquanto os **Nodes** são as **estações**, o **Monotrilho** vai representar a própria linked list.

Outras duas partes importante nas linked list são o headNode e o tail o head é o primeiro item que entra na lista e o tail o último, no caso do tail o **pointer** dele aponta para **nil**, o que representa o fim da lista. Essas duas partes também são importantes pois são referência para a criação dos métodos que precisaremos usar na **linked list**.

![](https://cdn-images-1.medium.com/max/800/1*V89SwrMrm2agrgn1Ymk0og.png)

Claro que uma estrutura da de dados não é nada seus seus métodos, e vamos implementá-los aqui. Para entender por completo é necessário que você saiba como usar ponteiros, então vou deixar alguns recursos aqui pra você consultar antes de começar a falar de código.

[Go by Example: Pointers](https://gobyexample.com/pointers)  
[Aprenda Go: O que são ponteiros?](https://www.youtube.com/watch?v=l2YJ-5GpGr8)

#### A estrutura

  

O struct Estação é o que vai representar os nós dessa Linked List, então temos a propriedade que é o valor _nome_ e _proximaEstacao_ que faz a ligação entre os nós apontando para o próximo _Node_.

Com a struct **Monotrilho** temos agora como setar a **primeiraEstacao** (_head_) que é onde começa a linked list, então **primeiraEstacao** vai apontar para **Estacao** (_Node_) que é onde fica a **proximaEstacao**.

![](https://cdn-images-1.medium.com/max/800/1*YYSdowi88jCwKCDAh2O_eQ.png)

#### Adicionar um head

  

O método **AddPrimeiraEstacao** vai adicionar um valor ao primeiro **Node** que é **primeiraEstacao**.

**AddPrimeiraEstacao** usa a struct **Monotrilho** como receiver e possui um parâmetro **string**. Instanciamos **Estacao** e atribuímos o parâmetro a propriedade **nome**.

Agora _linha.primeiraEstacao_ recebe a referencia de memória presente na **Estacao** instanciada.

![](https://cdn-images-1.medium.com/max/800/1*_GnsMeyv5zODTo7e1tyXhQ.png)

#### IterateList

  

Basicamente o que faríamos com um for statement em um array porém usando as propriedades das structs para percorrer a lista de **Estações**.  
Enquanto a Estacao instanciada for diferente de **nil** (representação do fim da lista) **estacao** é igual a **proximaEstacao** análogo ao i++.

![](https://cdn-images-1.medium.com/max/800/1*hkWR9tGSmQ2IsRsXaw8Khg.png)

#### LastNode

  

O método **UltimaEstacao** do **Monotrilho** retorna o node no final da linha. A linha é percorrida para verificar se **proximaEstacao** é **nil** caso seja a variável **ultimaEstacao** recebe o valor atual de **estacao** feito no for.

![](https://cdn-images-1.medium.com/max/800/1*SrS_ZToG2wJdurCNuFd0Ag.png)

#### AddToEnd

  

O método **AdicionaEstacao** adiciona um node no final da linha.

Primeiro buildamos um Node, **nome** = _parâmetro_ e o **proximaEstacao** que recebe _nil_ pois como essa estação é adicionada no final da **linha** ela precisa ser um **tail.**

Em seguida encontramos o fim da linha atual reutilizando **UltimaEstacao**() o método criado anteriormente e setamos seu valor com o a **estacao** criada.

![](https://cdn-images-1.medium.com/max/800/1*51btxEhiHDqi1RVzOHtQlQ.png)

#### NodeWithValue

  

O método **ProcuraEstacao** do **Monotrilho** retorna o node com o valor do **parâmetro**. A lista é percorrida e verificada para ver se o valor da propriedade é igual ao parâmetro caso contrário retorna **nil**.

![](https://cdn-images-1.medium.com/max/800/1*PcW6E39we4kypRcb8Z9rpA.png)

#### AddAfter

  

O método **AddProximaEstacao** adiciona um Node após uma **Estacao** específica.

O método **AddProximaEstacao** do **Monotrilho** possui dois parâmetros, onde o primeiro é o node de referencia e o segundo o novo node a ser adicionado na posição seguinte a referência.

Uma estacao com o valor **nomeDaEstacaoAnterior** é recuperada reutilizando o método **ProcuraEstacao**().

Um node (parâmetro 2) é criado e adicionada após o retorno de ProcuraEstacao().

![](https://cdn-images-1.medium.com/max/800/1*QDGjOopnL07eONhwm62WzA.png)

#### RemoveNode

  

O método **RemoveEstacao** verifica se o **primeiraEstacao** é nil, caso o **Monotrilho** esteja vazio.

Em seguida se a propriedade do **primeiraEstacao** for a mesma do parâmetro então **proximaEstacao** é movido para **primeiraEstacao** assim ocupando a posição do **Node** removido.

Em seguida uma nova condição onde **estacaoAtual** é setada como **head** para que possamos verificar os itens além da **primeiraEstacao** e assim caso seja encontrado fazemos a atribuição para a esquerda novamente alocando os nodes seguintes a posição do node removido.

![](https://cdn-images-1.medium.com/max/800/1*Krs67Ns0IqHJspHjU6oGIQ.png)

#### SizeList

  

**TamanhoDaLinha** faz o mesmo papel que **len**() em estruturas de array retornando a quantidade de **Nodes** presentes no **Monotrilho**.

![](https://cdn-images-1.medium.com/max/800/1*zN1OSseP8r_6TjmQJ4D1vQ.png)

#### FuncMain

  

Depois disso tudo você pode utilizar os métodos para construir a sua LinkedList como preferir.

![](https://cdn-images-1.medium.com/max/800/1*BX3JGQ868P47nA1yyWGt8A.png)

Output

![](https://cdn-images-1.medium.com/max/800/1*naXrVJgxxdgsmzI11ZvXiQ.png)

<<<<<<< HEAD
A versão em código no repositório não utiliza a os mesmos exemplos com o Monotrilho, os nomes no arquivo original contam com os nomes comuns usados na estrutura de LinkeList.
=======
A versão em código no repositório não utiliza a os mesmos exemplos com o Monotrilho, os nomes no arquivo original contam com os nomes comuns usados na estrutura de LinkeList.

#### LinkedList vs Arrays

- Um array é a estrutura de dados que contém uma coleção de elementos de dados de tipo semelhante, enquanto a linked list é considerada como uma estrutura de dados não primitiva contém uma coleção de elementos vinculados não ordenados conhecidos como nodes.

- No array, os elementos pertencem a índices, ou seja, se você deseja o quarto elemento, deve escrever o nome da variável com seu índice ou localização dentro do colchete.

- Em uma linked list, porém, você tem que começar do head e ir trabalhando até chegar ao quarto elemento. O acesso a um elemento em um array é rápido, enquanto a linked list leva um tempo linear, portanto, é um pouco mais lento.

- Operações como insert e delete em arrays consomem muito tempo. Por outro lado, o desempenho dessas operações em linked lists é rápido.

- Arrays são de tamanho fixo. Em contraste, as linked lists são dinâmicas e flexíveis e podem expandir e diminuir seu tamanho.

- Em um array, a memória é atribuída durante o tempo de compilação, enquanto uma linked list é alocada em tempo de execução.

- Os elementos são armazenados consecutivamente na memória com arrays, enquanto que são armazenados aleatoriamente em linked lists.


Tudo em relação as estruturas terão tradeoff então a aplicação de cada uma vai depender da necessidade do desenvolvedor.
>>>>>>> f57565e22afe458d64d5bf09230f93569ae287d8
