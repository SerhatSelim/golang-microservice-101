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

type SendMoneyMessage struct {
	Amount int    `form:"amount" json:"amount"`
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
	Flag   bool   `form:"flag" json:"flag"`
	Desc   string `form:"desc" json:"desc"`
}

type MerchantMessage struct {
	Amount int    `form:"amount" json:"amount"`
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
	Flag   bool   `form:"flag" json:"flag"`
	Desc   string `form:"desc" json:"desc"`
}

func main() {

	app := fiber.New()

	//curl localhost:30002/api/v1/moneygram-gateway/healty
	app.Get("/api/v1/moneygram-gateway/healty", func(c *fiber.Ctx) error {
		return c.SendString("payment-gateway healty")

	})

	//curl localhost:30002/api/v1/moneygram-gateway/send-money-api/healty
	app.Get("/api/v1/moneygram-gateway/send-money-api/healty", func(c *fiber.Ctx) error {
		fmt.Println("/api/v1/moneygram-gateway/send-money-api/healty")
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodGet, "http://localhost:44321/api/v1/healty", nil)
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

	//curl -X POST -H "Content-Type: application/json"  -d '{"amount": 100,"from": "cobadeff","to": "nabatww","flag": true, "desc":"send money"}' localhost:30002/api/v1/moneygram-gateway/send-money
	app.Post("/api/v1/moneygram-gateway/send-money", func(c *fiber.Ctx) error {
		fmt.Println("api/v1/moneygram-gateway/send-money")

		var dto SendMoneyMessage
		c.BodyParser(&dto)
		sendMoneyMessage := &dto
		jsonData, err := json.Marshal(sendMoneyMessage)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))

		req, err := http.NewRequest(http.MethodPost, "http://localhost:44321/api/v1/send-money", bytes.NewBuffer(jsonData))
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

	//curl localhost:30002/api/v1/moneygram-gateway/merchant-api/healty
	app.Get("/api/v1/moneygram-gateway/merchant-api/healty", func(c *fiber.Ctx) error {
		fmt.Println("/api/v1/moneygram-gateway/merchant-api/healty")
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodGet, "http://localhost:52321/api/v1/healty", nil)
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

	//curl -X POST -H "Content-Type: application/json"  -d '{"amount": 100,"from": "cobadeff","to": "nabatww","flag": true, "desc":"merchant"}' localhost:30002/api/v1/moneygram-gateway/merchant
	app.Post("/api/v1/moneygram-gateway/merchant", func(c *fiber.Ctx) error {
		fmt.Println("api/v1/moneygram-gateway/merchant")

		var dto MerchantMessage
		c.BodyParser(&dto)
		merchantMessage := &dto
		jsonData, err := json.Marshal(merchantMessage)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))

		req, err := http.NewRequest(http.MethodPost, "http://localhost:52321/api/v1/merchant", bytes.NewBuffer(jsonData))
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

	app.Listen(":30002")

}
