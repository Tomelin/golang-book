
1. Sintaxe básica da linguagem golang	1
1. 1 Tipos de dados e variáveis	1
1. 4 Struct, funções e métodos	1
   Sintaxe básica da linguagem golang
1. 1 Dominando as Variáveis em Golang
   A linguagem Go é fortemente tipada. Uma linguagem fortemente tipada, significa que suas constantes e/ou variáveis são de um tipo (float, number, string, …) e durante a existência dessas variáveis, o tipo delas não mudam, ou seja, uma variável que nasce do tipo string, morre do tipo string sem ser alterada o tipo da mesma no meio do caminho. Veremos alguns exemplos sobre variáveis logo em seguida.
   Vou me basear que estamos usando linux, MacOs ou WSL, para criar a estrutura de diretórios e arquivos, mas caso não queira fazer nenhum procedimento na sua máquina, pode usar o Golang playgorund no seguinte link: https://go.dev/play
   crie o diretório com o nome de 1-sintaxe, com o seguinte comando: `mkdir 1-sintaxe`;
   acessando o diretório com o comando cd: `cd 1-sintaxe`;
   crie um novo arquivo chamado main.go, podendo usar o comando: `touch main.go`;
   agora, vamos editar o arquivo, vou ter como base que estamos usando o Visual Code, então dentro do diretório 1-sintaxe execute o comando `code .`;
   Se seu visual code está instalado na sua máquina, deverá abrir uma nova página do visual code.

Agora, estando com nossa estrutura criada, ou usando o Golang playground, vamos iniciar nossos exemplos de variáveis.
Dentro do arquivo main.go, vamos começar criando a seguinte estrutura, lembrando que no Golang, toda estrutura nova, deve iniciar com o arquivo chamado main.go, package main e a função (func) main
```
package main

import "fmt"

func main() { 
	fmt.Println("Hello")
}

```
No exemplo acima, conseguimos identificar de forma clara o package main, função ***func main()*** e dentro da função um ***fmt.Print*** que é o famoso teste "hello world".  Para testar o código, basta executar o comando ```go run main.go```.
O package no golang é como se fosse o namespace em algumas linguagens, ele que indica onde aquela instrução (package) está sendo executada dentro do seu sistemas, veremos os packages um pouco adiante.
Vamos agora brincar com as nossas variáveis, que é o nosso objetivo desse tópico, mas antes precisamos que temos basicamente 3 formas de declarar uma variável, que são:
```
package main

import "fmt"

var (
	myVar1 string
)

func main() {
	var myVar2 string
	myVar3 := "variavel do tipo string"
	fmt.Println(myVar1)
	fmt.Println(myVar2)
	fmt.Println(myVar3)
}
```

A primeira variável, foi declarada fora da função main, ou seja, está no escopo do package, a segunda está declarada dentro da função main, mas apenas foi inicializada, sem nenhum valor dentro da mesma e a terceira, declaramos passando um texto para dentro.
Vamos recriar o código acima, agora declarando todas as variáveis com informações dentro


Vamos ver mais alguns exemplos e declaramos as variáveis apenas dentro da função main:
``` 
package main

import "fmt"

func main() {

	varString := "variável do tipo string"
	fmt.Println(varString) // Estamos usando o fmt.Println, pois o Println faz uma quebra de linha no final

	varString2 := `de string`
	fmt.Printf("interpolação %s \n", varString2) // Estamos usando o fmt.Println, pois o Println faz uma quebra de linha no final

	varInt := 5 // variável do tipo inteiro, mas pode ser valor negativo
	fmt.Println(varInt)

	varInt = -5 // varInt recebendo valor negativo
	fmt.Println(varInt)

	var varUint uint // variável apenas com números positivos
	varUint = 0
	fmt.Println(varUint)

	//varUint = -5 // retornará erro: cannot use -5 (untyped int constant) as uint value in assignment (overflows)
	//fmt.Println(varUint)

	varBool := true // variável do tipo boolean
	fmt.Println(varBool)

	varStrintList := []string{"value1", "value2"} // variável do tipo list de string
	fmt.Println(varStrintList)

	varUintList := []uint{1, 10} // variável do tipo list de inteiros
	fmt.Println(varUintList)

	var varInterface interface{} // variável do tipo interface, semelhante ao any
	varInterface = "pode receber qualquer tipo de valor"
	fmt.Println(varInterface)

	varInterface = 25 // inclusive pode mudar o tipo NÃO RECOMENDADO
	fmt.Println(varInterface)

	varMap := make(map[string]int) // variável do tipo map string que recebe o valor numérico
	varMap["key1"] = 5
	varMap["key2"] = 10
	fmt.Println(varMap)

}

```

