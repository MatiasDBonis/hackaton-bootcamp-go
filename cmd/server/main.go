package main

import (
	"fmt"
	"os"
	"strings"

	customers "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/customers"
	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func main() {
	customerFileName := "customers.txt"
	customersSlice, err := parseCustomers(customerFileName)
	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()

	repoCustomers := customers.NewRepository(db.StorageDB)
	serviceCustomers := customers.NewService(repoCustomers)

	//docs.SwaggerInfo.Host = os.Getenv("HOST")
	//router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//	router.Use(TokenAuthMiddleware())

	router.POST("/customers/insertAll", insertAll(serviceCustomers, customersSlice))

	router.Run()
}

func parseCustomers(fileName string) ([]customers.Customers, error) {
	var customersSlice []customers.Customers

	data, err := os.ReadFile("../../datos/" + fileName)
	if err != nil {
		return []customers.Customers{}, err
	}

	dataString := string(data)
	dataStringReplaced := strings.ReplaceAll(dataString, "#$%#", ",")

	titles := "id,last_name,first_name,condition\n"

	finalDataString := titles + dataStringReplaced
	// json.Unmarshal([]byte(dataStringReplaced), &customers)
	gocsv.UnmarshalString(finalDataString, &customersSlice)

	fmt.Println(dataStringReplaced)
	fmt.Println("SALTO DE LINEA")
	fmt.Println(customersSlice)

	return customersSlice, nil
}

func insertAll(service customers.Service, parsedCustomers []customers.Customers) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		rowsAffected, err := service.InsertAll(parsedCustomers)

		if err != nil {
			ctx.String(400, "Hubo un error %v", err.Error())
		} else {
			ctx.String(200, "Registros afectados: %v", rowsAffected)
		}
	}
}
