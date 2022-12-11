package service

import (
	"fmt"
	"product/entity"
)

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
