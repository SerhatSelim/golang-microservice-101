package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/api/v1/healty", func(c *fiber.Ctx) error {
		return c.SendString("payment-api healty")
	})

	app.Post("/api/v1/payment", func(c *fiber.Ctx) error {
		c.Method() // "POST"
		return c.SendString("payment-api payment")

	})

	app.Listen(":54321")

}
