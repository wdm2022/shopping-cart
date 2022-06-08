package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"shopping-cart/pkg/api/handlers"
)

func bindOrdersApi(app *fiber.App) {
	orders := app.Group("/orders")

	orders.Post("/create/:userId", handlers.CreateOrder)
	orders.Delete("/remove/:orderId", handlers.DeleteOrder)
	orders.Get("/find/:orderId", handlers.GetOrder)
	orders.Post("/addItem/:orderId/:itemId", handlers.AddItem)
	orders.Delete("/removeItem/:orderId/:itemId", handlers.DeleteItem)
	orders.Post("/checkout/:orderId", handlers.Checkout)
}

func bindStockApi(app *fiber.App) {
	stock := app.Group("/stock")

	stock.Get("/find/:itemId", handlers.GetStock)
	stock.Post("/subtract/:itemId/:amount", handlers.SubtractStock)
	stock.Post("/add/:itemId/:amount", handlers.AddStock)
	stock.Post("/item/create/:price", handlers.CreateItem)
}

func bindPaymentApi(app *fiber.App) {
	payment := app.Group("/payment")

	payment.Post("/pay/:userId/:orderId/:amount", handlers.PlaceOrderPayment)
	payment.Post("/cancel/:userId/:orderId", handlers.CancelOrderPayment)
	payment.Get("/status/:userId/:orderId", handlers.GetOrderPayment)
	payment.Post("/add_funds/:userId/:amount", handlers.AddFunds)
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