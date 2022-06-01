package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func PlaceOrderPayment(c *fiber.Ctx) error {

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

	return c.JSON(fiber.Map{
		"userId":  itemId,
		"orderId": orderId,
		"amount":  amount,
	})
}

func CancelOrderPayment(c *fiber.Ctx) error {

	itemId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}
	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"userId": itemId, "orderId": orderId,
	})
}

func GetOrderPayment(c *fiber.Ctx) error {

	itemId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}
	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"userId": itemId, "orderId": orderId,
	})
}

func AddFunds(c *fiber.Ctx) error {

	itemId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}
	amount, err := c.ParamsInt("amount")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"userId": itemId, "amount": amount,
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

	itemId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"userId": itemId,
	})
}
