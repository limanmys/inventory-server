package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limanmys/inventory-server/app/controllers/assets"
	"github.com/limanmys/inventory-server/app/controllers/discoveries"
	"github.com/limanmys/inventory-server/app/controllers/jobs"
	"github.com/limanmys/inventory-server/app/controllers/packages"
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

	// Asset routes
	assetGroup := app.Group("/assets")
	{
		// Index records
		assetGroup.Get("/", assets.Index)
		// Asset packages
		assetGroup.Get("/packages/:id", assets.AssetPackages)
	}

	// Package routes
	packageGroup := app.Group("/packages")
	{
		// Index records
		packageGroup.Get("/", packages.Index)
		// Create report
		packageGroup.Get("/report/:file_type", packages.Report)
	}

	// Job routes
	jobGroup := app.Group("/jobs")
	{
		// Index records
		jobGroup.Get("/", jobs.Index)
		// Download report
		jobGroup.Get("/:id", jobs.Download)
	}
}
