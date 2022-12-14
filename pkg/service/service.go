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

// Função que retorna lista de logs
func (ps *Produto_service) GetLog() *entity.LogList {
	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query("SELECT * FROM log")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo LogList (vazia)
	lista_logs := &entity.LogList{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo Log (vazia)
		log := entity.Log{}

		// pega dados da query e atribui a variável log, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&log.ID, &log.Method, &log.Description, &log.Data); err != nil {
			fmt.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			lista_logs.List = append(lista_logs.List, &log)
		}

	}

	// retorna lista de logs
	return lista_logs
}

// Função que cria produto e retorna id criado
func (ps *Produto_service) Create(produto *entity.Produto) int {
	// pega database
	database := ps.dbp.GetDB()

	// prepara query para ser executada no database
	stmt, err := database.Prepare("INSERT INTO product (pro_name, pro_code, pro_price) VALUES (?, ?, ?)")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// prepara query para ser executada no database
	stmt_log, err := database.Prepare("INSERT INTO log (log_method, log_description) VALUES (?, ?)")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer stmt.Close()
	// fecha linha da query, quando sair da função
	defer stmt_log.Close()

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	result, err := stmt.Exec(produto.Name, produto.Code, produto.Price)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// pega id do último product inserido
	lastId, err := result.LastInsertId()
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// preenche descrição do log
	description := "user '" + "grupob" + "' inserted product '" + strconv.Itoa(int(lastId)) + "' from database"

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	_, err = stmt_log.Exec("POST", description)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// retorna último id (converte int64 para int)
	return int(lastId)
}

// Função que deleta produto
func (ps *Produto_service) Delete(id *int) int {
	// pega database
	database := ps.dbp.GetDB()

	// prepara query para ser executada no database
	stmt, err := database.Prepare("DELETE FROM product WHERE pro_id = ?")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// prepara query para ser executada no database
	stmt_log, err := database.Prepare("INSERT INTO log (log_method, log_description) VALUES (?, ?)")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer stmt.Close()
	// fecha linha da query, quando sair da função
	defer stmt_log.Close()

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	result, err := stmt.Exec(id)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// RowsAffected retorna número de linhas afetadas com delete
	rowsaff, err := result.RowsAffected()
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// preenche descrição do log
	description := "user '" + "grupob" + "' deleted product '" + strconv.Itoa(*id) + "' from database"

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	_, err = stmt_log.Exec("DELETE", description)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// retorna rowsaff (converte int64 para int)
	return int(rowsaff)
}

// Função que retorna lista de products
func (ps *Produto_service) GetAll() *entity.ProdutoList {
	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query("SELECT * FROM product")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	// prepara query para ser executada no database
	stmt_log, err := database.Prepare("INSERT INTO log (log_method, log_description) VALUES (?, ?)")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()
	// fecha linha da query, quando sair da função
	defer stmt_log.Close()

	// variável do tipo ProductList (vazia)
	lista_produtos := &entity.ProdutoList{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo Produto (vazia)
		produto := entity.Produto{}

		// pega dados da query e atribui a variável produto, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&produto.ID, &produto.Name, &produto.Code, &produto.Price, &produto.CreatedAt, &produto.UpdatedAt); err != nil {
			fmt.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			lista_produtos.List = append(lista_produtos.List, &produto)
		}

	}

	// preenche descrição do log
	description := "user '" + "grupob" + "' gets all products from database"

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	_, err = stmt_log.Exec("GET", description)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// retorna lista de produtos
	return lista_produtos
}

// Função que retorna product
func (ps *Produto_service) GetProduto(ID *int) *entity.Produto {
	// pega database
	database := ps.dbp.GetDB()

	// prepara query para ser executada no database
	stmt, err := database.Prepare("SELECT * FROM product WHERE pro_id = ?")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// prepara query para ser executada no database
	stmt_log, err := database.Prepare("INSERT INTO log (log_method, log_description) VALUES (?, ?)")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer stmt.Close()
	// fecha linha da query, quando sair da função
	defer stmt_log.Close()

	// variável do tipo Produto (vazia)
	produto := entity.Produto{}

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	err = stmt.QueryRow(ID).Scan(&produto.ID, &produto.Name, &produto.Code, &produto.Price, &produto.CreatedAt, &produto.UpdatedAt)
	// verifica se teve erro
	if err != nil {
		log.Println("error: cannot find produto", err.Error())
	}

	// preenche descrição do log
	description := "user '" + "grupob" + "' gets product '" + strconv.Itoa(*ID) + "' from database"

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	_, err = stmt_log.Exec("GET", description)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// retorna produto
	return &produto
}

// Função que altera produto
func (ps *Produto_service) Update(ID *int, produto *entity.Produto) int {
	// pega database
	database := ps.dbp.GetDB()

	// prepara query para ser executada no database
	stmt, err := database.Prepare("UPDATE product SET pro_name = ?, pro_code = ?, pro_price = ? WHERE pro_id = ?")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// prepara query para ser executada no database
	stmt_log, err := database.Prepare("INSERT INTO log (log_method, log_description) VALUES (?, ?)")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer stmt.Close()
	// fecha linha da query, quando sair da função
	defer stmt_log.Close()

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	result, err := stmt.Exec(produto.Name, produto.Code, produto.Price, ID)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// RowsAffected retorna número de linhas afetadas com update
	rowsaff, err := result.RowsAffected()
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// preenche descrição do log
	description := "user '" + "grupob" + "' updated product '" + strconv.Itoa(*ID) + "' from database"

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	_, err = stmt_log.Exec("PUT", description)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// retorna rowsaff (converte int64 para int)
	return int(rowsaff)
}
