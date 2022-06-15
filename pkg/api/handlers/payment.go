package handlers

import (
	"fmt"
	"log"
	paymentApi "shopping-cart/api/proto/payment"
	"shopping-cart/pkg/payment"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PlaceOrderPayment(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	// Invalid id / default value returned by c.params
	if userId == "" {
		return c.SendStatus(400)
	}
	orderId := c.Params("order_id")
	// Invalid id / default value returned by c.params
	if orderId == "" {
		return c.SendStatus(400)
	}

	amountStr := c.Params("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return err
	}
	amount *= 100

	_, err = payment.Pay(&paymentApi.PayRequest{UserId: userId, OrderId: orderId, Amount: int64(amount)})
	if err != nil {
		log.Printf("Error when executing Pay: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func CancelOrderPayment(c *fiber.Ctx) error {

	userId := c.Params("user_id")
	// Invalid id / default value returned by c.params
	if userId == "" {
		return c.SendStatus(400)
	}
	orderId := c.Params("order_id")
	// Invalid id / default value returned by c.params
	if orderId == "" {
		return c.SendStatus(400)
	}
	_, err := payment.Cancel(&paymentApi.CancelRequest{UserId: userId, OrderId: orderId})
	if err != nil {
		log.Printf("Error when executing Cancel: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetOrderPayment(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	// Invalid id / default value returned by c.params
	if userId == "" {
		return c.SendStatus(400)
	}
	orderId := c.Params("order_id")
	// Invalid id / default value returned by c.params
	if orderId == "" {
		return c.SendStatus(400)
	}

	response, err := payment.Status(&paymentApi.StatusRequest{UserId: userId, OrderId: orderId})
	if err != nil {
		log.Printf("Error when executing Status: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"paid": response.Paid,
	})
}

func AddFunds(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	// Invalid id / default value returned by c.params
	if userId == "" {
		return c.SendStatus(400)
	}
	amountStr := c.Params("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		logger := c.App().Server().Logger
		logger.Printf("error opening file: %v", err)
		return err
	}
	amount *= 100
	response, err := payment.AddFunds(&paymentApi.AddFundsRequest{UserId: userId, Amount: int64(amount)})
	if err != nil {
		log.Printf("Error when executing AddFunds: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"done": response.Success,
	})
}

// TODO: which operation should this function call?
func CreatePaymentUser(c *fiber.Ctx) error {
	user, err := payment.CreateUser(&paymentApi.EmptyMessage{})
	if err != nil {
		log.Printf("Error when executing CreateUser: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"user_id": user.UserId,
	})
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("user_id")

	// Invalid id / default value returned by c.params
	if userId == "" {
		return c.SendStatus(400)
	}

	response, err := payment.FindUser(&paymentApi.FindUserRequest{UserId: userId})
	if err != nil {
		log.Printf("Error when executing FindUser: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"user_id": response.UserId,
		"credit":  fmt.Sprintf("%f", float64(response.Credits)/100.0),
	})
}
