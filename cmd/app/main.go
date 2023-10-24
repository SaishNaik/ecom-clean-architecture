package main

import (
	"ecom_clean_architecture/config"
	"ecom_clean_architecture/dependency_injection"
	"ecom_clean_architecture/infrastructure/datastore"
	"ecom_clean_architecture/infrastructure/router"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	//load config
	config.ReadConfig()

	// orm and connection to database
	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	di := dependency_injection.NewDepInjection(db)

	e := echo.New()
	router.NewRouter(e, di.NewAppController())

	fmt.Println("Server listening at port:" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
