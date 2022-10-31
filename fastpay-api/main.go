package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type FastPayMessage struct {
	Amount int    `form:"Amount" json:"Amount"`
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
}

func main() {

	app := fiber.New()

	//curl localhost:63321/api/v1/healty
	app.Get("/api/v1/healty", func(c *fiber.Ctx) error {
		return c.SendString("fastpay-api healty")
	})

	//curl -X POST -H "Content-Type: application/json"  -d '{"Amount": 100,"From": "cobadeff","To": "nabatww"}' localhost:63321/api/v1/fastpay
	app.Post("/api/v1/fastpay", func(c *fiber.Ctx) error {
		fmt.Println("fastpay-api payment")
		var dto FastPayMessage
		c.BodyParser(&dto)

		fastPayMessage := &dto
		jsonData, err := json.Marshal(fastPayMessage)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))
		return c.JSON(dto)

	})

	app.Listen(":63321")

}
