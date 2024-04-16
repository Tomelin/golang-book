# Trabalhando com interfaces

Nesse tópico, vamos falar de interfaces no Golang.  Vou confesar que no início tive bastante problema para entender interfaces no GO e vou tentar trazer de uma forma diferentes que constumamos ver nos exemplos.

Primeiro ponto que precisamos entender, que a interface está ***ligada aos métodos das struct***. Deixa eu reescrever esse ponto, a interface é baseado num ***assinatura***, ou seja, na assinatura de uma struct que possui N métodos.

Vamos ao exemplo, que acredito que ficará bem mais claro. Nesse exemplo, vamos trabalhar para gravar os dados de uma struct de produto, agora, criaremos a nossa struct.

Antes de começar, vamos ter 3 package no nosso exemplo, main, entity e storage.

1. Criando e acessando o diretório;
   ```
   mkdir basic_interface
   cd basic_interface
   ```
2. Vamos iniciar o nosso GO mod;
    ```
    go mod init github.com/tomelin/golang-book/interfaces
    ```
3. vamos criar o nosso package entity, já com as nossas structs
  3.1 Criando o diretório e o arquivo entity;
    ```
    mkdir entity
    cd entity
    touch entity.go
    ```
  3.2 criando a nossa struct
    ```
      package entity

      type Produto struct {
        ID string `json:"id,omitempty"`
        Name string `json:"name" binding:"required"`
        Descricao string `json:"description,omitempty"`
        Preco uint `json:"preco" binding:"required"`
        Estoque uint `json:"estoque" binding:"required"`
      }

      // Essa interface não precisa estar nesse package, pois estar em qualquer outro package.  Para fins de dadicos, deixarei dentro de entity
      type ProdutoInterface interface{
        Create(*Produto) error
        FindAll() ([]Produto, error)
        FindById(id *string) (*Produto,error)
        Delete(id *string) error
      }
    ```

Nesse nosso exemplo, a struct produto e a interface ProdutoInterface, estão em letra maiuscula, pois são "publicas" ou acessadas de outros packages.

**Perceba, criamos a interface e não conectamos em nenhum momento a struct, há não ser o nome que contem produto** 
Esse é o ponto importante a ser entendido, nesse momento a interface, não está vinculada a nenhuma struct, apenas tem o que chamamos assinatura, e essa assinatura possui o seguinte:
```
Create(*Produto) error
FindAll() ([]Produto, error)
FindById(id *string) (*Produto,error)
Delete(id *string) error
```

Essa assinatura nos informa que toda e qualquer struct que tiver os métodos acima, podemos usar a interface ProdutoInterface

Vamos agora, criar o package storage e dentro de storage, teremos 3 structs para trabalharmos com (database, cache, files). Para poder usar a interface que criamos, as structs que criamos precisarão ter a mesma assinatura.

Vamos entender conforme o nosso exemplo:

1. vamos criar o nosso package storage, com a struct para trabalhar com database
```
  mkdir storage
  cd storage
  touch storage.go
```
2. Agora iremos criar uma struct para trabahar com database
  2.1
    ```
      package storage
      
      type ProdutoStorageDB struct {
        DB *sql.DB
      }
      
      // Aqui estamos recebendo a conexão do SQL e associando a struct.  e estmos retornando a nossa struct ProdutoStorageDB, porém estamos informando que ela é do tipo entity.ProdutoInterface (ou possui as assinaturas) que a interface precisa para trabalhar
      func NewProdutoStorageDB(db *sql.DB) entity.ProdutoInterface{
        return &ProdutoStorageDB{DB: db}
      }
    ```
    Nesse momento a struct ProdutoStorageDB, não possui a assinatura para trabalhar com a interface entity.ProdutoInterface, ou seja, nesse momento teremos um erro de assinatura.
  
    Agora vamos implementar a nossa assinatura para trabalharmos com a interface

  2.2 Implementando a assinatura a nossa struct ProdutoStorageDB, continuaremos com o nosso exemplo acima
    ```
      package storage
      
      type ProdutoStorageDB struct {
        DB *sql.DB
      }

      // Aqui estamos recebendo a conexão do SQL e associando a struct.  e estmos retornando a nossa struct ProdutoStorageDB, porém estamos informando que ela é do tipo entity.ProdutoInterface (ou possui as assinaturas) que a interface precisa para trabalhar
      func NewProdutoStorageDB(db *sql.DB) entity.ProdutoInterface{
        return &ProdutoStorageDB{DB: db}
      }

      func (p *ProdutoStorageDB) Create(*Produto) error {return nil}
      func (p *ProdutoStorageDB) FindAll() ([]Produto, error) {return nil,nil }
      func (p *ProdutoStorageDB) FindById(id *string) (*Produto,error)  {return nil,nil }
      func (p *ProdutoStorageDB) Delete(id *string) error {return nil }
    ```
   Aqui é onde realmente acontece a mágica, a struct ProdutoStorageDB, contém os métodos necessário para implementar a intefrcae entity.ProdutoInterface, ou seja, nesse caso já podemos usar a nossa interface.

3. Agora iremos criar uma struct para trabahar com cache redis
    3.1
     ```
       package storage
  
       type ProdutoStorageCache struct {
       ConnectionString *string
       }
  
       // Aqui estamos recebendo a conexão do SQL e associando a struct.  e estmos retornando a nossa struct ProdutoStorageCache, porém estamos informando que ela é do tipo entity.ProdutoInterface (ou possui as assinaturas) que a interface precisa para trabalhar
       func NewProdutoStorageDB(connection *string) entity.ProdutoInterface{
       return &ProdutoStorageCache{ConnectionString: connection}
       }
     ```
     Da mesma forma que o exemplo do item 2, nesse momento a struct ProdutoStorageCache não está preparada para trabalhar com a nossa interface entity.ProdutoInterface, simplesmente porque não possui os métodos necessários.
      
     Agora vamos implementar a nossa assinatura para trabalharmos com a interface.
    3.2
       ```
         package storage
    
         type ProdutoStorageCache struct {
         ConnectionString *string
         }
  
         // Aqui estamos recebendo a conexão do SQL e associando a struct.  e estmos retornando a nossa struct ProdutoStorageCache, porém estamos informando que ela é do tipo entity.ProdutoInterface (ou possui as assinaturas) que a interface precisa para trabalhar
         func NewProdutoStorageDB(connection *string) entity.ProdutoInterface{
         return &ProdutoStorageCache{ConnectionString: connection}
         }
  
         func (p *ProdutoStorageCache) Create(*Produto) error {return nil}
         func (p *ProdutoStorageCache) FindAll() ([]Produto, error) {return nil,nil }
         func (p *ProdutoStorageCache) FindById(id *string) (*Produto,error)  {return nil,nil }
         func (p *ProdutoStorageCache) Delete(id *string) error {return nil }
       ```
      Observe que agora temos praticamente o mesmo exemplo que ProdutoStorageDB, o que muda é a struct no qual associamos os métodos.

      **O que realmente importa entender, é que tanto a struct ProdutoStorageDB e ProdutoStorageCache possuem os mesmos métodos para a assinatura da interface**