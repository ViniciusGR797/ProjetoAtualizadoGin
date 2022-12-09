package service

import (
	"api-produto/database"
	"api-produto/entity"
)

type ProdutoServiceInterface interface {
	GetProduto(ID *int64) *entity.Produto
	GetAll() *entity.ListaDeProduto
	Create(produto *entity.Produto) int64
	Update(ID *int64, produto *entity.Produto) int64
	Delete(id *int64) int64
}

type produto_service struct {
	dbp database.DatabaseInterface
}

func NewProdutoService(dabase_pool database.DatabaseInterface) *produto_service {
	return &produto_service{
		dabase_pool,
	}
}
