package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limanmys/inventory-server/app/controllers/profiles"
)

func Routes(app *fiber.App) {
	profileGroup := app.Group("/profiles")
	{
		// Create record
		profileGroup.Post("/", profiles.Create)
		// Index records
		profileGroup.Get("/", profiles.Index)
		// Update record
		profileGroup.Patch("/:id", profiles.Update)
		// Delete record
		profileGroup.Delete("/:id", profiles.Delete)
	}
}
