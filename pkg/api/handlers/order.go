package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	orderApi "shopping-cart/api/proto/order"
	"shopping-cart/pkg/order"
)

// TODO: This is not a function belonging to the Order, it should be in payment
func CreateUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	// Invalid id / default value returned by c.params
	if userId == "" {
		return c.SendStatus(400)
	}

	response, err := order.CreateOrder(&orderApi.CreateOrderRequest{UserId: userId})

	if err != nil {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"userId": response.OrderId,
	})
}

func GetOrder(c *fiber.Ctx) error {
	orderId := c.Params("orderId")

	// Invalid id / default value returned by c.params
	if orderId == "" {
		return c.SendStatus(400)
	}

	response, err := order.GetOrder(&orderApi.GetOrderRequest{OrderId: orderId})

	if err != nil {
		return c.SendStatus(400)
	}

	err = c.SendStatus(200)
	if err != nil {
		return err
	}

	fmt.Println(response)
	//grpc serializes / deserializes empty array as null/ nil or something
	if response.ItemIds == nil {
		response.ItemIds = []string{}
	}

	return c.JSON(fiber.Map{
		"order_id":   response.OrderId,
		"paid":       response.Paid,
		"items":      response.ItemIds,
		"user_id":    response.UserId,
		"total_cost": response.TotalCost,
	})
}

func AddItem(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	// Invalid id / default value returned by c.params
	if orderId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	itemId := c.Params("itemId")
	// Invalid id / default value returned by c.params
	if itemId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err := order.AddItem(&orderApi.AddItemRequest{OrderId: orderId, ItemId: itemId})

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteItem(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	// Invalid id / default value returned by c.params
	if orderId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	itemId := c.Params("itemId")
	// Invalid id / default value returned by c.params
	if itemId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err := order.RemoveItem(&orderApi.RemoveItemRequest{OrderId: orderId, ItemId: itemId})

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteOrder(c *fiber.Ctx) error {
	orderId := c.Params("orderId")

	// Invalid id / default value returned by c.params
	if orderId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err := order.RemoveOrder(&orderApi.RemoveOrderRequest{OrderId: orderId})

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func CreateOrder(c *fiber.Ctx) error {
	userId := c.Params("userId")

	// Invalid id / default value returned by c.params
	if userId == "" {
		return c.SendStatus(400)
	}

	response, err := order.CreateOrder(&orderApi.CreateOrderRequest{UserId: userId})

	if err != nil {
		return c.SendStatus(400)
	}

	err = c.SendStatus(200)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"order_id": response.OrderId,
	})
}
