package database

import (
	"database/sql"
	"fmt"

	"api-produto/config"
)

var (
	db *sql.DB
)

type DatabaseInterface interface {
	GetDB() (DB *sql.DB)
	Close() error
}

type dabase_pool struct {
	DB *sql.DB
}

var dbpool = &dabase_pool{}

func NewDB(conf *config.Config) *dabase_pool {

	// db, err := sql.Open("mysql", "root:<yourMySQLdatabasepassword>@tcp(127.0.0.1:3306)/test")

	conf.DBConfig.DB_DSN = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.DBConfig.DB_USER, conf.DBConfig.DB_PASS, conf.DBConfig.DB_HOST, conf.DBConfig.DB_PORT, conf.DBConfig.DB_NAME)
	dbpool = Mysql(conf)

	return dbpool
}

func (d *dabase_pool) Close() error {

	err := d.DB.Close()
	if err != nil {
		return err
	}

	dbpool = &dabase_pool{}

	return err
}

func (d *dabase_pool) GetDB() (DB *sql.DB) {
	return d.DB
}
