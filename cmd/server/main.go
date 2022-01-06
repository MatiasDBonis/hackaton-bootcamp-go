package main

import (
	"encoding/csv"
	"io"

	customers "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/customers"
	invoices "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/invoices"
	products "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/products"
	sales "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/sales"
	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/db"
	parser "github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/parser"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func main() {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';' // This is our separator now
		return r
	})

	router := gin.Default()

	repoCustomers := customers.NewRepository(db.StorageDB)
	serviceCustomers := customers.NewService(repoCustomers)

	repoProducts := products.NewRepository(db.StorageDB)
	serviceProducts := products.NewService(repoProducts)

	repoSales := sales.NewRepository(db.StorageDB)
	serviceSales := sales.NewService(repoSales)

	repoInvoices := invoices.NewRepository(db.StorageDB)
	serviceInvoices := invoices.NewService(repoInvoices)

	router.POST("/customers/insertAll", insertAllCustomers(serviceCustomers))
	router.POST("/products/insertAll", insertAllProducts(serviceProducts))
	router.POST("/sales/insertAll", insertAllSales(serviceSales))
	router.POST("/invoices/insertAll", insertAllInvoices(serviceInvoices))

	router.Run()
}

//Endpoints customers
func insertAllCustomers(service customers.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		parsedCustomers, err := parser.ParseDataCustomers()
		if err != nil {
			panic(err.Error())
		}

		rowsAffected, err := service.InsertAll(parsedCustomers)

		if err != nil {
			ctx.String(400, "Hubo un error %v", err.Error())
		} else {
			ctx.String(200, "Registros afectados: %v", rowsAffected)
		}
	}
}

//Endpoints products
func insertAllProducts(service products.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		parsedProducts, err := parser.ParseDataProducts()
		if err != nil {
			panic(err.Error())
		}

		rowsAffected, err := service.InsertAll(parsedProducts)

		if err != nil {
			ctx.String(400, "Hubo un error %v", err.Error())
		} else {
			ctx.String(200, "Registros afectados: %v", rowsAffected)
		}
	}
}

//Endpoints sales
func insertAllSales(service sales.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		parsedSales, err := parser.ParseDataSales()
		if err != nil {
			panic(err.Error())
		}

		rowsAffected, err := service.InsertAll(parsedSales)

		if err != nil {
			ctx.String(400, "Hubo un error %v", err.Error())
		} else {
			ctx.String(200, "Registros afectados: %v", rowsAffected)
		}
	}
}

//Endpoints sales
func insertAllInvoices(service invoices.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		parsedInvoices, err := parser.ParseDataInvoices()
		if err != nil {
			panic(err.Error())
		}

		rowsAffected, err := service.InsertAll(parsedInvoices)

		if err != nil {
			ctx.String(400, "Hubo un error %v", err.Error())
		} else {
			ctx.String(200, "Registros afectados: %v", rowsAffected)
		}
	}
}
