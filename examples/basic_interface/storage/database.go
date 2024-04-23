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

	log.Println("database: the produto é: ", produto)
	return nil
}
func (p *ProdutoStorageDB) FindAll() ([]entity.Produto, error)           { return nil, nil }
func (p *ProdutoStorageDB) FindById(id *string) (*entity.Produto, error) { return nil, nil }
func (p *ProdutoStorageDB) Delete(id *string) error                      { return nil }
func (p *ProdutoStorageDB) CloseDB() error                               { return nil }
