package db

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
)

func init() {
	dataSource := "root:@tcp(localhost:3306)/hackaton_db"
	txdb.Register("txdb", "mysql", dataSource)
}

func InitTxdb() (*sql.DB, error) {
	db, err := sql.Open("txdb", "indentificar")
	if err != nil {
		return nil, err
	}
	return db, nil
}
