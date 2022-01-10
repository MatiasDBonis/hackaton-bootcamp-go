package products

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

	parsedProducts, err := parser.ParseDataProducts()
	assert.Nil(t, err)
	fmt.Println(parsedProducts)

	cant, err := service.InsertAll(parsedProducts)
	assert.Nil(t, err)
	assert.Equal(t, 100, cant)
}

func TestInsertExistent(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	parsedProducts, err := parser.ParseDataProducts()
	assert.Nil(t, err)

	cant, err := service.InsertAll(parsedProducts)
	assert.Equal(t, 100, cant)
	assert.Nil(t, err)

	cant, err = service.InsertAll(parsedProducts)
	assert.Equal(t, 0, cant)
	assert.NotNil(t, err)
}

func TestUpdateOK(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	newProduct := []domain.Products{
		{
			Id:          1,
			Description: "Beer - Scottish",
			Price:       25.99,
		},
	}

	editedProduct := domain.Products{
		Id:          1,
		Description: "Beer - Golden Ipa",
		Price:       30.55,
	}

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	service.InsertAll(newProduct)

	productResult, err := service.Update(editedProduct)

	assert.Nil(t, err)
	assert.Equal(t, editedProduct, productResult)
}

func TestUpdateIdInexistent(t *testing.T) {
	db, err := db.InitTxdb()
	assert.Nil(t, err)

	editedProduct := domain.Products{
		Id:          1,
		Description: "Beer - Scottish",
		Price:       25.99,
	}

	repo := NewRepository(db)
	defer db.Close()

	service := NewService(repo)

	productResult, err := service.Update(editedProduct)

	assert.Error(t, err)
	assert.NotEqual(t, editedProduct, productResult)
}
