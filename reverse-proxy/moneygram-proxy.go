package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Moneygram struct {
	key string
}

func MoneygramProxy(key string) Moneygram {
	return Moneygram{
		key: key,
	}
}

func (m Moneygram) Accept(key string) bool {
	return m.key == key
}

func (m Moneygram) Proxy(c *fiber.Ctx) error {

	fmt.Printf("Moneygram Proxy ")

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:30002/api/v1/moneygram-gateway/healty", nil)
	if err != nil {
		log.Fatal(err)
	}

	// appending to existing query args
	q := req.URL.Query()
	q.Add("foo", "bar")

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

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
