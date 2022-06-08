package handlers

import (
	"github.com/gofiber/fiber/v2"
	stockApi "shopping-cart/api/proto/stock"
	"shopping-cart/pkg/stock"
)

func GetStock(c *fiber.Ctx) error {

	itemId := c.Params("item_id")

	// Invalid id / default value returned by c.params
	if itemId == "" {
		return c.SendStatus(400)
	}

	response, err := stock.Find(&stockApi.FindRequest{ItemId: itemId})

	if err != nil {
		return c.SendStatus(400)
	}

	err = c.SendStatus(200)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"stock": response.Stock,
		"price": response.Price,
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
		return c.SendStatus(400)
	}

	return c.SendStatus(200)
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
		return c.SendStatus(400)
	}

	return c.SendStatus(200)
}

func CreateItem(c *fiber.Ctx) error {
	// TODO: Check whether price has to be a float or can be an integer
	price, err := c.ParamsInt("price")
	if err != nil {
		return err
	}

	response, err := stock.Create(&stockApi.CreateRequest{Price: int64(price)})

	if err != nil {
		return c.SendStatus(400)
	}

	err = c.SendStatus(200)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"item_id": response.ItemId,
	})
}
