package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/k0kubun/pp"
	"pesa_api/config"
	"pesa_api/config/database"
	"pesa_api/server/routes"
)

func main() {
	//app
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	//db
	_, err := database.Connect()
	if err != nil {
		log.Errorf("error occurred while connecting to db: %v", err)
	}
	defer database.Close()
	//router
	routes.Routers(app)
	//server
	cfg := config.New()
	address := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
	pp.Println("Server on: ", address)
	err = app.Listen(address)
	if err != nil {
		log.Errorf("error occurred while start server: %v", err)
	}

}
