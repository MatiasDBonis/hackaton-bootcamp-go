package sales

import (
	"testing"

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
