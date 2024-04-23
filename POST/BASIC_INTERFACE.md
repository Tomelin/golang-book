# Entendendo interface no Golang de uma vez por todas

Nesse tópico, vamos abordar o uso de interfaces no Golang. Vou confessar que no início tive bastante dificuldade para entender interface no GO e vou tentar trazer de uma forma diferente que constumamos ver nos exemplos.

Primeiro ponto que precisamos entender, é que a interface está **se referenciando aos métodos das struct**. Deixa eu reescrever esse ponto, a interface é baseado numa ***assinatura***, ou seja, na assinatura de uma struct que possui N métodos. Todas as assinaturas que a interface contiver, é necessário que a struct contenha os mesmos métodos ou mais.

Vamos ao exemplo, que acredito que ficará mais claro. Nesse exemplo, vamos trabalhar para gravar os dados de uma struct de produto. Vamos iniciar criando a nossa struct.

Antes de começar, vamos ter 3 package no nosso exemplo, main, entity e storage.

1. Criando e acessando o diretório;
```
  mkdir basic_interface
  cd basic_interface
```
2. Vamos iniciar o nosso GO mod;
```
  go mod init github.com/tomelin/golang-book/examples/basic_interface
```
3. vamos criar o nosso package entity, já com as nossa struct
  
   3.1 Criando o diretório e o arquivo entity;
```
  mkdir entity
  cd entity
  touch entity.go
```

   3.2 criando a nossa struct
```
  package entity

  // Produto
  type Produto struct {
    ID string `json:"id,omitempty"`
    Name string `json:"name" binding:"required"`
    Descricao string `json:"description,omitempty"`
    Preco float32 `json:"preco" binding:"required"`
    Estoque uint `json:"estoque" binding:"required"`
  }

  // ProdutoInterface essa interface não precisa estar nesse package, pode estar em qualquer outro package.  Para fins didáticos, deixarei dentro de entity
  type ProdutoInterface interface {
    Create(*Produto) error
    FindAll() ([]Produto, error)
    FindById(id *string) (*Produto, error)
    Delete(id *string) error
  }
```

Nesse passo, a struct Produto e a interface ProdutoInterface, estão em letra maiuscula, pois são **publicas** e serão acessadas por outros packages.

**Perceba, criamos a interface e não conectamos em nenhum momento a struct.** 
Esse é o ponto importante a ser entendido, nesse momento a interface, não está vinculada a nenhuma struct, apenas tem o que chamamos assinatura, e essa assinatura possui o seguinte:
```
  Create(*Produto) error
  FindAll() ([]Produto, error)
  FindById(id *string) (*Produto,error)
  Delete(id *string) error
```

Essa assinatura nos informa que toda e qualquer struct que tiver os métodos acima, podemos usar a interface ProdutoInterface

Vamos agora, criar o package storage e dentro de storage. Teremos 3 structs para trabalharmos com (database, cache, files), para usarmos a interface que criamos  
As structs que criaremos precisarão ter a mesma assinatura (métodos).

Vamos entender conforme o nosso exemplo:

1. vamos criar o nosso package storage, com a struct para trabalhar com database
```
  mkdir storage
  cd storage
  touch database.go
```
2. DATABASE  
   2.1 Agora iremos criar uma struct para trabahar com database

```
  package storage
    
  import (
    "database/sql"
    "github.com/tomelin/golang-book/examples/basic_interface/entity"
    "log"
  )
  
  type ProdutoStorageDB struct {
    DB *sql.DB
  }
  
  // NewProdutoStorageDB receberemos a conexão do SQL e associando a struct. E estmos retornando a nossa struct ProdutoStorageDB, porém estamos informando que ela é do tipo entity.ProdutoInterface (ou possui as assinaturas) que a interface precisa para trabalhar
  func NewProdutoStorageDB(db *sql.DB) entity.ProdutoInterface {
    _return &ProdutoStorageDB{DB: db}
  }
```
   
    Nesse momento a struct ProdutoStorageDB, não possui a assinatura para trabalhar com a interface entity.ProdutoInterface, ou seja, nesse momento teremos um erro de assinatura.
  
    Agora vamos implementar a nossa assinatura para trabalharmos com a interface

   2.2 Implementando a assinatura a nossa struct ProdutoStorageDB, continuaremos com o nosso exemplo acima  

