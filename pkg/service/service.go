package service

import (
	"fmt"
	"log"
	"strconv"

	// Import interno de packages do próprio sistema
	"product/pkg/database"
	"product/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Produto (tudo que tiver os métodos abaixo do CRUD são serviços de produto)
type ProdutoServiceInterface interface {
	// Pega todos os logs, logo lista todos os logs
	GetLog() *entity.LogList
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
type Produto_service struct {
	dbp database.DatabaseInterface
}

// Cria novo serviço de CRUD para pool de conexão
func NewProdutoService(dabase_pool database.DatabaseInterface) *Produto_service {
	return &Produto_service{
		dabase_pool,
	}
}

func (ps *Produto_service) GetLog() *entity.LogList {
	database := ps.dbp.GetDB()

	rows, err := database.Query("SELECT * FROM log")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	lista_logs := &entity.LogList{}

	for rows.Next() {
		log := entity.Log{}

		if err := rows.Scan(&log.ID, &log.Method, &log.Description, &log.Data); err != nil {
			fmt.Println(err.Error())
		} else {
			lista_logs.List = append(lista_logs.List, &log)
		}

	}

	return lista_logs
}

func (ps *Produto_service) Create(produto *entity.Produto) int {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO product (pro_name, pro_code, pro_price) VALUES (?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	stmt_log, err := database.Prepare("INSERT INTO log (log_method, log_description) VALUES (?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()
	defer stmt_log.Close()

	result, err := stmt.Exec(produto.Name, produto.Code, produto.Price)
	if err != nil {
		log.Println(err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
	}

	description := "user '" + "admin" + "' inserted product '" + strconv.Itoa(int(lastId)) + "' from database"

	_, err = stmt_log.Exec("POST", description)
	if err != nil {
		log.Println(err.Error())
	}

	return int(lastId)
}

func (ps *Produto_service) Delete(id *int) int {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("DELETE FROM product WHERE pro_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt_log, err := database.Prepare("INSERT INTO log (log_method, log_description) VALUES (?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()
	defer stmt_log.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Println(err.Error())
	}

	aff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	description := "user '" + "admin" + "' deleted product '" + strconv.Itoa(*id) + "' from database"

	_, err = stmt_log.Exec("DELETE", description)
	if err != nil {
		log.Println(err.Error())
	}

	return int(aff)
}

func (ps *Produto_service) GetAll() *entity.ProdutoList {
	database := ps.dbp.GetDB()

	rows, err := database.Query("SELECT * FROM product")
	if err != nil {
		fmt.Println(err.Error())
	}

	stmt_log, err := database.Prepare("INSERT INTO log (log_method, log_description) VALUES (?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer rows.Close()
	defer stmt_log.Close()

	lista_produtos := &entity.ProdutoList{}

	for rows.Next() {
		produto := entity.Produto{}

		if err := rows.Scan(&produto.ID, &produto.Name, &produto.Code, &produto.Price, &produto.CreatedAt, &produto.UpdatedAt); err != nil {
			fmt.Println(err.Error())
		} else {
			lista_produtos.List = append(lista_produtos.List, &produto)
		}

	}

	description := "user '" + "admin" + "' gets all products from database"

	_, err = stmt_log.Exec("GET", description)
	if err != nil {
		log.Println(err.Error())
	}

	return lista_produtos
}

func (ps *Produto_service) GetProduto(ID *int) *entity.Produto {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT * FROM product WHERE pro_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt_log, err := database.Prepare("INSERT INTO log (log_method, log_description) VALUES (?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()
	defer stmt_log.Close()

	produto := entity.Produto{}

	err = stmt.QueryRow(ID).Scan(&produto.ID, &produto.Name, &produto.Code, &produto.Price, &produto.CreatedAt, &produto.UpdatedAt)
	if err != nil {
		log.Println("error: cannot find produto", err.Error())
	}

	description := "user '" + "admin" + "' gets product '" + strconv.Itoa(*ID) + "' from database"

	_, err = stmt_log.Exec("GET", description)
	if err != nil {
		log.Println(err.Error())
	}

	return &produto
}

func (ps *Produto_service) Update(ID *int, produto *entity.Produto) int {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE product SET pro_name = ?, pro_code = ?, pro_price = ? WHERE pro_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt_log, err := database.Prepare("INSERT INTO log (log_method, log_description) VALUES (?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()
	defer stmt_log.Close()

	result, err := stmt.Exec(produto.Name, produto.Code, produto.Price, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	description := "user '" + "admin" + "' updated product '" + strconv.Itoa(*ID) + "' from database"

	_, err = stmt_log.Exec("PUT", description)
	if err != nil {
		log.Println(err.Error())
	}

	return int(rowsaff)
}
