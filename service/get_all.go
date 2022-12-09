package service

import (
	"api-produto/entity"
	"fmt"
)

func (ps *produto_service) GetAll() *entity.ListaDeProduto {
	database := ps.dbp.GetDB()

	rows, err := database.Query("SELECT id, nome, codigo, valor FROM produto LIMIT 100")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	lista_produtos := &entity.ListaDeProduto{}

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
