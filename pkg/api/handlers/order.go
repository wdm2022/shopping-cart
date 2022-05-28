package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateOrder(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}
	values := map[string]int{"userId": userId}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	resp, err := http.Post("create-order", "order", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return resp.StatusCode
}

func DeleteOrder(c *fiber.Ctx) error {
	client := &http.Client{}

	userId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}
	values := map[string]int{"userId": userId}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", "delete-order", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)

	return resp.StatusCode
}

func GetOrder(c *fiber.Ctx) error {
	client := &http.Client{}

	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}
	values := map[string]int{"orderId": orderId}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", "get-order", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)

	return resp
}

func AddItem(c *fiber.Ctx) error {
	client := &http.Client{}

	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}
	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}
	values := map[string]int{"orderId": orderId, "itemId": itemId}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "add-item", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)

	return resp
}

func DeleteItem(c *fiber.Ctx) error {
	client := &http.Client{}

	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}
	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}
	values := map[string]int{"orderId": orderId, "itemId": itemId}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", "delete-item", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)

	return resp
}
