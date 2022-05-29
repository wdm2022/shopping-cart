package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func PlaceOrderPayment(c *fiber.Ctx) error {
	client := &http.Client{}

	itemId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}
	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}
	amount, err := c.ParamsInt("amount")
	if err != nil {
		return err
	}
	values := map[string]int{"userId": itemId, "orderId": orderId, "amount": amount}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "place-order", bytes.NewBuffer(data))
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

func CancelOrderPayment(c *fiber.Ctx) error {
	client := &http.Client{}

	itemId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}
	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}
	values := map[string]int{"userId": itemId, "orderId": orderId}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "cancel-order", bytes.NewBuffer(data))
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

func GetOrderPayment(c *fiber.Ctx) error {
	client := &http.Client{}

	itemId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}
	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}
	values := map[string]int{"userId": itemId, "orderId": orderId}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", "get-order", bytes.NewBuffer(data))
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

func AddFunds(c *fiber.Ctx) error {
	client := &http.Client{}

	itemId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}
	amount, err := c.ParamsInt("amount")
	if err != nil {
		return err
	}
	values := map[string]int{"userId": itemId, "amount": amount}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "add-funds", bytes.NewBuffer(data))
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

func CreatePaymentUser(c *fiber.Ctx) error {
	resp, err := http.Post("create-order", "order", nil)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"resp": resp,
	})
}

func GetUser(c *fiber.Ctx) error {
	client := &http.Client{}

	itemId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}
	values := map[string]int{"userId": itemId}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", "get-user", bytes.NewBuffer(data))
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
