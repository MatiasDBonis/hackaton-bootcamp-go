package internal

import (
	"errors"
	"log"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/db"
)

type Customers struct {
	id         int
	last_name  string
	first_name string
	condition  string
}

type Repository interface {
	Insert(customer Customers) (Customers, error)
	Update(customer Customers) (Customers, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Insert(customer Customers) (Customers, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO customers (id, lastname, firstname, condition) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(customer.id, customer.last_name, customer.first_name, customer.condition)
	if err != nil {
		return Customers{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Customers{}, errors.New("no se insertó el customer")
	}

	return customer, nil
}

func (r *repository) Update(customer Customers) (Customers, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("UPDATE customers SET lastname = ?, firstname = ?, condition = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(customer.last_name, customer.first_name, customer.condition, customer.id)
	if err != nil {
		return Customers{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Customers{}, errors.New("no se insertó el customer")
	}

	return customer, nil
}
