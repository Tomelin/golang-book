# Trabalhando com struct, funções e metodos
A linguagem Go por sua definição não é Orientada a Objeto e por sua vez não possui classe propriamente dita. O Go trabalha no conceito de Struct, talvez podemos fazer um paralelo entre Struct e classes.

Apesar da struct e classe ter alguma semelhança, a struct traz alguns pontos que deixa a linguagem mais robusta e são pontos bem importantes que iremos abordar dos posts. Para quem já trabalhou com orientação a objetos e com classes, pode estranhar no início a sintaxe, mas entendendo bem o conceito, terá um aprendizado rápido, pois realmente é fácil. abaixo teremos alguns exemplos das structs e seus comportamentos.

Vamos declarar uma struct de categoria, para podermos iniciar de forma simples.

## CRIANDO A STRUCT CATEGORIA
Abaixo está um exemplo de uma struct chamada categoria, onde também já colocarei algumas TAGs e explicarei rapidamente.

```
package entity

type Categoria struct {
    ID string `json:"id,omitempty"`              // ID do tipo string
    Name string `json:"name" binding:"required"` // Name do tipo string
    Description string `json:"description"`      // Description do tipo string
}
```
No exemplo acima, declaramos a categoria dentro do package chamado entity. Dentro da nossa struct, temos o ID, Name e Description, todos do tipo string. Ao lado de cada tipo (string) temos o que chamamos de TAG no Go e vamos entendê-las.

**`json:"id,omitempty"`**, aqui informaremos que quando apresentar a struct Categoria no formato json, o id será minúsculo; já o omitempty, irá omitir esse campo, caso o mesmo esteja vazio.
**`json:"name"`**, o campo name, quando apresentado no formato JSON, irá aparecer com o name todo em minúsculo e “mesmo que vazio o campo irá aparecer, pois tem o parâmetro omitempty”.
**`binding:"required"`**, ainda no name, temos a TAG binding, essa TAG irá nos ajudar para informar se o valor é requerido ou não e no caso no name, sim, o valor é requerido. Isso irá ajudar em algumas validações posterioeres;
**`json:"description"`**, já o campo description, não é requerido e irá ser apresentando no formato JSON em minúsculo, mesmo que esteja com o valor vazio.

## CRIANDO A STRUCT PRODUTO

Agora vamos avançar um pouco, vamos criar uma struct chamada produto, mas que possui uma categoria.

```
package entity

type Produto struct {
    ID string `json:"id,omitempty"`
    Name string `json:"name" binding:"required"`
    Description string `json:"description"`
    Categoria *Categoria `json:"categoria"`
}
```

Bom, não iremos comentar os três primeiros campos (ID, Name e Description) pois tem o mesmo comportamento da struct Categoria, vamos entender o objeto chamado Categoria dentro de Produto.

Assim como os outros campos eram do tipo string, o Categoria é do tipo (*Categoria — a primeira struct que criamos), ou seja, possui as mesmas instruções, “quase como se fosse uma herança”, porém ainda assim observamos que tem um * antes do Categoria. Esse * está nos referindo que o objeto Categoria é do tipo endereço de memória da struct Categoria, ou seja, teremos apenas o apontamento da memória de Categoria e não uma cópia dos valores. Isso no fim do dia, está nos informando que iremos economizar memória, pois não precisamos copiar os valores, apenas vamos ler o valor que já está na memória daquele objeto.

Sim, pode ficar um pouco confuso no início, mas em seguida ficará mais claro.

## CRIANDO OS MÉTODOS

Vamos criar os métodos da struct categoria e em seguida entenderemos os apontamentos de memórias.

Já que possuímos a nossa struct, vamos criar 3 métodos para entendermos um pouco, os métodos que criaremos são (Create, List e validator). Sim, dois métodos com a primeira letra maiúscula e o validate com a primeira letra minúscula.

Vamos recordar que no Golang quando colocamos a primeira letra minúscula de uma Struct, Interface, Variável, Constante ou dos Métodos, essa variável será do tipo private (será vista apenas dentro do package), e se colocarmos em maiúscula, o parâmetro declarado será do tipo Public (será visto também fora do package). No nosso exemplo, queremos deixar o validate no formato private, apenas para ser acessado de dentro do package entity.

Vamos ao exemplo:

```
package entity

type Categoria struct {
    ID string `json:"id,omitempty"`              // ID do tipo string
    Name string `json:"name" binding:"required"` // Name do tipo string
    Description string `json:"description"`      // Description do tipo string
}

func (c *Categoria) Create() error{}
func (c *Categoria) List() ([]Category,error){}
func (c *Categoria) validator() (bool,error){}
```
No exemplo acima, criamos 3 métodos que pertence à Categoria. O parâmetro (c *Categoria) após o func, está nos informando que a função (Create, List, validator) pertence à Categoria e que dentro de cada função iremos acessar os valores através do endereço de memória que está representado pela variável c (sim, no Golang é comum declarar variáveis apenas com uma letra), ou seja, tudo que for alterado da Categoria dentro da função (ou seja, na memória), poderá ser acessado na outra função sem problemas, pois apontaremos a variável de memória e não fazendo uma cópia do objeto na totalidade.

## CRIANDO A FUNÇÃO PARA CRIAR UMA NOVA CATEGORIA
Agora, iremos criar uma função chamada NewCategoria, para que essa função receba dois valores (name e description) e retorne toda a struct com o ID preenchido.

```
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
```

## CRIANDO A FUNÇÃO MAIN

Agora podemos criar uma função main, que chame a Catelogria. A nossa função seria algo assim:

```
import (
    "my-package/entity"
    "fmt"
)
func main(){

    categoria, err := entity.Categoria("meu nome", "minhas description")
    if err != nil {
        Panic(err)
    }

    fmt.Println(categoria)
}
```

Na nossa função main acima, iniciamos a Categoria, passando o nome e description, caso de erro podemos tratar o erro e teremos o nosso resultado, os atributos de categoria e o ID.

Nesse exemplo, deixaremos o Create e List em aberto, para evoluirmos posteriormente.

Espero que esse post, ajude a entender melhor trabalhar com struct e ponteiros (endereço de memória).

Deixe aqui o seu comentário, para evoluirmos juntos no conhecimento!