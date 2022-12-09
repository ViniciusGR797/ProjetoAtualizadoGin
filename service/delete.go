package service

import (
	"log"
)

func (ps *produto_service) Delete(id *int64) int64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("DELETE FROM produto WHERE id = ?")
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

	return aff

}
