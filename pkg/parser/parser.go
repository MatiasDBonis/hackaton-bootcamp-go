package parser

import (
	"fmt"
	"os"
	"strings"

	customers "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/customers"
	invoices "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/invoices"
	products "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/products"
	sales "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/sales"
	"github.com/gocarina/gocsv"
)

func ParseDataCustomers() ([]customers.Customers, error) {
	var customersSlice []customers.Customers

	data, err := os.ReadFile("../../datos/customers.txt")
	if err != nil {
		return []customers.Customers{}, err
	}

	dataString := string(data)
	dataStringReplaced := strings.ReplaceAll(dataString, "#$%#", ";")

	titles := "id;last_name;first_name;condition\n"

	finalDataString := titles + dataStringReplaced

	gocsv.UnmarshalString(finalDataString, &customersSlice)

	// fmt.Println("Customer:")
	// fmt.Println(finalDataString)
	// fmt.Println("SALTO DE LINEA")
	// fmt.Println(customersSlice)

	return customersSlice, nil
}

func ParseDataProducts() ([]products.Products, error) {
	var productsSlice []products.Products

	data, err := os.ReadFile("../../datos/products.txt")
	if err != nil {
		return []products.Products{}, err
	}

	dataString := string(data)
	dataStringReplaced := strings.ReplaceAll(dataString, "#$%#", ";")

	titles := "id;description;price\n"

	finalDataString := titles + dataStringReplaced
	gocsv.UnmarshalString(finalDataString, &productsSlice)

	fmt.Println("Products:")
	fmt.Println(finalDataString)
	fmt.Println("SALTO DE LINEA")
	fmt.Println(productsSlice)

	return productsSlice, nil
}

func ParseDataSales() ([]sales.Sales, error) {
	var salesSlice []sales.Sales

	data, err := os.ReadFile("../../datos/sales.txt")
	if err != nil {
		return []sales.Sales{}, err
	}
	dataString := string(data)
	dataStringReplaced := strings.ReplaceAll(dataString, "#$%#", ";")

	titles := "id;id_invoice;id_product;quantity\n"

	finalDataString := titles + dataStringReplaced
	gocsv.UnmarshalString(finalDataString, &salesSlice)

	fmt.Println("Sales:")
	fmt.Println(finalDataString)
	fmt.Println("SALTO DE LINEA")
	fmt.Println(salesSlice)

	return salesSlice, nil
}

func ParseDataInvoices() ([]invoices.Invoices, error) {
	var invoicesSlice []invoices.Invoices

	data, err := os.ReadFile("../../datos/invoices.txt")
	if err != nil {
		return []invoices.Invoices{}, err
	}

	dataString := string(data)
	dataStringReplaced := strings.ReplaceAll(dataString, "#$%#", ";")

	titles := "id;datetime;id_customer;total\n"

	finalDataString := titles + dataStringReplaced
	gocsv.UnmarshalString(finalDataString, &invoicesSlice)

	//fmt.Println("Invoices:")
	//fmt.Println(finalDataString)
	//fmt.Println("SALTO DE LINEA")
	//fmt.Println(invoicesSlice)

	return invoicesSlice, nil
}
