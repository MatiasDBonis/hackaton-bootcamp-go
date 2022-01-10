package products

import (
	"database/sql"
	"errors"
	"log"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
)

type Repository interface {
	Insert(customer domain.Products) (domain.Products, error)
	Update(customer domain.Products) (domain.Products, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(product domain.Products) (domain.Products, error) {
	stmt, err := r.db.Prepare("INSERT INTO products (id, description, price) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Id, product.Description, product.Price)
	if err != nil {
		return domain.Products{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.Products{}, errors.New("no se insertó el product")
	}

	return product, nil
}

func (r *repository) Update(product domain.Products) (domain.Products, error) {
	stmt, err := r.db.Prepare("UPDATE products SET description = ?, price = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Description, product.Price, product.Id)
	if err != nil {
		return domain.Products{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.Products{}, errors.New("no se actualizó el product")
	}

	return product, nil
}
