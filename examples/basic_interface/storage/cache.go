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
	log.Println("cache: the produto é: ", produto)
	return nil
}
func (p *ProdutoStorageCache) FindAll() ([]entity.Produto, error)           { return nil, nil }
func (p *ProdutoStorageCache) FindById(id *string) (*entity.Produto, error) { return nil, nil }
func (p *ProdutoStorageCache) Delete(id *string) error                      { return nil }
func (p *ProdutoStorageCache) CloseCache() error                            { return nil }
