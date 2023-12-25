package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limanmys/inventory-server/app/controllers/discoveries"
	"github.com/limanmys/inventory-server/app/controllers/profiles"
)

func Routes(app *fiber.App) {
	// Profile routes
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

	// Discovery routes
	discoveryGroup := app.Group("/discoveries")
	{
		// Create record
		discoveryGroup.Post("/", discoveries.Create)
		// Run discovery
		discoveryGroup.Post("/:id", discoveries.Run)
		// Index records
		discoveryGroup.Get("/", discoveries.Index)
		// Delete record
		discoveryGroup.Delete("/:id", discoveries.Delete)
	}
}
