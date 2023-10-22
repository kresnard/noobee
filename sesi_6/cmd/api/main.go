package main

import (
	"log"
	"sesi_6/config"
	"sesi_6/pkg/databse"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New(fiber.Config{
		AppName: "Produts services",
		Prefork: true,
	})

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error connect to config.yaml", err.Error())
	}

	db, err := databse.ConnectGORMPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	router.Listen(config.Cfg.App.Port)
}
