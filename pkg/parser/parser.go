package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	domain "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
	"github.com/gocarina/gocsv"
)

func ParseDataCustomers() ([]domain.Customers, error) {
	setPuntoYComaSeparator()
	var customersSlice []domain.Customers

	data, err := os.ReadFile("../../datos/customers.txt")
	if err != nil {
		return []domain.Customers{}, err
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

func ParseDataProducts() ([]domain.Products, error) {
	setPuntoYComaSeparator()
	var productsSlice []domain.Products

	data, err := os.ReadFile("../../datos/products.txt")
	if err != nil {
		return []domain.Products{}, err
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

func ParseDataSales() ([]domain.Sales, error) {
	setPuntoYComaSeparator()
	var salesSlice []domain.Sales

	data, err := os.ReadFile("../../datos/sales.txt")
	if err != nil {
		return []domain.Sales{}, err
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

func ParseDataInvoices() ([]domain.Invoices, error) {
	setPuntoYComaSeparator()
	var invoicesSlice []domain.Invoices

	data, err := os.ReadFile("../../datos/invoices.txt")
	if err != nil {
		return []domain.Invoices{}, err
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

func setPuntoYComaSeparator() {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';' // This is our separator now
		return r
	})
}
