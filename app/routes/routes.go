package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limanmys/inventory-server/app/controllers/assets"
	"github.com/limanmys/inventory-server/app/controllers/discoveries"
	"github.com/limanmys/inventory-server/app/controllers/jobs"
	"github.com/limanmys/inventory-server/app/controllers/metrics"
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
		// Read latest log
		discoveryGroup.Get("/logs/:id", discoveries.ReadLatestLog)
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
		// Show record
		assetGroup.Get("/:id", assets.Show)
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

	// Metric routes
	metricGroup := app.Group("/metrics")
	{
		// Asset Count
		metricGroup.Get("/asset_count", metrics.AssetCount)
		// Discovery Count
		metricGroup.Get("/discovery_count", metrics.DiscoveryCount)
		// Package Count
		metricGroup.Get("/package_count", metrics.PackageCount)
		// Latest Discovery Time
		metricGroup.Get("/latest_discovery_time", metrics.LatestDiscoveryTime)
		// Added Assets As Timeseries
		metricGroup.Get("/added_assets", metrics.AddedAssetsAsTimeseries)
		// Vendor Counts
		metricGroup.Get("/vendor_counts", metrics.VendorCounts)
		// Most Used Packages
		metricGroup.Get("/most_used_packages", metrics.MostUsedPackages)
	}
}
