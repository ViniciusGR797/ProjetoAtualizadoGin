package service

import (
	// Import interno de packages do próprio sistema
	"product/database"
	"product/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Produto (tudo que tiver os métodos abaixo do CRUD são serviços de produto)
type ProdutoServiceInterface interface {
	// Pega produto em específico passando o id dele como parâmetro
	GetProduto(ID *int64) *entity.Produto
	// Pega todos os produtos, logo lista todos os produtos
	GetAll() *entity.ProdutoList
	// Cria um novo produto passando seus dados como parâmetro
	Create(produto *entity.Produto) int64
	// Atualiza dados de um produto, passando id do produto e dados a serem alterados por parâmetro
	Update(ID *int64, produto *entity.Produto) int64
	// Deletar produto passando id por parâmetro
	Delete(id *int64) int64
}

// Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type produto_service struct {
	dbp database.DatabaseInterface
}

// Cria novo serviço de CRUD para pool de conexão
func NewProdutoService(dabase_pool database.DatabaseInterface) *produto_service {
	return &produto_service{
		dabase_pool,
	}
}
