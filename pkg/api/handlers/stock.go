package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetStock(c *fiber.Ctx) error {

	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"itemId": itemId,
	})
}

func SubtractStock(c *fiber.Ctx) error {

	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}
	amount, err := c.ParamsInt("amount")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"itemId": itemId, "amount": amount,
	})
}

func AddStock(c *fiber.Ctx) error {

	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}
	amount, err := c.ParamsInt("amount")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"itemId": itemId, "amount": amount,
	})
}

func CreateItem(c *fiber.Ctx) error {
	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}
	price, err := c.ParamsInt("price")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"itemId": itemId, "price": price,
	})
}
