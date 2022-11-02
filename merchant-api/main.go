package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MerchantMessage struct {
	Amount int    `form:"amount" json:"amount"`
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
	Flag   bool   `form:"flag" json:"flag"`
	Desc   string `form:"desc" json:"desc"`
}

func main() {

	app := fiber.New()

	//curl localhost:52321/api/v1/healty
	app.Get("/api/v1/healty", func(c *fiber.Ctx) error {
		return c.SendString("merchant-api healty")
	})

	//curl -X POST -H "Content-Type: application/json"  -d '{"amount": 100,"from": "cobadeff","to": "nabatww","flag": true, "desc":"merchant"}' localhost:52321/api/v1/merchant
	app.Post("/api/v1/merchant", func(c *fiber.Ctx) error {
		fmt.Println("merchant-api")
		var dto MerchantMessage
		c.BodyParser(&dto)

		merchantMessage := &dto
		jsonData, err := json.Marshal(merchantMessage)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))
		return c.JSON(dto)

	})

	app.Listen(":52321")

}
