package service

import (
	"log"
	"product/entity"
)

func (ps *produto_service) GetProduto(ID *int64) *entity.Produto {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT id, nome, codigo, valor FROM produto WHERE id = ?")
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
