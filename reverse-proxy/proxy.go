package main

import (
	"github.com/gofiber/fiber/v2"
)

type Proxy interface {
	Accept(key string) bool
	Proxy(c *fiber.Ctx) error
}

var Proxies = []Proxy{
	MoneygramProxy("moneygram"),
	PaymentProxy("payment"),
	MerchantProxy("merchant"),
}

func ProxyHandler(c *fiber.Ctx) error {
	for _, v := range Proxies {
		if v.Accept(c.Params("key")) {
			return v.Proxy(c)
		}
	}

	c.Response().SetStatusCode(404)
	return nil
}
