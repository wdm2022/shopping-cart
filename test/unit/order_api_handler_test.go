package unit

import (
	"net/http/httptest"
	handlers "shopping-cart/pkg/api/handlers"
	"strconv"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderCorrect(t *testing.T) {
	tests := []struct {
		description  string
		orderIndex   int
		expectedCode int
	}{
		{
			description:  "Create a new order",
			orderIndex:   1,
			expectedCode: 200,
		},
		{
			description:  "Create a new order (negative number)",
			orderIndex:   -1,
			expectedCode: 200,
		},
	}
	app := fiber.New()
	app.Post("/create/:userId", handlers.CreateOrder)
	for _, test := range tests {
		request_str := strings.Join([]string{"/create/", strconv.Itoa(test.orderIndex)}, "")
		req := httptest.NewRequest("POST", request_str, nil)
		resp, _ := app.Test(req, -1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestCreateOrderFails(t *testing.T) {
	tests := []struct {
		description     string
		orderIdentifier string
		expectedCode    int
	}{
		{
			description:     "Create a new order (weird identifier)",
			orderIdentifier: "dffsd",
			expectedCode:    400,
		},
		{
			description:     "Create a new order (empty)",
			orderIdentifier: "",
			expectedCode:    404,
		},
	}
	app := fiber.New()
	app.Post("/create/:userId", handlers.CreateOrder)
	for _, test := range tests {
		request_str := strings.Join([]string{"/create/", test.orderIdentifier}, "")
		req := httptest.NewRequest("POST", request_str, nil)
		resp, _ := app.Test(req, -1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
