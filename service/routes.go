package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vsantos1/websockets-go/middleware"
)



func SetupRoutes(app *fiber.App) {
	app.Use(middleware.Sameorigin)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	
	app.Get("/chat/:id",websocket.New(Websocket))
	

}