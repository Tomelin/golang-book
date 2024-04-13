# Golang, operadores e expressões aritméticas
Dando continuidade na sintaxe básica do Golang, hoje vamos ver como o Golang trata os operadores e expressões aritméticas.

Antes de iniciarmos, vamos entender primeiro o que são operadores e expressões aritméticas:

***Expressões aritméticas são formas de representar as operações matemáticas em que os operadores são aritméticos (sinais matemáticos) e os operandos são valores do tipo numérico, que podem ser inteiros ou reais (float)***. [aprenda-a-resolver-expressoes-aritmeticas-de-maneira-simples](https://www.terra.com.br/noticias/educacao/enem/aprenda-a-resolver-expressoes-aritmeticas-de-maneira-simples,f24dc01376626cd7a4e0b29bfe424330hivnu3ro.html)

## SOMA
Vamos começar pela soma, que é um dos operadores mais utilizados e de fácil entendimento.

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
No exemplo acima temos 3 somas, sendo realizadas de formas diferentes, e vamos explicar cada uma delas:

1. A primeira soma é realizada por constantes (const), ou seja, esse valor não mudará ao longo da execução do processo, após a declaração, realizamos uma soma básica colocando o resultado na variável de nome integerResult e por último, realizamos o print do resultado na tela, que nesse exemplo o resultado será 3. É muito importante ressaltar que o int8 na linha 7, aceita valores negativos.
2. A segunda soma é realizada por variáveis, ou seja, o valor da variável pode mudar ao longo da execução do programa. Realizamos uma soma básica, igual ao exemplo anterior e depois imprimimos, o resultado ficou 11.00. Os dois zeros no final, ficou, pois quando imprimimos com o fmt.Printf, informamos ao Go para colocar duas casas decimais após o zero com o parâmetro “%.2f”.
3. Por último, mas não menos importante, declaramos duas variáveis apenas dentro do escopo da função main do tipo uint. O tipo uint não aceita valores negativos. Essa opção é interessante, quando falaremos de salário, notas da escola que não podem ser menores que 0, …

O resultado da execução que devemos ter, é semelhante ao resultado abaixo:

```go run main.go```

O resultado da soma de int é: 3
O resultado da soma de int é: 11.00
O resultado da soma de uint é: 2

## SUBTRAÇÃO
A subtração é muito semelhante ao exemplo anterior, exceto que o resultado pode retornar valores negativos em algum momento, ou seja, temos que cuidar ao usar o type uint.

Vamos nos basear no exemplo anterior, exceto por declarar as variáveis de formas diferentes ou em escopos diferentes, apenas para treinarmos algumas lógicas básicas.

```
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
```

1. A primeira subtração é realizada por constantes, ou seja, o valor da constante não mudará ao longo da execução do processo. Realizamos uma subtração básica, igual ao exemplo anterior e depois imprimimos, o resultado ficou 17. O valor ficou 17, pois tínhamos um valor positivo (10) e subtraímos um valor negativo (-7) no qual acabou inverntendo o sinal de — para +, pois menos com menos é mais (10 — (-7)).
2. A segunda subtração é realizada por variáveis, ou seja, o valor da variável pode mudar ao longo da execução do programa. Realizamos a subtração básica, igual ao exemplo anterior e depois imprimimos, o resultado ficou 10.70. Os dois zeros no final, ficou, pois quando imprimimos com o fmt.Printf, informamos ao Go para colocar duas casas decimais após o zero com o parâmetro “%.2f”.
3. A terceira subtração, colocamos dentro de if, apenas para mostrar que os valores valueUint01 e valueUint02 estão dentro do if, nesse contexto o escopo da variável será somente dentro do if, ou seja, fora do if a variável não existirá.

## MULTIPLICAÇÃO
A multiplicação acaba sendo simples também e vamos nos basear no exemplo anterior.

```
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
```

1. A primeira multiplicação é realizada por constantes, ou seja, o valor da constante não mudará ao longo da execução do processo. Realizamos uma multiplicação básica, igual ao exemplo anterior e depois imprimimos, o resultado ficou -70. O valor ficou -70, pois tínhamos um valor positivo (10) e multiplicamos um valor negativo (-7) ou seja, a consulta seria igual ao exemplo ao lado: (10 * (-7)).
2. A segunda multiplicação é realizada por variáveis, ou seja, o valor da variável pode mudar ao longo da execução do programa. Realizamos a multiplicação básica, igual ao exemplo anterior e depois imprimimos, o resultado ficou 1.07. Os dois zeros no final, ficou, pois quando imprimimos com o fmt.Printf, informamos ao Go para colocar duas casas decimais após o zero com o parâmetro “%.2f”.
3. A terceira multiplicação, declaramos as variáveis fora do IF e somente se a condicional for TRUE, entraremos no IF e configuramos os valores, caso contrário o valor será zero, pois o Golang quando inicia um int o seu valor inicial será sempre 0.

O resultado da execução que devemos ter, é semelhante ao resultado abaixo:

```go run main.go```
O resultado da subtração de int é: -70
O resultado da subtração de int é: 1.07
O resultado da subtração de uint é: 0

## DIVISÃO
Conforme os exemplos anteriores, continuaremos com o mesmo código:

```
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
```
O resultado da execução que devemos ter, é semelhante ao resultado abaixo:

```go run main.go```
O resultado da subtração de int é: -1
O resultado da subtração de int é: 107.00
O resultado da subtração de uint é: 2

## MÓDULO
O módulo é muito útil para identificar se um número é dividido por outro. Quando o resultado do módulo for 0 quer dizer que um número é divisivel por outro. Utilizamos muito módulo para saber se um número é par ou ímpar.

```
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
```
1. No primeiro exemplo, estamos validando se o número 10 é divisível por 5 e o resultado é verdadeiro;
2. No segundo exemplo, observem que mudamos a condicional, ou seja, só entrará no IF se o resultado for diferente de 0.
3. e o terceiro exemplo, estamos validando se o número 21 é divisível por 3 e como veremos o resultado, sim, 21 é divisível por 3.

```go run main.go```

O resultado do modulo é par: 0
O resultado do modulo é impar: 1
O resultado do módulo é impar: 0


## PACKAGE MATH
Além dos exemplos acima, o Golang tem o package math ‘https://pkg.go.dev/math', que contém muitos cálculos matemáticos.

Esses foram alguns exemplos de operações matemáticas na qual conseguimos realizar com o Golang.