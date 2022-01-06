package main

import (
	"fmt"
	"os"
	"strings"

	internal "github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/customers"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func main() {
	customerFileName := "customers.txt"
	customers, err := parseCustomers(customerFileName)
	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()

	db := store.New(store.FileType, "./personasSalida.json")
	repo := personas.NewRepository(db)
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	//docs.SwaggerInfo.Host = os.Getenv("HOST")
	//router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//	router.Use(TokenAuthMiddleware())

	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())
	router.PUT("/personas/:id", controller.Update())
	router.PATCH("/personas/:id", controller.UpdateNombre())
	router.DELETE("/personas/:id", controller.Delete())

	router.Run()
}

func parseCustomers(fileName string) ([]internal.Customers, error) {
	var customers []internal.Customers

	data, err := os.ReadFile("../../datos/" + fileName)
	if err != nil {
		return []internal.Customers{}, err
	}

	dataString := string(data)
	dataStringReplaced := strings.ReplaceAll(dataString, "#$%#", ",")

	titles := "id,last_name,first_name,condition\n"

	finalDataString := titles + dataStringReplaced
	// json.Unmarshal([]byte(dataStringReplaced), &customers)
	gocsv.UnmarshalString(finalDataString, &customers)

	fmt.Println(dataStringReplaced)
	fmt.Println("SALTO DE LINEA")
	fmt.Println(customers)

	return customers, nil
}
