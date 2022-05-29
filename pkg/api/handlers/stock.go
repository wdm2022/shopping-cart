package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetStock(c *fiber.Ctx) error {
	client := &http.Client{}

	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}
	values := map[string]int{"itemId": itemId}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", "get-stock", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"resp": resp,
	})
}

func SubtractStock(c *fiber.Ctx) error {
	client := &http.Client{}

	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}
	amount, err := c.ParamsInt("amount")
	if err != nil {
		return err
	}
	values := map[string]int{"itemId": itemId, "amount": amount}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "subtract-stock", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"resp": resp,
	})
}

func AddStock(c *fiber.Ctx) error {
	client := &http.Client{}

	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}
	amount, err := c.ParamsInt("amount")
	if err != nil {
		return err
	}
	values := map[string]int{"itemId": itemId, "amount": amount}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "add-stock", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"resp": resp,
	})
}

func CreateItem(c *fiber.Ctx) error {
	client := &http.Client{}

	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}
	price, err := c.ParamsInt("price")
	if err != nil {
		return err
	}
	values := map[string]int{"itemId": itemId, "price": price}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "create-item", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"resp": resp,
	})
}
