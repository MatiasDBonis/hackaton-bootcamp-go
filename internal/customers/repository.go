package internal

import (
	"errors"
	"log"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/db"
)

type Customers struct {
	Id        int    `csv:"id"`
	LastName  string `csv:"last_name"`
	FirstName string `csv:"first_name"`
	Condition string `csv:"condition"`
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

	result, err := stmt.Exec(customer.Id, customer.LastName, customer.FirstName, customer.Condition)
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

	result, err := stmt.Exec(customer.LastName, customer.FirstName, customer.Condition, customer.Id)
	if err != nil {
		return Customers{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Customers{}, errors.New("no se insertó el customer")
	}

	return customer, nil
}
