package internal

import (
	"database/sql"
	"errors"
	"log"
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

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(customer Customers) (Customers, error) {
	stmt, err := r.db.Prepare("INSERT INTO customers(id, lastname, firstname, condicion) VALUES(?, ?, ?, ?)")
	// db.Prepare("INSERT INTO personas(nombre, apellido, edad) VALUES( ?, ?, ? )")
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
	stmt, err := r.db.Prepare("UPDATE customers SET lastname = ?, firstname = ?, condition = ? WHERE id = ?")
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
