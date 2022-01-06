package internal

import (
	"database/sql"
	"errors"
	"log"
)

type Invoices struct {
	Id         int     `csv:"id"`
	Datetime   string  `csv:"datetime"`
	IdCustomer string  `csv:"id_customer"`
	Total      float64 `csv:"total"`
}

type Repository interface {
	Insert(invoices Invoices) (Invoices, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(invoice Invoices) (Invoices, error) {
	stmt, err := r.db.Prepare("INSERT INTO invoices (id, datetime, idcustomer, total) VALUES(?, ?, ?, (SELECT SUM(S.quantity) FROM sales S WHERE S.idinvoice = ?))")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(invoice.Id, invoice.Datetime, invoice.IdCustomer, invoice.Id)
	if err != nil {
		return Invoices{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Invoices{}, errors.New("no se insertó el Invoice")
	}

	return invoice, nil
}
