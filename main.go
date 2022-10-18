package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/websockets-go/service"
)


func main() {
	app := fiber.New()
	service.SetupRoutes(app)
	
	log.Fatal(app.Listen(":3000"))
}