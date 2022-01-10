package main

import (
	"fmt"
	"strconv"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/customers"
	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/invoices"
	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/products"
	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/sales"
	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/db"
	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/parser"
	"github.com/gin-gonic/gin"
)

func main() {
	// gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
	// 	r := csv.NewReader(in)
	// 	r.Comma = ';' // This is our separator now
	// 	return r
	// })

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

	router.PUT("/customers/update/:id", updateCustomers(serviceCustomers))
	router.PUT("/products/update/:id", updateProducts(serviceProducts))
	router.PUT("/sales/update/:id", updateSales(serviceSales))
	router.PUT("/invoices/update/:id", updateInvoices(serviceInvoices))

	router.GET("/customers", GetTotals(serviceCustomers))

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

func updateCustomers(service customers.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var customer domain.Customers
		err := ctx.ShouldBindJSON(&customer)
		if err != nil {
			ctx.String(400, "Error al procesar JSON")
		}

		paramId := ctx.Param("id")

		customer.Id, err = strconv.Atoi(paramId)

		if err != nil {
			ctx.String(400, "Error al procesar el ID")
		}

		rowsAffected, err := service.Update(customer)

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

func updateProducts(service products.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var product domain.Products
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.String(400, "Error al procesar JSON")
		}
		paramId := ctx.Param("id")

		product.Id, err = strconv.Atoi(paramId)

		if err != nil {
			ctx.String(400, "Error al procesar el ID")
		}

		rowsAffected, err := service.Update(product)

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
			ctx.String(500, err.Error())
		}
		fmt.Println(parsedSales)
		rowsAffected, err := service.InsertAll(parsedSales)

		if err != nil {
			ctx.String(400, "Hubo un error %v", err.Error())
		} else {
			ctx.String(200, "Registros afectados: %v", rowsAffected)
		}
	}
}

func updateSales(service sales.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var sale domain.Sales
		err := ctx.ShouldBindJSON(&sale)
		if err != nil {
			ctx.String(400, "Error al procesar JSON")
		}
		paramId := ctx.Param("id")

		sale.Id, err = strconv.Atoi(paramId)

		if err != nil {
			ctx.String(400, "Error al procesar el ID")
		}

		rowsAffected, err := service.Update(sale)

		if err != nil {
			ctx.String(400, "Hubo un error %v", err.Error())
		} else {
			ctx.String(200, "Registros afectados: %v", rowsAffected)
		}
	}
}

//Endpoints invoices
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

func updateInvoices(service invoices.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var invoice domain.Invoices
		err := ctx.ShouldBindJSON(&invoice)
		if err != nil {
			ctx.String(400, "Error al procesar JSON")
		}
		paramId := ctx.Param("id")

		invoice.Id, err = strconv.Atoi(paramId)

		if err != nil {
			ctx.String(400, "Error al procesar el ID")
		}

		rowsAffected, err := service.Update(invoice)

		if err != nil {
			ctx.String(400, "Hubo un error %v", err.Error())
		} else {
			ctx.String(200, "Registros afectados: %v", rowsAffected)
		}
	}
}

func GetTotals(service customers.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		data, err := service.GetTotals()

		if err != nil {
			ctx.String(400, "Hubo un error: %v", err.Error())
		} else {
			ctx.String(200, "Data: %v", data)
		}
	}
}
