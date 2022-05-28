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
	orders.Delete("/create/remove/:orderId", handlers.DeleteOrder)
	orders.Get("/create/find/:orderId", handlers.GetOrder)
	orders.Get("/addItem/:orderId/:itemId", handlers.AddItem)
	orders.Delete("/removeItem/:orderId/:itemId", handlers.DeleteItem)
	// TODO: uncomment after implementing handlers
	//orders.Post("/checkout/:orderId")
}

func bindStockApi(app *fiber.App) {
	// TODO: uncomment after implementing handlers
	//stock := app.Group("/stock")

	//stock.Get("/find/:itemId")
	//stock.Post("/subtract/:itemId/:amount")
	//stock.Post("/add/:itemId/:amount")
	//stock.Post("/item/create/:price")
}

func bindPaymentService(app *fiber.App) {
	// TODO: uncomment after implementing handlers
	//payment := app.Group("/payment")

	//payment.Post("/pay/:userId/:orderId/:amount")
	//payment.Post("/cancel/:userId/:orderId")
	//payment.Get("/status/:userId/:orderId")
	//payment.Post("/add_funds/:userId/:amount")
	//payment.Post("/create_user")
	//payment.Get("/find_user/:user_id")
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
	bindPaymentService(app)

	return app.Listen(fmt.Sprintf(":%d", *port))
}
