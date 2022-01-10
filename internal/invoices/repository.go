package invoices

import (
	"database/sql"
	"errors"
	"log"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
)

type Repository interface {
	Insert(invoices domain.Invoices) (domain.Invoices, error)
	Update(invoice domain.Invoices) (domain.Invoices, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(invoice domain.Invoices) (domain.Invoices, error) {
	stmt, err := r.db.Prepare("INSERT INTO invoices (id, datetime, idcustomer, total) VALUES(?, ?, ?, (SELECT SUM(S.quantity) FROM sales S WHERE S.idinvoice = ?))")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(invoice.Id, invoice.Datetime, invoice.IdCustomer, invoice.Id)
	if err != nil {
		return domain.Invoices{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.Invoices{}, errors.New("no se insertó el invoice")
	}

	return invoice, nil
}

func (r *repository) Update(invoice domain.Invoices) (domain.Invoices, error) {
	stmt, err := r.db.Prepare("UPDATE invoices SET idcustomer = ?, total = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(invoice.Datetime, invoice.IdCustomer, invoice.Total, invoice.Id)
	if err != nil {
		return domain.Invoices{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.Invoices{}, errors.New("no se actualizó el invoice")
	}

	return invoice, nil
}
