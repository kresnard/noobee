package main

import (
	"log"
	"sesi_6/app/product"
	"sesi_6/config"
	"sesi_6/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New(fiber.Config{
		AppName: "Products services",
		Prefork: true,
	})

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error connect to config.yaml", err.Error())
	}

	db, err := database.ConnectGORMPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	dbSqlx, err := database.ConnectSqlxPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	product.RegisterServiceProduct(router, db, dbSqlx)

	router.Listen(config.Cfg.App.Port)
}
