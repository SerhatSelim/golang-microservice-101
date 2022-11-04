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

type MerchantMessage struct {
	Amount int    `form:"amount" json:"amount"`
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
	Flag   bool   `form:"flag" json:"flag"`
	Desc   string `form:"desc" json:"desc"`
}

type Merchant struct {
	key string
}

func MerchantProxy(key string) Merchant {
	return Merchant{
		key: key,
	}
}

func (m Merchant) Accept(key string) bool {
	return m.key == key
}

func (m Merchant) Proxy(c *fiber.Ctx) error {

	var dto MerchantMessage
	c.BodyParser(&dto)
	merchantMessage := &dto
	jsonData, err := json.Marshal(merchantMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	req, err := http.NewRequest(http.MethodPost, "localhost:30002/api/v1/moneygram-gateway/merchant", bytes.NewBuffer(jsonData))
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

	return nil
}
