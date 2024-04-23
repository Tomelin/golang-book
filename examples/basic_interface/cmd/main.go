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
