package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PaymentMessage struct {
	Amount int    `form:"Amount" json:"Amount"`
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
}

func main() {

	app := fiber.New()

	//curl localhost:30001/api/v1/gateway/healty
	app.Get("/api/v1/gateway/healty", func(c *fiber.Ctx) error {
		return c.SendString("payment-gateway healty")

	})

	// //curl localhost:30001/api/v1/gateway/payment-api/healty
	app.Get("/api/v1/gateway/payment-api/healty", func(c *fiber.Ctx) error {
		fmt.Println("/api/v1/gateway/payment-api/healty")
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodGet, "http://localhost:54321/api/v1/healty", nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Errored when sending request to the server")
			return err
		}

		defer resp.Body.Close()
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.Status)
		fmt.Println(string(responseBody))

		return c.SendString(string(responseBody))
	})

	//curl -X POST -H "Content-Type: application/json"  -d '{"Amount": 100,"From": "cobadeff","To": "nabatww"}' localhost:30001/api/v1/gateway/payment
	app.Post("/api/v1/gateway/payment", func(c *fiber.Ctx) error {
		fmt.Println("api/v1/gateway/payment")

		var dto PaymentMessage
		c.BodyParser(&dto)
		paymentMessage := &dto
		jsonData, err := json.Marshal(paymentMessage)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))

		req, err := http.NewRequest(http.MethodPost, "http://localhost:54321/api/v1/payment", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Errored when sending request to the server")
			return err
		}

		defer resp.Body.Close()
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.Status)
		fmt.Println(string(responseBody))
		return c.SendString(string(responseBody))
	})

	app.Listen(":30001")

}
