package sales

import (
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

	parsedSales, err := parser.ParseDataSales()
	assert.Nil(t, err)
	cant, err := service.InsertAll(parsedSales)
	assert.Nil(t, err)
	assert.Equal(t, 1000, cant)
}

func TestInsertCustomersExistents(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	parsedSales, err := parser.ParseDataSales()
	assert.Nil(t, err)

	cant, err := service.InsertAll(parsedSales)
	assert.Equal(t, 1000, cant)
	assert.Nil(t, err)

	cant, err = service.InsertAll(parsedSales)
	assert.Equal(t, 0, cant)
	assert.NotNil(t, err)
}

func TestUpdateOK(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	newSales := []domain.Sales{
		{
			Id:        1,
			IdInvoice: 1,
			IdProduct: 1,
			Quantity:  10,
		},
	}

	salesEdited := domain.Sales{
		Id:        1,
		IdInvoice: 2,
		IdProduct: 2,
		Quantity:  20,
	}

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	service.InsertAll(newSales)

	salesResult, err := service.Update(salesEdited)

	assert.Nil(t, err)
	assert.Equal(t, salesEdited, salesResult)
}

func TestUpdateIdInexistent(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	salesEdited := domain.Sales{
		Id:        2,
		IdInvoice: 1,
		IdProduct: 1,
		Quantity:  10,
	}

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	salesResult, err := service.Update(salesEdited)

	assert.Error(t, err)
	assert.NotEqual(t, salesEdited, salesResult)
}
