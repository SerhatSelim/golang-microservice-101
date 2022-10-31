package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type FastPayMessage struct {
	Amount int    `form:"amount" json:"amount"`
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
	Flag   bool   `form:"flag" json:"flag"`
}

func main() {

	app := fiber.New()

	//curl localhost:63321/api/v1/healty
	app.Get("/api/v1/healty", func(c *fiber.Ctx) error {
		return c.SendString("fastpay-api healty")
	})

	//curl -X POST -H "Content-Type: application/json"  -d '{"amount": 100,"from": "cobadeff","to": "nabatww","flag": true}' localhost:63321/api/v1/fastpay
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
