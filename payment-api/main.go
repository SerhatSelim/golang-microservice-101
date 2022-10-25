package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type PaymentMessage struct {
	Amount int    `form:"Amount" json:"Amount"`
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
}

func main() {

	app := fiber.New()

	//curl localhost:54321/api/v1/healty
	app.Get("/api/v1/healty", func(c *fiber.Ctx) error {
		return c.SendString("payment-api healty")
	})

	//curl -X POST -H "Content-Type: application/json"  -d '{"Amount": 100,"From": "cobadeff","To": "nabatww"}' localhost:54321/api/v1/payment
	app.Post("/api/v1/payment", func(c *fiber.Ctx) error {
		fmt.Println("payment-api payment")
		var dto PaymentMessage
		c.BodyParser(&dto)

		paymentMessage := &dto
		jsonData, err := json.Marshal(paymentMessage)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))
		return c.JSON(dto)

	})

	app.Listen(":54321")

}