```
  package storage
  
  import (
  "database/sql"
  "github.com/tomelin/golang-book/examples/basic_interface/entity"
  "log"
  )
  
  type ProdutoStorageDB struct {
  DB *sql.DB
  }
  
  // NewProdutoStorageDB receberemos a conexão do SQL e associando a struct. E estmos retornando a nossa struct ProdutoStorageDB, porém estamos informando que ela é do tipo entity.ProdutoInterface (ou possui as assinaturas) que a interface precisa para trabalhar
  func NewProdutoStorageDB(db *sql.DB) entity.ProdutoInterface {
  return &ProdutoStorageDB{DB: db}
  }
  
  func (p *ProdutoStorageDB) Create(produto *entity.Produto) error {
  
    log.Println(produto)
    return nil
  }
  func (p *ProdutoStorageDB) FindAll() ([]entity.Produto, error)           { return nil, nil }
  func (p *ProdutoStorageDB) FindById(id *string) (*entity.Produto, error) { return nil, nil }
  func (p *ProdutoStorageDB) Delete(id *string) error               { return nil }
  func (p *ProdutoStorageDB) CloseDB() error                        { return nil }
```

   Aqui é onde realmente acontece a mágica, a struct ProdutoStorageDB, contém os métodos necessários para implementar a intefrcae entity.ProdutoInterface, ou seja, nesse caso já podemos usar a nossa interface.

3. Agora iremos criar uma struct para trabahar com cache redis
  3.1 Vamos criar o arquivo cache, dentro do nosso diretório storage

```
  cd storage
  touch cache.go
```

   3.2

 ```
  package storage

  import (
  "github.com/tomelin/golang-book/examples/basic_interface/entity"
  "log"
  )
  
  type ProdutoStorageCache struct {
  ConnectionString *string
  }
  
  // NewProdutoStorageCache receberemos a concection string para se conectar ao Redis e será associado a struct.  e estmos retornando a nossa struct ProdutoStorageCache, porém estamos informando que ela é do tipo entity.ProdutoInterface (ou possui as assinaturas) que a interface precisa para trabalhar
  func NewProdutoStorageCache(connection *string) entity.ProdutoInterface {
  return &ProdutoStorageCache{ConnectionString: connection}
  }
 ```

     Da mesma forma que o exemplo do item 2, nesse momento a struct ProdutoStorageCache não está preparada para trabalhar com a nossa interface entity.ProdutoInterface, simplesmente porque não possui os métodos necessários.
      
     Agora vamos implementar a nossa assinatura para trabalharmos com a interface.

      3.3

```
  package storage
  
  import (
    "github.com/tomelin/golang-book/examples/basic_interface/entity"
    "log"
  )
  
  type ProdutoStorageCache struct {
    ConnectionString *string
  }
  
  // NewProdutoStorageCache receberemos a concection string para se conectar ao Redis e será associado a struct.  e estmos retornando a nossa struct ProdutoStorageCache, porém estamos informando que ela é do tipo entity.ProdutoInterface (ou possui as assinaturas) que a interface precisa para trabalhar
  func NewProdutoStorageCache(connection *string) entity.ProdutoInterface {
    return &ProdutoStorageCache{ConnectionString: connection}
  }
  
  func (p *ProdutoStorageCache) Create(produto *entity.Produto) error {
    log.Println(produto)
    return nil
  }
  func (p *ProdutoStorageCache) FindAll() ([]entity.Produto, error)           { return nil, nil }
  func (p *ProdutoStorageCache) FindById(id *string) (*entity.Produto, error) { return nil, nil }
  func (p *ProdutoStorageCache) Delete(id *string) error               { return nil }
  func (p *ProdutoStorageCache) CloseCache() error                     { return nil } 
```

      Observe que agora temos praticamente o mesmo exemplo que ProdutoStorageDB, o que muda é a struct no qual associamos os métodos.

      **O que realmente importa entender, é que tanto a struct ProdutoStorageDB e ProdutoStorageCache possuem os mesmos métodos para a assinatura da interface**

4. Agora iremos criar uma struct para trabahar com arquivos
   4.1 Criaremos o arquivo files, dentro do nosso diretório storage

  ```
    cd storage
    touch files.go
  ```
      4.2
  ```
    package storage
    
    import (
      "github.com/tomelin/golang-book/examples/basic_interface/entity"
      "log"
    )
    
    type ProdutoStorageFiles struct {
      FileName *string
      FilePath *string
    }
    
    // Aqui estamos recebendo a conexão do SQL e associando a struct.  e estmos retornando a nossa struct ProdutoStorageCache, porém estamos informando que ela é do tipo entity.ProdutoInterface (ou possui as assinaturas) que a interface precisa para trabalhar
    func NewProdutoStorageFile(name, dir string) entity.ProdutoInterface {
      return &ProdutoStorageFiles{
        FileName: &name,
        FilePath: &dir,
      }
    }
  ```
