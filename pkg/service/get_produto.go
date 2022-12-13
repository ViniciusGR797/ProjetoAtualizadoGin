package service

import (
	"log"
	"product/pkg/entity"
)

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
