package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId")
		orderId := ctx.Params("orderId")
		return ctx.SendString("Order: " + orderId + " from " + userId)
	})

	request := httptest.NewRequest("GET", "/users/100/orders/9812", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "Order: 9812 from 100", string(bytes))

}
