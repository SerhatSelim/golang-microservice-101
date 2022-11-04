package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//curl localhost:3000/payment
	//curl localhost:3000/moneygram
	app.Get("/:key/*", ProxyHandler)

	//TODO post endpoints
	app.Listen(":3000")

}
