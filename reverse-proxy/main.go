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

	//curl -X POST -H "Content-Type: application/json"  -d '{"amount": 100,"from": "cobadeff","to": "nabatww","flag": true, "desc":"merchant"}' localhost:3000/merchant
	app.Post("/:key/*", ProxyHandler)

	//TODO post endpoints
	app.Listen(":3000")

}