Acima temos alguns exemplos de variáveis que podemos utilizar.  Para que possamos criar constantes no Golang, basta alterar o var por const, conforme os exemplos abaixo:

```
package main

import "fmt"

func main() {

	const constString string = "constante do tipo string"
	fmt.Println(constString) // Uma constante não pode ser alterada durante seu tempo de vida

	const constInt int = 5 // constante do tipo inteiro, mas pode ser valor negativo
	fmt.Println(constInt)

	const constUint uint = 0 // constante apenas com números positivos
	fmt.Println(constUint)
}
```

Podemos ver que no Golang a declaração de variáveis e constante são simples. Lembrando que quando uma variável/constante estiver fora de uma função, o valor está disponível para todo o package.

Uma observação muito importante, é que no Golang temos variáveis do tipo public e private.  O que muda entre public e private é a primeira letra, se for MAIÚSCULA é do tipo public se for MINÚSCULA é do tipo privada.
Vamos a alguns exemplos:

```
package main

import "fmt"

func main() {

	const ConstPublic string = "constante do tipo string"
	fmt.Println(ConstPublic) // constante Pública, pois tem a primeira letra MAIÚSCULA, pode ser acessado de qualquer lugar do sistema

	const constPublic string = "constante do tipo string"
	fmt.Println(constPublic) // constante privada, pois tem a primeira letra MINÚSCULA, restrito ao package

	var VarInt int = 5 // variável Pública, pois tem a primeira letra MAIÚSCULA, podendo ser acessada de qualquer lugar do sistema
	fmt.Println(VarInt)

	var varBool bool = true // variável privada, pois tem a primeira letra MINÚSCULA, restrito ao package
	fmt.Println(varBool)

}
 
```
No exemplo acima de variáveis public e private é fácil de identificar, apenas pela primeira letra tornando muito simples a declaração da mesma, inclusive ajudando na padronização de escrita de código.
Veremos mais à frente, mas o padrão de public e private no Golang através de letra MAIÚSCULA e MINÚSCULA, será adotado nas funções (fuc), struct, methods, interfaces, …

Contudo, esse é apenas um breve resumo sobre variáveis no Golang, recomendo fortemente dar uma olhada na documentação.  Deixarei dos links da documentação do Golang para consultas:
https://go.dev/ref/spec#Variables
https://go.dev/doc/effective_go

