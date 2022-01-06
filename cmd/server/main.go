package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("../../datos/")
	if err == nil {
		file := string(data)
		fmt.Println(file)
	} else {
		fmt.Println("El archivo no existe...")
	}
}

func parseCustomers(customers *[]Customers)
