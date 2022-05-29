package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {

	userId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"userId": userId,
	})
}

func GetOrder(c *fiber.Ctx) error {
	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"orderId": orderId,
	})
}

func AddItem(c *fiber.Ctx) error {

	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}
	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"orderId": orderId,
		"itemId":  itemId,
	})
}

func DeleteItem(c *fiber.Ctx) error {

	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}
	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"orderId": orderId,
		"itemId":  itemId,
	})
}

func DeleteOrder(c *fiber.Ctx) error {

	userId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"userId": userId,
	})
}

func CreateOrder(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("orderId")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"orderId": userId,
	})
}
