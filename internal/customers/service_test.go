package customers

import (
	"fmt"
	"testing"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/db"
	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestInsertOk(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	parsedCustomers, err := parser.ParseDataCustomers()
	assert.Nil(t, err)
	fmt.Println(parsedCustomers)

	cant, err := service.InsertAll(parsedCustomers)
	assert.Nil(t, err)
	assert.Equal(t, 50, cant)
}

func TestInsertCustomersExistents(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	parsedCustomers, err := parser.ParseDataCustomers()
	assert.Nil(t, err)

	cant, err := service.InsertAll(parsedCustomers)
	assert.Equal(t, 50, cant)
	assert.Nil(t, err)

	cant, err = service.InsertAll(parsedCustomers)
	assert.Equal(t, 0, cant)
	assert.NotNil(t, err)
}

func TestUpdateOK(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	newCustomer := []domain.Customers{
		{
			Id:        1,
			LastName:  "asd",
			FirstName: "asd",
			Condition: "asd",
		},
	}

	customerEdited := domain.Customers{
		Id:        1,
		LastName:  "mati",
		FirstName: "facu",
		Condition: "hola",
	}

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	service.InsertAll(newCustomer)

	customerResult, err := service.Update(customerEdited)

	assert.Nil(t, err)
	assert.Equal(t, customerEdited, customerResult)
}

func TestUpdateIdInexistent(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	customerEdited := domain.Customers{
		Id:        2,
		LastName:  "mati",
		FirstName: "facu",
		Condition: "hola",
	}

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	customerResult, err := service.Update(customerEdited)

	assert.Error(t, err)
	assert.NotEqual(t, customerEdited, customerResult)
}
