package invoices

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

	parsedInvoices, err := parser.ParseDataInvoices()
	assert.Nil(t, err)

	cant, err := service.InsertAll(parsedInvoices)
	assert.Nil(t, err)
	assert.Equal(t, 100, cant)
}

func TestInsertExistent(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	parsedInvoices, err := parser.ParseDataInvoices()
	assert.Nil(t, err)

	cant, err := service.InsertAll(parsedInvoices)
	assert.Equal(t, 100, cant)
	assert.Nil(t, err)

	cant, err = service.InsertAll(parsedInvoices)
	assert.Equal(t, 0, cant)
	assert.NotNil(t, err)
}

func TestUpdateOK(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	newInvoice := []domain.Invoices{
		{
			Id:         1,
			Datetime:   "2021-12-8 22:53:32",
			IdCustomer: 5,
			Total:      250.44,
		},
	}

	invoiceEdited := domain.Invoices{
		Id:         1,
		Datetime:   "2021-12-8 22:53:35",
		IdCustomer: 8,
		Total:      255.55,
	}

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	service.InsertAll(newInvoice)

	invoiceResult, err := service.Update(invoiceEdited)

	assert.Nil(t, err)
	assert.Equal(t, invoiceEdited, invoiceResult)
}

func TestUpdateIdInexistent(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	invoiceEdited := domain.Invoices{
		Id:         2,
		Datetime:   "2021-12-8 22:53:32",
		IdCustomer: 5,
		Total:      250.44,
	}

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	invoiceResult, err := service.Update(invoiceEdited)

	assert.Error(t, err)
	assert.NotEqual(t, invoiceEdited, invoiceResult)
}
