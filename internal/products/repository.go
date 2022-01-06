package internal

import (
	"errors"
	"log"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/db"
)

type Products struct {
	Id          int     `csv:"id"`
	Description string  `csv:"description"`
	Price       float32 `csv:"price"`
}

type Repository interface {
	Insert(customer Products) (Products, error)
	Update(customer Products) (Products, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Insert(product Products) (Products, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO products (id, description, price) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Id, product.Description, product.Price)
	if err != nil {
		return Products{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Products{}, errors.New("no se insertó el product")
	}

	return product, nil
}

func (r *repository) Update(product Products) (Products, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("UPDATE products SET description = ?, price = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Description, product.Price, product.Id)
	if err != nil {
		return Products{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Products{}, errors.New("no se insertó el product")
	}

	return product, nil
}