Da mesma forma que o exemplo do item 2, nesse momento a struct ProdutoStorageFiles não está preparada para trabalhar com a nossa interface entity.ProdutoInterface, simplesmente porque não possui os métodos necessários.

Agora vamos implementar a nossa assinatura para trabalharmos com a interface.
4.3
```
  package storage
  
  import (
    "github.com/tomelin/golang-book/examples/basic_interface/entity"
    "log"
  )
  
  type ProdutoStorageFiles struct {
    FileName *string
    FilePath *string
  }
  
  // Aqui estamos recebendo a conexão do SQL e associando a struct.  e estmos retornando a nossa struct ProdutoStorageCache, porém estamos informando que ela é do tipo entity.ProdutoInterface (ou possui as assinaturas) que a interface precisa para trabalhar
  func NewProdutoStorageFile(name, dir string) entity.ProdutoInterface {
    return &ProdutoStorageFiles{
      FileName: &name,
      FilePath: &dir,
    }
  }
  
  func (p *ProdutoStorageFiles) Create(produto *entity.Produto) error {
  
    log.Println(produto)
    return nil
  }
  func (p *ProdutoStorageFiles) FindAll() ([]entity.Produto, error)           { return nil, nil }
  func (p *ProdutoStorageFiles) FindById(id *string) (*entity.Produto, error) { return nil, nil }
  func (p *ProdutoStorageFiles) Delete(id *string) error               { return nil }
  func (p *ProdutoStorageFiles) RemoveFile(f string) error             { return nil }
```
Observe que agora temos praticamente o mesmo exemplo que ProdutoStorageDB e ProdutoStorageCache, o que muda é a struct no qual associamos os métodos.

**Mais uma vez, o que importa entender, é que tanto a struct ProdutoStorageDB, ProdutoStorageCache e ProdutoStorageFiles possuem os mesmos métodos que a interface entity.ProdutoInterface está esperado.**

Aqui nesse exemplo, não iremos desenvolver a lógica dentro da função, o nosso propósito é auxiliar no entendimento de interfaces no Golang.

5. Agora, vamos implementar a lógica no arquivo mais, de como chamariamos essas funções.

5.1 Vamos criar o diretório cmd na raiz do nosso projeto
```
  mkdir cmd
  cd cmd
  touch main.go
```

5.2 Agora vamos implemntar a nossa lógica de como consumir nossos storages.
```
  package main
  
  import (
  "database/sql"
  "fmt"
  "log"
  
    _ "github.com/mattn/go-sqlite3"
    "github.com/tomelin/golang-book/examples/basic_interface/entity"
    "github.com/tomelin/golang-book/examples/basic_interface/storage"
  )
  
  func main() {
  
    produto := &entity.Produto{
      Name:      "notebook",
      Descricao: "notebook para games",
      Preco:     5000.35,
      Estoque:   1,
    }
  
    dbOpen, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
      log.Fatal(err)
    }
  
    // Database
    produtoInterface := storage.NewProdutoStorageDB(dbOpen)
    err = produtoInterface.Create(produto)
    if err != nil {
      log.Fatal(err)
    }
  
    // Cache Redis
    var cacheuser, cachePass string
    cacheConnectionString := fmt.Sprintf("redis://%s:%s@localhost:6379/0", cacheuser, cachePass)
    produtoInterface = storage.NewProdutoStorageCache(&cacheConnectionString)
    err = produtoInterface.Create(produto)
    if err != nil {
      log.Fatal(err)
    }
  
    // Files
    produtoInterface = storage.NewProdutoStorageFile("example.txt", "/tmp/golang")
    err = produtoInterface.Create(produto)
    if err != nil {
      log.Fatal(err)
    }
  }

```

Observer que a veriável **ProdutoInterface** é reutilizada em outras chamdas do storage (não é uma boa pratica fazer isso), pois as outras funções retornando o produtoInterface também.

Mas novamente, as structs (ProdutoStorageDB, ProdutoStorageCache e ProdutoStorageFiles), implementam os métodos que a assinatura do ProdutoInterface precisa para fazer o match e passar a funcionar.

Também precebemos que a struct pode ter outros métodos além da ProdutoInterface, não pode ter menos, pois se tiver menos, a assinatura não será válida.
 