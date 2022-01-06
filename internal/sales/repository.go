package internal

import (
	"database/sql"
	"errors"
	"log"
)

type Sales struct {
	Id        int     `csv:"id"`
	IdInvoice int     `csv:"id_invoice"`
	IdProduct int     `csv:"id_product"`
	Quantity  float32 `csv:"quantity"`
}

type Repository interface {
	Insert(customer Sales) (Sales, error)
	Update(customer Sales) (Sales, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(sale Sales) (Sales, error) {
	stmt, err := r.db.Prepare("INSERT INTO sales (id, idinvoice, idproduct, quantity) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(sale.Id, sale.IdInvoice, sale.IdProduct, sale.Quantity)
	if err != nil {
		return Sales{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Sales{}, errors.New("no se insertó la sale")
	}

	return sale, nil
}

func (r *repository) Update(sale Sales) (Sales, error) {
	stmt, err := r.db.Prepare("UPDATE sales SET idinvoice = ?, idproduct = ?, quantity = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(sale.IdInvoice, sale.IdProduct, sale.Quantity, sale.Id)
	if err != nil {
		return Sales{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Sales{}, errors.New("no se insertó la sale")
	}

	return sale, nil
}
