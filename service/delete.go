package service

import (
	"log"
)

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
