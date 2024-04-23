package entity

// Produto
type Produto struct {
	ID        string  `json:"id,omitempty"`
	Name      string  `json:"name" binding:"required"`
	Descricao string  `json:"description,omitempty"`
	Preco     float32 `json:"preco" binding:"required"`
	Estoque   uint    `json:"estoque" binding:"required"`
}

// ProdutoInterface essa interface não precisa estar nesse package, pode estar em qualquer outro package.  Para fins didáticos, deixarei dentro de entity
type ProdutoInterface interface {
	Create(*Produto) error
	FindAll() ([]Produto, error)
	FindById(id *string) (*Produto, error)
	Delete(id *string) error
}
