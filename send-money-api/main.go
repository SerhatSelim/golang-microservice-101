package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type SendMoneyMessage struct {
	Amount int    `form:"amount" json:"amount"`
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
	Flag   bool   `form:"flag" json:"flag"`
	Desc   string `form:"desc" json:"desc"`
}

func main() {

	app := fiber.New()

	//curl localhost:44321/api/v1/healty
	app.Get("/api/v1/healty", func(c *fiber.Ctx) error {
		return c.SendString("send-money-api healty")
	})

	//curl -X POST -H "Content-Type: application/json"  -d '{"amount": 100,"from": "cobadeff","to": "nabatww","flag": true, "desc":"send money"}' localhost:44321/api/v1/send-money
	app.Post("/api/v1/send-money", func(c *fiber.Ctx) error {
		fmt.Println("send-money")
		var dto SendMoneyMessage
		c.BodyParser(&dto)

		sendMoneyMessage := &dto
		jsonData, err := json.Marshal(sendMoneyMessage)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))
		return c.JSON(dto)

	})

	app.Listen(":44321")

}
