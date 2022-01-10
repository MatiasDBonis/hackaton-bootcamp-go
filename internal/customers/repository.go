package customers

import (
	"database/sql"
	"errors"
	"log"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
)

type Repository interface {
	Insert(customer domain.Customers) (domain.Customers, error)
	Update(customer domain.Customers) (domain.Customers, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(customer domain.Customers) (domain.Customers, error) {
	stmt, err := r.db.Prepare("INSERT INTO customers(id, lastname, firstname, condicion) VALUES(?, ?, ?, ?)")
	// db.Prepare("INSERT INTO personas(nombre, apellido, edad) VALUES( ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(customer.Id, customer.LastName, customer.FirstName, customer.Condition)
	if err != nil {
		return domain.Customers{}, err
	}

	return customer, nil
}

func (r *repository) Update(customer domain.Customers) (domain.Customers, error) {
	stmt, err := r.db.Prepare("UPDATE customers SET lastname = ?, firstname = ?, condicion = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(customer.LastName, customer.FirstName, customer.Condition, customer.Id)
	if err != nil {
		return domain.Customers{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.Customers{}, errors.New("no se actualizó el customer")
	}

	return customer, nil
}
