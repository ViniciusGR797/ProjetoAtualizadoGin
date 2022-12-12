package service

import (
	"log"
	"product/entity"
)

func (ps *produto_service) Update(ID *int, produto *entity.Produto) int {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE produto SET nome = ?, codigo = ?, valor = ? WHERE id = ?")
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
