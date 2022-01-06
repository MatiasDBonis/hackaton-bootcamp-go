package internal

import (
	"errors"
	"log"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/db"
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

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Insert(invoice Invoices) (Invoices, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO invoices (id, datetime, idcustomer, total) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(invoice.Id, invoice.Datetime, invoice.IdCustomer, invoice.Total)
	if err != nil {
		return Invoices{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Invoices{}, errors.New("no se insertó el Invoice")
	}

	return invoice, nil
}
