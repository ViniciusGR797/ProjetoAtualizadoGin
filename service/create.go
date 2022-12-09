package service

import (
	"log"
	"product/entity"
)

func (ps *produto_service) Create(produto *entity.Produto) int64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO produto (nome, codigo, valor) VALUES (?, ?, ?)")
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

	return lastId
}
