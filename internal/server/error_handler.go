package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limanmys/inventory-server/internal/validation"
)

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*validation.Errors); ok {
		code = fiber.StatusUnprocessableEntity
		return c.Status(code).JSON(e)
	}
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	var message interface{}
	if code == fiber.StatusOK {
		message = struct {
			Message string `json:"message"`
		}{err.Error()}
	} else {
		message = struct {
			Error string `json:"error"`
		}{err.Error()}
	}
	return c.Status(code).JSON(message)
}