1. 2 Operadores e expressões em Golang
   Dando continuidade na sintaxe básica do Golang, hoje vamos ver como o Golang trata os operadores e expressões aritméticas.
   Antes de iniciarmos, vamos entender primeiro o que são operadores e expressões aritméticas:
   “Expressões aritméticas são formas de representar as operações matemáticas em que os operadores são aritméticos (sinais matemáticos) e os operandos são valores do tipo numérico, que podem ser inteiros ou reais (float)“. (https://www.terra.com.br/noticias/educacao/enem/aprenda-a-resolver-expressoes-aritmeticas-de-maneira-simples,f24dc01376626cd7a4e0b29bfe424330hivnu3ro.html)
   1.2.1 SOMA
   Vamos começar pela a soma, que é um dos operadores mais utilizados e de fácil entendimento.
```
package main


import "fmt"


const (
   value01 int8 = 10
   value02 int8 = -7
)


var (
   float01 float32 = 10.7
   float02 float32 = 0.3
)


func main() {


   integerResult := value01 + value02
   fmt.Println("O resultado da soma de int é: ", integerResult)


   floatResult := float01 + float02
   fmt.Printf("O resultado da soma de int é: %.2f", floatResult)


   var valueUint01, valueUint02 uint
   valueUint01 = 2
   valueUint02 = 0
   fmt.Printf("\nO resultado da soma de uint é: %d", valueUint01+valueUint02)
}

```

No exemplo acima temos 3 somas, sendo realizadas de forma diferentes, e vamos explicar cada uma delas:
A primeira soma é realizada através de constantes (const), ou seja, esse valor não mudará ao longo da execução do processo, após a declaração, realizamos uma soma básica colocando o resultado na variável de nome integerResult e por último, realizamos o print do resultado na tela, que nesse exemplo o resultado será 3.  É muito importante ressaltar que o int8 na linha 7, aceita valores negativos.
A segunda soma é realizada através de variáveis, ou seja, o valor da variável pode mudar ao longo da execução do programa.  Realizamos uma soma básica, igual ao exemplo anterior e depois imprimimos, o resultado ficou 11.00.  Os dois zeros no final, ficou pois quando imprimimos com o fmt.Printf, informamos ao Go para colocar duas casas decimais após o zero com o parâmetro "%.2f".
Por último, mas não menos importante, declaramos duas variáveis apenas dentro do escopo da função main do tipo uint.  O tipo uint não aceita valores negativos.  Essa opção é interessante, quando estamos falando de salário, notas da escola que não podem ser menores que 0, …

O resultado da execução que devemos ter, é semelhante a imagem abaixo:

```
➜  cmd go run main.go
O resultado da soma de int é:  3
O resultado da soma de int é: 11.00
O resultado da soma de uint é: 2

```

1.2.1 SUBTRAÇÃO
A subtração é muito semelhante ao exemplo anterior, exceto que o resultado pode retornar valores negativos em algum momento, ou seja, temos que cuidar ao usar o type uint.
Vamos nos basear no exemplo anterior, exceto por declarar as variáveis de formas diferentes ou em escopos diferentes, apenas para que possamos ir treinando algumas lógicas básicas.

package main


import "fmt"


const (
value01 int8 = 10
value02 int8 = -7
)


var (
float01 float32 = 10.7
float02 float32 = 0.
)


func main() {


integerResult := value01 - value02
fmt.Println("O resultado da subtração de int é: ", integerResult)


floatResult := float01 - float02
fmt.Printf("O resultado da subtração de int é: %.2f", floatResult)


if floatResult < 11 {
var valueUint01, valueUint02 uint
valueUint01 = 2
valueUint02 = 0
fmt.Printf("\nO resultado da subtração de uint é: %d", valueUint01+valueUint02)
}
}

A primeira subtração é realizada através de constantes, ou seja, o valor da constante não mudará ao longo da execução do processos.  Realizamos uma subtração básica, igual ao exemplo anterior e depois imprimimos, o resultado ficou 17.  O valor ficou 17, pois tínhamos um valor positivo (10) e subtraímos um valor negativo (-7) no qual acabou inverntendo o sinal de - para +, pois menos com menos é mais (10 - (-7)).
A segunda subtração é realizada através de variáveis, ou seja, o valor da variável pode mudar ao longo da execução do programa.  Realizamos a subtração básica, igual ao exemplo anterior e depois imprimimos, o resultado ficou 10.70.  Os dois zeros no final, ficou pois quando imprimimos com o fmt.Printf, informamos ao Go para colocar duas casas decimais após o zero com o parâmetro "%.2f".
A terceira subtração, colocamos dentro de um if, apenas para mostrar que os valores valueUint01 e valueUint02 estão dentro do if, nesse contexto o escopo da variável será somente dentro do if, ou seja, fora do if a variável não existirá.
1.2.2 MULTIPLICAÇÃO
A multiplicação acaba sendo simples também e vamos nos basear no exemplo anterior.
package main


import "fmt"


const (
value01 int16 = 10
value02 int16 = -7
)


var (
float01 float32 = 10.7
float02 float32 = 0.1
)


func main() {


integerResult := value01 * value02
fmt.Println("O resultado da subtração de int é: ", integerResult)


floatResult := float01 * float02
fmt.Printf("O resultado da subtração de int é: %.2f", floatResult)


var valueUint01, valueUint02 uint8
if floatResult > 11 {
valueUint01 = 2
valueUint02 = 1
}
fmt.Printf("\nO resultado da subtração de uint é: %d", valueUint01*valueUint02)
}


A primeira multiplicação é realizada através de constantes, ou seja, o valor da constante não mudará ao longo da execução do processos.  Realizamos uma multiplicação básica, igual ao exemplo anterior e depois imprimimos, o resultado ficou -70.  O valor ficou -70, pois tínhamos um valor positivo (10) e multiplicamos um valor negativo (-7)  ou seja, a consulta seria igual ao exemplo ao lado: (10 * (-7)).
A segunda multiplicação é realizada através de variáveis, ou seja, o valor da variável pode mudar ao longo da execução do programa.  Realizamos a multiplicação básica, igual ao exemplo anterior e depois imprimimos, o resultado ficou 1.07.  Os dois zeros no final, ficou pois quando imprimimos com o fmt.Printf, informamos ao Go para colocar duas casas decimais após o zero com o parâmetro "%.2f".
A terceira multiplicação, declaramos as variáveis fora do IF e somente se a condicional for TRUE, entraremos no IF e configuramos os valores, caso contrário o valor será zero, pois o Golang quando inicia um int o seu valor inicial será sempre 0.

O resultado da execução que devemos ter, é semelhante a imagem abaixo:
```
➜  cmd go run main.go
O resultado da subtração de int é:  -70
O resultado da subtração de int é: 1.07
O resultado da subtração de uint é: 0
```

1.2.3 DIVISÃO
Conforme os exemplos anteriores, continuaremos com o mesmo código:
package main


import "fmt"


const (
value01 int16 = 10
value02 int16 = -7
)


var (
float01 float32 = 10.7
float02 float32 = 0.1
)


func main() {


integerResult := value01 / value02
fmt.Println("O resultado da subtração de int é: ", integerResult)


floatResult := float01 / float02
fmt.Printf("O resultado da subtração de int é: %.2f", floatResult)


var valueUint01, valueUint02 uint8
if floatResult > 11 {
valueUint01 = 2
valueUint02 = 1
}
fmt.Printf("\nO resultado da subtração de uint é: %d", valueUint01/valueUint02)
}



O resultado da execução que devemos ter, é semelhante a imagem abaixo:
```
➜  cmd go run main.go
O resultado da subtração de int é:  -1
O resultado da subtração de int é: 107.00
O resultado da subtração de uint é: 2
```

1.2.3 MÓDULO
O módulo é muito útil para identificar se um número é dividido por outro, sendo o resultado 0, ou se um número é divisível por outro.  Utilizamos muito módulo para saber se um número é par ou ímpar.
package main


import "fmt"


var (
value01 uint = 10
value02 uint = 5
value03 uint = 3
value04 uint = 21
)


func main() {


if value01%value02 == 0 {
fmt.Printf("O resultado do modulo é par: %d", value01%value02)
}


if value01%value03 != 0 {
fmt.Printf("\nO resultado do modulo é impar: %d", value01%value03)
}


if value04%value03 == 0 {
fmt.Printf("\nO número 21 é divisivel por 3, o resultado é: %d", value04%value03)
}
}

No primeiro exemplo, estamos validando se o número 10 é divisível por 5 e o resultado é verdadeiro;
No segundo exemplo, observem que mudamos a condicional, ou seja, só entrará no IF se o resultado for diferente de 0.
e o terceiro exemplo, estamos validando se o número 21 é divisível por 3 e como veremos o resultado, sim, 21 é divisível por 3.
```
➜  cmd go run main.go
O resultado do modulo é par: 0
O resultado do modulo é impar: 1
O resultado do módulo é impar: 0
```

Além dos exemplos acima, o Golang tem o package math 'https://pkg.go.dev/math', que contém muitos cálculos matemáticos.

Esses foram alguns exemplos de operações matemáticas na qual conseguimos realizar com o Golang.


1. 3 Struct, funções e métodos
   A linguagem Go por sua definição não é Orientada a Objeto e por sua vez não possui classe propriamente dita. O Go trabalha no conceito de Struct, talvez podemos fazer um paralelo entre Struct e classes.
   Apesar da struct e classe ter alguma semelhança, a struct trás alguns pontos que deixa a linguagem mais robusta e são pontos bem importantes que iremos abordar dos posts. Para quem já trabalhou com orientação a objetos e com classes, pode estranhar no início a sintaxe, mas entendendo bem o conceito, terá um aprendizado rápido, pois realmente é fácil. abaixo teremos alguns exemplos das structs e seus comportamentos.

Vamos declarar uma struct de categoria, para podermos iniciar de forma simples.

CRIANDO A STRUCT CATEGORIA
Abaixo está um exemplo de uma struct chamada categoria, onde também já colocarei algumas TAGs e explicarei rapidamente.

package entity

type Categoria struct {
ID string `json:"id,omitempty"`              // ID do tipo string
Name string `json:"name" binding:"required"` // Name do tipo string
Description string `json:"description"`      // Description do tipo string
}


No exemplo acima, declaramos a categoria dentro do package chamado entity.  Dentro da nossa struct, temos o ID, Name e Description, todos do tipo string.  Ao lado de cada tipo (string) temos o que chamamos de TAG no Go e vamos entendê-las.
`json:"id,omitempty"`,  aqui estamos informando que quando apresentar a struct Categoria no formato json, o id será minúsculo;  já o omitempty, irá omitir esse campo, caso o mesmo esteja vazio.
`json:"name"`, o campo name, quando apresentado no formato JSON, irá aparecer com o name todo em minúsculo e "mesmo que vazio o campo irá aparecer pois tem tem o parâmetro omitempty".  
`binding:"required"`, ainda no name, temos a TAG binding, essa TAG irá nos ajudar para informar se o valor é requerido ou não e no caso no name, sim o valor é requerido.  Isso irá ajudar em algumas validações posterioeres;
`json:"description"`, já o campo description, não é requerido e irá ser apresentando no formato JSON em minúsculo, mesmo que esteja com o valor vazio.

CRIANDO A STRUCT PRODUTO
Agora vamos avançar um pouco, vamos criar uma struct chamada produto mas que possui uma categoria.

package entity

type Produto struct {
ID string `json:"id,omitempty"`
Name string `json:"name" binding:"required"`
Description string `json:"description"`
Categoria *Categoria `json:"categoria"`
}


Bom, não iremos comentar os três primeiros campos (ID, Name e Description) pois tem o mesmo comportamento da struct Categoria, vamos entender o objeto chamado Categoria dentro de Produto.

Assim como os outros campos eram do tipo string, o Categoria é do tipo (*Categoria - a primeira struct que criamos), ou seja, possui as mesmas instruções, "quase como se fosse uma herança", porém ainda assim observamos que tem um * antes do Categoria. Esse * está nos referindo que o objeto Categoria é do tipo endereço de memória da struct Categoria, ou seja, teremos apenas o apontamento da memória de Categoria e não uma cópia dos valores.  Isso no fim do dia, está nos informando que iremos economizar memória, pois não precisamos copiar os valores, apenas vamos ler o valor que já está na memória daquele objeto.

Sim, pode ficar um pouco confuso no início, mas em seguida ficará mais claro.

CRIANDO OS MÉTODOS

Vamos criar os métodos da struct categoria e em seguida iremos entender os apontamentos de memórias.

Já que possuímos a nossa struct, vamos criar 3 métodos para que possamos entender um pouco, os métodos que criaremos são (Create, List e validator).   Sim, dois métodos com a primeira letra maiúscula e o validate com a primeira letra minúscula.
Vamos recordar que no Golang quando colocamos a primeira letra minúscula de uma Struct,Interface, Variável, Constante ou dos Métodos, essa variável será do tipo private (será vista apenas dentro do package), e se colocarmos em maiúscula, o parâmetro declarado será do tipo Public (será visto também fora do package).   No nosso exemplo, queremos deixar o validate no formato private, apenas para ser acessado de dentro do package entity.

Vamos ao exemplo:

package entity

type Categoria struct {
ID string `json:"id,omitempty"`              // ID do tipo string
Name string `json:"name" binding:"required"` // Name do tipo string
Description string `json:"description"`      // Description do tipo string
}

func (c *Categoria) Create() error{}
func (c *Categoria) List() ([]Category,error){}
func (c *Categoria) validator() (bool,error){}


No exemplo acima, criamos 3 métodos que pertence a Categoria. o parâmetro (c *Categoria) após o func, está nos informando que a função (Create, List, validator) pertence a Categoria e que dentro de cada função iremos acessar os valores através do endereço de memória que está representado pela variável c (sim, no Golang é comum declarar variáveis apenas com uma letra), ou seja tudo que for alterado da Categoria dentro de uma função (ou seja, na memória), poderá ser acessado na outra função sem problemas, pois estamos apontando a variável de memória e não fazendo uma cópia do objeto como um todo.

Agora, iremos criar uma função chamada NewCategoria, para que essa função receba dois valores (name e description) e retorne toda a struct com o ID preenchido.


package entity

type Categoria struct {
ID string `json:"id,omitempty"`              // ID do tipo string
Name string `json:"name" binding:"required"` // Name do tipo string
Description string `json:"description"`      // Description do tipo string
}
func NewCategoria(name, description string) (*Categoria, error){


// Aqui estamos instanciando a struct e passando os valores para ela
// O & comercial na frente de Categoria, indica que estamos instanciado o objeto, mas a partir do endereço de memória do mesmo
c := &Categoria{
Name: name
Description: description
}


// Como o validator pertence a struct Categoria, podemos acessar a partir da variável c que é do tipo Categoria
err := c.validator()
if err != nil {
return nil, err
}

// Após a chamado do c.validator(), podemos acessar o ID aqui nesse nessa função, pois alteramos a variável no endereço de memoria

return c, nil
}
func (c *Categoria) Create() error{}
func (c *Categoria) List() ([]Category,error){}
func (c *Categoria) validator() error {

// Estamos validando se o Name está vazio
if c.Name == ""{
return errors.New("o campo name não pode ser vazio")
}

// Estamos validando se o Description está vazio a partir do package reflect
if reflect.ValueOf(c.Description).IsZero(){
return errors.New("o campo description não pode ser vazio")
}

// Estamos passando um ID caso o mesmo esteja vazio
if reflect.ValueOf(c.ID).IsZero(){
c.ID = uuid.New()
}

return nil
}



Agora podemos criar uma função main, que chame a Catelogria.  Nossa função seria algo assim:

import (
"entity"
"fmt"
)
func main(){

categoria, err := entity.Categoria("meu nome", "minhas description")
if err != nil {
Panic(err)
}

fmt.Println(categoria)
}


Na nossa função main acima, iniciamos a Categoria, passando o nome e description, caso de erro podemos tratar o erro e teremos o nosso resultado final, os atributos de categoria e o ID.

Nesse exemplo, deixaremos o Create e List em aberto, para que possamos evoluir posteriormente.

Espero que esse post, ajude a entender melhor trabalhar com struct e ponteiros (endereço de memória).

Deixe aqui seu comentário, para evoluirmos juntos no conhecimento!

Sintaxe Básica da Linguagem
Tipos de dados e variáveis.
Operadores e expressões.
Controle de fluxo: instruções if, for, switch, select.
Funções e métodos.
Pacotes que precisamos conhecer
strings, strconv, io. net/http, manipulação de arquivos, network, stdout, slog, fmt, log, error, Defer, panic e recover, reflect
Context
Pacotes
Goroutines e Canais
Introdução à concorrência e paralelismo
Goroutines e programação concorrente
Canais e comunicação entre goroutines
Padrões de concorrência: pipeline, fan-in, fan-out
pipeline
Estruturas de Dados
Arrays e slices.
Maps e structs, struct tags
Interfaces e tipos customizados, Interfaces e polimorfismo
Manipulação JSON e YAML
Tests, docs, benachmark, fuzzy
gRPC
Protobuf
Interface (interface, 3 struct (database, redis, file)
Protofile
API/Gin
Kubernetes operators


wg, mutex, sync, lock






Sintaxe Básica da Linguagem

