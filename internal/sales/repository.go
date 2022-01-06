package internal

import (
	"errors"
	"log"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/db"
)

type Sales struct {
	Id        int     `csv:"id"`
	IdProduct int     `csv:"id_product"`
	Price     float32 `csv:"price"`
}

type Repository interface {
	Insert(customer Sales) (Sales, error)
	Update(customer Sales) (Sales, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Insert(sale Sales) (Sales, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO sales (id, id_product, price) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(sale.Id, sale.IdProduct, sale.Price)
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
	db := db.StorageDB

	stmt, err := db.Prepare("UPDATE sales SET id_product = ?, price = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(sale.IdProduct, sale.Price, sale.Id)
	if err != nil {
		return Sales{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Sales{}, errors.New("no se insertó la sale")
	}

	return sale, nil
}
