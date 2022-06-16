package api

import (
	"fmt"
	"shopping-cart/pkg/api/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func bindOrdersApi(app *fiber.App) {
	orders := app.Group("/orders")

	orders.Post("/create/:user_id", handlers.CreateOrder)
	orders.Delete("/remove/:order_id", handlers.DeleteOrder)
	orders.Get("/find/:order_id", handlers.GetOrder)
	orders.Post("/addItem/:order_id/:item_id", handlers.AddItem)
	orders.Delete("/removeItem/:order_id/:item_id", handlers.DeleteItem)
	orders.Post("/checkout/:order_id", handlers.Checkout)
}

func bindStockApi(app *fiber.App) {
	stock := app.Group("/stock")

	stock.Get("/find/:item_id", handlers.GetStock)
	stock.Post("/subtract/:item_id/:amount", handlers.SubtractStock)
	stock.Post("/add/:item_id/:amount", handlers.AddStock)
	stock.Post("/item/create/:price", handlers.CreateItem)
}

func bindPaymentApi(app *fiber.App) {
	payment := app.Group("/payment")

	payment.Post("/pay/:user_id/:order_id/:amount", handlers.PlaceOrderPayment)
	payment.Post("/cancel/:user_id/:order_id", handlers.CancelOrderPayment)
	payment.Get("/status/:user_id/:order_id", handlers.GetOrderPayment)
	payment.Post("/add_funds/:user_id/:amount", handlers.AddFunds)
	payment.Post("/create_user", handlers.CreatePaymentUser)
	payment.Get("/find_user/:user_id", handlers.GetUser)
}

func bindKubernetesApi(app *fiber.App) {
	kube := app.Group("/kube")

	kube.Get("/liveness", handlers.Liveness)
	kube.Get("/readiness", handlers.Readiness)
}

func RunHttpServer(port *int, prefork *bool) error {
	app := fiber.New(fiber.Config{
		Prefork: *prefork,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	bindOrdersApi(app)
	bindStockApi(app)
	bindPaymentApi(app)
	bindKubernetesApi(app)

	return app.Listen(fmt.Sprintf(":%d", *port))
}
