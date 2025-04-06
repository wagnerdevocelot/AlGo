# AlGo

Este repositório contém implementações e explicações de algoritmos e estruturas de dados em Go (Golang). Ele foi projetado para ajudar desenvolvedores a entender conceitos fundamentais de algoritmos, estruturas de dados e análise de complexidade, com exemplos práticos e bem documentados.

## Estrutura do Projeto

O projeto está organizado em pastas, cada uma representando um tópico ou conceito específico. Abaixo está uma visão geral das pastas e seus conteúdos:

### 1. **RECURSION**
- **Descrição:** Exemplos de recursão direta, indireta, tail recursion e head recursion.
- **Arquivos principais:**
  - `sumrange.go`: Soma recursiva de números.
  - `tailcall.go`: Exemplo de recursão em cauda.
  - `fatorial.go`: Implementação do cálculo de fatorial.
  - `indirect.go`: Exemplo de recursão indireta.
  - `head.go`: Exemplo de recursão em cabeça.
  - `callstack.go`: Demonstração do funcionamento da pilha de chamadas.
- **Documentação:** [README](RECURSION/README.MD)

---

### 2. **PROBLEM_SOLVING**
- **Descrição:** Estratégias e padrões para resolução de problemas.
- **Arquivos principais:**
  - `bruteforce.go`: Busca linear usando força bruta.
  - `dividenconquer.go`: Exemplo de divisão e conquista com Fibonacci.
  - `backtracing.go`: Algoritmo de backtracking para combinações de soma.
- **Documentação:** [README](PROBLEM_SOLVING/README.MD)

---

### 3. **LINEAR_DATA_STRUCTURES**
- **Descrição:** Implementações de estruturas de dados lineares.
- **Subpastas:**
  - **STACK:** Implementação de pilhas (LIFO).
    - Arquivos: `stack.go`, `sliceStack.go`
    - Documentação: [README](LINEAR_DATA_STRUCTURES/STACK/README.md)
  - **QUEUE:** Implementação de filas (FIFO) e filas circulares.
    - Arquivos: `queue.go`, `linkedQueue.go`, `circularQueue.go`
    - Documentação: [README](LINEAR_DATA_STRUCTURES/QUEUES/README.md)
  - **SET:** Implementação de conjuntos com operações de união e interseção.
    - Arquivo: `sets.go`
    - Documentação: [README](LINEAR_DATA_STRUCTURES/SET/README.md)
  - **SINGLE_LINKED_LIST:** Implementação de listas encadeadas simples.
    - Arquivo: `linkedlist.go`
    - Documentação: [README](LINEAR_DATA_STRUCTURES/SINGLE_LINKED_LIST/README.md)
  - **DOUBLY_LINKED_LIST:** Implementação de listas duplamente encadeadas.
    - Arquivo: `doubly.go`
    - Documentação: [README](LINEAR_DATA_STRUCTURES/DOUBLY_LINKED_LIST/README.md)

---

### 4. **BIG O NOTATION**
- **Descrição:** Explicações e exemplos de análise de complexidade de algoritmos.
- **Subpastas:**
  - **constante:** Exemplos de complexidade constante `O(1)`.
    - Arquivo: `constant.go`
    - Documentação: [README](BIG_O_NOTATION/constante/README.md)
  - **linear:** Exemplos de complexidade linear `O(n)`.
    - Arquivo: `linear.go`
    - Documentação: [README](BIG_O_NOTATION/linear/README.md)
  - **polinomial:** Exemplos de complexidade quadrática `O(n²)` e cúbica `O(n³)`.
    - Arquivos: `quadratico.go`, `cubico.go`
    - Documentação: [README](BIG_O_NOTATION/polinomial/README.md)
  - **logaritmo:** Exemplos de complexidade logarítmica `O(log n)`.
    - Arquivo: `logaritmo.go`
    - Documentação: [README](BIG_O_NOTATION/logaritmo/README.md)
- **Documentação geral:** [README](BIG_O_NOTATION/README.md)

---

## Como Executar

1. Certifique-se de ter o Go instalado em sua máquina. [Instale o Go](https://golang.org/doc/install) se necessário.
2. Clone este repositório:
   ```bash
   git clone https://github.com/seu-usuario/AlGo.git
   cd AlGo
   ```
3. Navegue até a pasta desejada e execute os arquivos `.go`:
   ```bash
   go run nome_do_arquivo.go
   ```

---

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests com melhorias, correções ou novos exemplos.
## 