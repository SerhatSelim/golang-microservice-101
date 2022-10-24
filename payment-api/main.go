package main

import (
	"github.com/gofiber/fiber/v2"
)

type PaymentMessage struct {
	Amount int    `form:"Amount" query:"Amount"`
	From   string `form:"from" query:"from"`
	To     string `form:"to" query:"to"`
}

func main() {

	app := fiber.New()

	//curl localhost:54321/api/v1/healty
	app.Get("/api/v1/healty", func(c *fiber.Ctx) error {
		return c.SendString("payment-api healty")
	})

	//curl -X POST -H "Content-Type: application/json"  -d '{"Amount": 100,"From": "cobadeff","To": "nabatww"}' localhost:54321/api/v1/payment
	app.Post("/api/v1/payment", func(c *fiber.Ctx) error {
		var dto PaymentMessage
		c.BodyParser(&dto)
		c.SendString("payment-api payment")
		return c.JSON(dto)

	})

	app.Listen(":54321")

}
