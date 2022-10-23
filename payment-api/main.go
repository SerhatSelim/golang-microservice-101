package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/healty", func(c *fiber.Ctx) error {
		return c.SendString("payment-api running")
	})

	app.Listen(":54321")

}
