package service

import (
	"fmt"
	"log"

	// Import interno de packages do próprio sistema
	"product/pkg/database"
	"product/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Produto (tudo que tiver os métodos abaixo do CRUD são serviços de produto)
type ProdutoServiceInterface interface {
	// Pega produto em específico passando o id dele como parâmetro
	GetProduto(ID *int) *entity.Produto
	// Pega todos os produtos, logo lista todos os produtos
	GetAll() *entity.ProdutoList
	// Cria um novo produto passando seus dados como parâmetro
	Create(produto *entity.Produto) int
	// Atualiza dados de um produto, passando id do produto e dados a serem alterados por parâmetro
	Update(ID *int, produto *entity.Produto) int
	// Deletar produto passando id por parâmetro
	Delete(id *int) int
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

func (ps *produto_service) Create(produto *entity.Produto) int {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO product (pro_name, pro_code, pro_price) VALUES (?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(produto.Name, produto.Code, produto.Price)
	if err != nil {
		log.Println(err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
	}

	return int(lastId)
}

func (ps *produto_service) Delete(id *int) int {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("DELETE FROM product WHERE pro_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Println(err.Error())
	}

	aff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return int(aff)
}

func (ps *produto_service) GetAll() *entity.ProdutoList {
	database := ps.dbp.GetDB()

	rows, err := database.Query("SELECT pro_id, pro_name, pro_code, pro_price FROM product LIMIT 100")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	lista_produtos := &entity.ProdutoList{}

	for rows.Next() {
		produto := entity.Produto{}

		if err := rows.Scan(&produto.ID, &produto.Name, &produto.Code, &produto.Price); err != nil {
			fmt.Println(err.Error())
		} else {
			lista_produtos.List = append(lista_produtos.List, &produto)
		}

	}

	return lista_produtos
}

func (ps *produto_service) GetProduto(ID *int) *entity.Produto {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT pro_id, pro_name, pro_code, pro_price FROM product WHERE pro_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	produto := entity.Produto{}

	err = stmt.QueryRow(ID).Scan(&produto.ID, &produto.Name, &produto.Code, &produto.Price)
	if err != nil {
		log.Println("error: cannot find produto", err.Error())
	}

	return &produto
}

func (ps *produto_service) Update(ID *int, produto *entity.Produto) int {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE product SET pro_name = ?, pro_code = ?, pro_price = ? WHERE pro_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(produto.Name, produto.Code, produto.Price, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return int(rowsaff)
}
