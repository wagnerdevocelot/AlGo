# Sets

Um _Set_ é uma estrutura de dados linear que possui uma coleção de valores que não se repetem. Um _Set_ pode
armazenar valores únicos sem qualquer ordem particular. 

No mundo real, os _Sets_ podem ser usados para coletar todas as tags para postagens de blog ou participantes em uma conversa de chat.

Não tem muito segredo, como queremos uma estrutura que não permita dados duplicados vamos basear a estrutura do _set_ em uma outra estrutura do _golang_ chamada _map_ que é baseada em _hash_, um conjunto de chave valor.

Então definimos o segundo tipo de dados como *boleano*, assim sendo possível fazer comparações de _true_ e _false_ entre os valores.

```go
type Set struct {
	estrutura map[int]bool
}

func (set *Set) Criar() {
	set.estrutura = make(map[int]bool)
}
```

# The insert method

O método _AddValor_ adiciona o valor a um _Set_. 

Aproveitados a o método _PossuiValor_ para verificar se existe o valor dentro do set.
Se o resultado do método não for verdadeiro então _set.estrutura_ recebe valor e definimos o segundo parâmetro do _map_ como _true_.

```go
func (set *Set) AddValor(valor int) {

	if !set.PossuiValor(valor) {
		set.estrutura[valor] = true
	}
}
```

# The delete method

Aqui é reaproveitado o _delete_ do próprio map da _sdtlib_ do _Golang_ então a implementação é mínima.

```go
func (set *Set) DeletaValor(valor int) {

	delete(set.estrutura, valor)
}
```

# The lookup method

O método PossuiValor do _Set_ verifica se o valor existe ou não em _set.estrutura_ e retorna um boolean com a resposta.

Definimos _localizado_ como um boolean, o zaro value dele é _false_

Esse underline cheguei a comentar antes, mas ele anula o valor de uma variável, no caso aqui ele é anula o valor _inteiro_ de _set.estrutura_ para que a comparação possa ser feita apenas com o os _boleanos_.


```go
func (set *Set) PossuiValor(valor int) bool {

	var localizado bool

	_, localizado = set.estrutura[valor]

	return localizado
}
```

# The InterSect method

No código a seguir, o método _CruzarCom_ retorna um _setCruzado_ que consiste na interseção de _set_ e _setDiferente_. 

Instanciamos um _set_ que será o retorno da **interseção**, em seguida fazemos um loop ao longo da estrutura já existente verificando apenas os valores, e anulando o segundo parâmetro do _map_ que seria o **boleano**.

Reaproveitamos o método _PossuiValor_ no parâmetro para fazer a comparação e com o valor da condição reutilizamos _AddValor_ para inserir o valor no _set_ que instanciamos.

```go
func (set *Set) CruzarCom(setDiferente *Set) *Set {

	var setCruzado = &Set{}
	setCruzado.Criar()
	var valor int

	for valor, _ = range set.estrutura {

		if setDiferente.PossuiValor(valor) {

			setCruzado.AddValor(valor)
		}
	}

	return setCruzado
}

```


# The Union method

O método _UnirCom_ retorna um _setReunido_ que consiste em uma união de set e _setDiferente_.

Instanciamos _setReunido_ que será o retorno. E reutilizamos _AddValor_ percorrendo o _set_ original e o _set_ no parâmetro inserindo os valores no _set_ instanciado.

```go
func (set *Set) UnirCom(setDiferente *Set) *Set {

	var setReunido = &Set{}
	setReunido.Criar()
	var valor int

	for valor, _ = range set.estrutura {
		setReunido.AddValor(valor)
	}

	for valor, _ = range setDiferente.estrutura {
		setReunido.AddValor(valor)
	}

	return setReunido
}
```

# The main method

Utilizando a estrutura set na main func.

```go
func main() {

	var set *Set
	set = &Set{}

	set.Criar()

	set.AddValor(1)
	set.AddValor(2)

	fmt.Println("set inicial", set)
	fmt.Println(set.PossuiValor(1))

	var setDiferente *Set
	setDiferente = &Set{}

	setDiferente.Criar()
	setDiferente.AddValor(2)
	setDiferente.AddValor(4)
	setDiferente.AddValor(5)

	fmt.Println("outro set", set)

	fmt.Println("cruzamento de sets", set.CruzarCom(setDiferente))

	fmt.Println("união de sets", set.UnirCom(setDiferente))

}
```
Output:

```text
set inicial &{map[1:true 2:true]}
true
outro set &{map[1:true 2:true]}
cruzamento de sets &{map[2:true]}
união de sets &{map[1:true 2:true 4:true 5:true]}
```