# Dominando as variáveis em Golang

A linguagem Go é fortemente tipada. Uma linguagem fortemente tipada, significa que suas constantes e/ou variáveis são de um tipo (float, number, string, …) e durante a existência dessas variáveis, o tipo delas não mudam, ou seja, uma variável que nasce do tipo string, morre do tipo string sem ser alterada o tipo da mesma no meio do caminho. Veremos alguns exemplos sobre variáveis logo em seguida.

Vou me basear que usaremos linux, MacOs ou WSL, para criar a estrutura de diretórios e arquivos, mas caso não queira fazer nenhum procedimento na sua máquina, pode usar o Golang playgorund no seguinte link: https://go.dev/play

1. crie o diretório com o nome de 1-sintaxe, com o seguinte comando: "mkdir 1-sintaxe";
2. acessando o diretório com o comando cd: "cd 1-sintaxe";
3. crie um novo arquivo chamado main.go, podendo usar o comando: "touch main.go";
4. agora, vamos editar o arquivo, vou ter como base que estamos usando o Visual Code, então dentro do diretório 1-sintaxe execute o comando "code .";
5. Se seu visual code está instalado na sua máquina, deverá abrir uma nova página do visual code.
6. 
Agora, estando com nossa estrutura criada, ou usando o Golang playground, iniciaremos nossos exemplos de variáveis.

Dentro do arquivo main.go, vamos começar criando a seguinte estrutura, lembrando que no Golang, toda estrutura nova, deve iniciar com o arquivo chamado main.go, package main e a função (func) main

```
package main
import "fmt"

func main() {
    fmt.Println("Hello")
}
```
No exemplo acima, conseguimos identificar de forma clara o package main, função func main() e dentro da função um fmt.Print que é o famoso teste hello world. Para testar o código, basta executar o comando

```go run main.go```

O package no golang é como se fosse o namespace em algumas linguagens, ele que indica onde aquela instrução (package) executará dentro do seu sistema, veremos os packages um pouco adiante.

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

Acima temos alguns exemplos de variáveis que podemos utilizar. Para que possamos criar constantes no Golang, basta alterar o var por const, conforme os exemplos abaixo:

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

Uma observação muito importante, é que no Golang temos variáveis do tipo public e private. O que muda entre public e private é a primeira letra, se for MAIÚSCULA é do tipo public se for MINÚSCULA é do tipo privada.

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

Veremos mais adiante, mas o padrão de public e private no Golang através de letra MAIÚSCULA e MINÚSCULA, será adotado nas funções (fuc), struct, methods, interfaces, …

Contudo, esse é apenas um breve resumo sobre variáveis no Golang, recomendo fortemente dar uma olhada na documentação. Deixarei dos links da documentação do Golang para consultas:

[https://go.dev/ref/spec#Variables](https://go.dev/ref/spec#Variables)

[https://go.dev/doc/effective_go](https://go.dev/doc/effective_go)