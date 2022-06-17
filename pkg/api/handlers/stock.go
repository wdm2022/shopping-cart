package handlers

import (
	"log"
	stockApi "shopping-cart/api/proto/stock"
	"shopping-cart/pkg/stock"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetStock(c *fiber.Ctx) error {

	itemId := c.Params("item_id")

	// Invalid id / default value returned by c.params
	if itemId == "" {
		return c.SendStatus(400)
	}

	response, err := stock.Find(&stockApi.FindRequest{ItemId: itemId})
	if err != nil {
		log.Printf("Error when executing Find: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	float := float64(response.Price) / 100.0
	return c.JSON(fiber.Map{
		"stock": response.Stock,
		"price": float,
	})
}

func SubtractStock(c *fiber.Ctx) error {
	itemId := c.Params("item_id")
	// Invalid id / default value returned by c.params
	if itemId == "" {
		return c.SendStatus(400)
	}

	amount, err := c.ParamsInt("amount")
	if err != nil {
		return err
	}

	_, err = stock.Subtract(&stockApi.SubtractRequest{ItemId: itemId, Amount: int64(amount)})
	if err != nil {
		log.Printf("Error when executing Subtract: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func AddStock(c *fiber.Ctx) error {
	itemId := c.Params("item_id")
	// Invalid id / default value returned by c.params
	if itemId == "" {
		return c.SendStatus(400)
	}

	amount, err := c.ParamsInt("amount")
	if err != nil {
		return err
	}

	_, err = stock.Add(&stockApi.AddRequest{ItemId: itemId, Amount: int64(amount)})
	if err != nil {
		log.Printf("Error when executing Add: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func CreateItem(c *fiber.Ctx) error {
	priceStr := c.Params("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return err
	}
	price *= 100

	response, err := stock.Create(&stockApi.CreateRequest{Price: int64(price)})
	if err != nil {
		log.Printf("Error when executing Create: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"item_id": response.ItemId,
	})
}
