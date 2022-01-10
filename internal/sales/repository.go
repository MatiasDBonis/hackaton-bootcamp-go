package sales

import (
	"database/sql"
	"errors"
	"log"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
)

type Repository interface {
	Insert(customer domain.Sales) (domain.Sales, error)
	Update(customer domain.Sales) (domain.Sales, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(sale domain.Sales) (domain.Sales, error) {
	stmt, err := r.db.Prepare("INSERT INTO sales (id, idinvoice, idproduct, quantity) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(sale.Id, sale.IdInvoice, sale.IdProduct, sale.Quantity)
	if err != nil {
		return domain.Sales{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.Sales{}, errors.New("no se insertó la sale")
	}

	return sale, nil
}

func (r *repository) Update(sale domain.Sales) (domain.Sales, error) {
	stmt, err := r.db.Prepare("UPDATE sales SET idinvoice = ?, idproduct = ?, quantity = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(sale.IdInvoice, sale.IdProduct, sale.Quantity, sale.Id)
	if err != nil {
		return domain.Sales{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.Sales{}, errors.New("no se actualizó la sale")
	}

	return sale, nil
}
