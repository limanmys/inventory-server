package metrics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/database"
	"github.com/limanmys/inventory-server/pkg/counter"
	"gorm.io/gorm"
)

// Asset count, returns total asset count
func AssetCount(c *fiber.Ctx) error {
	// Get count
	response, err := counter.Get("assets")
	if err != nil {
		return err
	}

	return c.JSON(response)
}

// Discovery count, returns total discovery count
func DiscoveryCount(c *fiber.Ctx) error {
	// Get count
	response, err := counter.Get("discoveries")
	if err != nil {
		return err
	}

	return c.JSON(response)
}

// Package count, returns total package count
func PackageCount(c *fiber.Ctx) error {
	// Get count
	response, err := counter.Get("packages")
	if err != nil {
		return err
	}

	return c.JSON(response)
}

// LatestDiscoveryTime, returns latest discovery time
func LatestDiscoveryTime(c *fiber.Ctx) error {
	var item entities.Discovery
	err := database.Connection().Model(&entities.Discovery{}).First(&item).Error
	if err != gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{"time": item.UpdatedAt})
	}

	return c.JSON(fiber.Map{"time": 0})
}

// AddedAssetsAsTimeseries, return last 7 days added asset counts day by day
func AddedAssetsAsTimeseries(c *fiber.Ctx) error {
	// Set result
	var result = make([]map[string]interface{}, 0)
	// Run query
	if err := database.Connection().Raw(`
SELECT
	time_interval.date,
	count(assets.id)
FROM
(
SELECT
	to_char(date_trunc('day',
	(current_date - offs)),
	'YYYY-MM-DD') as date
FROM
	generate_series(0,
	6,
	1) 
	as offs
	) time_interval
LEFT OUTER JOIN assets ON 
	(time_interval.date = to_char(date_trunc('day',
	assets.created_at),
	'YYYY-MM-DD'))
GROUP BY
	time_interval.date
ORDER BY
	time_interval.date`).Scan(&result).Error; err != nil {
		return err
	}

	return c.JSON(result)
}

// VendorCounts, returns asset count as vendor by vendor
func VendorCounts(c *fiber.Ctx) error {
	// Set result
	var result = make([]map[string]interface{}, 0)

	// Run query
	if err := database.Connection().Model(&entities.Asset{}).
		Select("vendor", "count(*)").Where("deleted_at is null").
		Group("vendor").Find(&result).Error; err != nil {
		return err
	}

	return c.JSON(result)
}

// MostUsedPackages, returns most used 5 packages
func MostUsedPackages(c *fiber.Ctx) error {
	// Set result
	var result = make([]map[string]interface{}, 0)

	if err := database.Connection().
		Model(&entities.Package{}).
		Select("packages.name", "count(*)").
		Joins("inner join asset_packages ap on ap.package_id = packages.id").
		Joins("inner join assets on assets.id = ap.asset_id").
		Group("packages.name").Order("count desc").Limit(5).Find(&result).
		Error; err != nil {
		return err
	}

	return c.JSON(&result)
}
