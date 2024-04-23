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

	log.Println("files: the produto é: ", produto)
	return nil
}
func (p *ProdutoStorageFiles) FindAll() ([]entity.Produto, error)           { return nil, nil }
func (p *ProdutoStorageFiles) FindById(id *string) (*entity.Produto, error) { return nil, nil }
func (p *ProdutoStorageFiles) Delete(id *string) error                      { return nil }
func (p *ProdutoStorageFiles) RemoveFile(f string) error                    { return nil }
