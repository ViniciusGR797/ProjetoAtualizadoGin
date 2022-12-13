package service

import (
	"log"
	"product/pkg/entity"
)

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
