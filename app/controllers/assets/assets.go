package assets

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/database"
	"github.com/limanmys/inventory-server/internal/paginator"
	"github.com/limanmys/inventory-server/internal/search"
)

// Index, Lists all records
func Index(c *fiber.Ctx) error {
	// Set query
	sub_query := database.Connection().
		Model(&entities.Asset{}).
		Select("assets.*", "count(*) as package_count").
		Joins("left join asset_packages ap on ap.asset_id = assets.id").
		Joins("left join packages on ap.package_id = packages.id").
		Group("assets.id").Order("package_count desc")

	db := database.Connection().Table("(?) as t1", sub_query)

	// Apply search, if exists
	if c.Query("search") != "" {
		search.Search(c.Query("search"), db)
	}

	// Get data
	var assets []map[string]interface{}
	page, err := paginator.New(db, c).Paginate(&assets)
	if err != nil {
		return err
	}

	return c.JSON(page)
}

// Show, gets a single asset
func Show(c *fiber.Ctx) error {
	// Check uuid validity
	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	// Set result
	var asset entities.Asset

	// Set query
	if err := database.Connection().
		Model(&asset).First(&asset, "id = ?", uuid).Error; err != nil {
		return err
	}

	return c.JSON(asset)
}

// AssetPackages, returns asset's packages
func AssetPackages(c *fiber.Ctx) error {
	// Check uuid validity
	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	db := database.Connection().Model(&entities.Package{}).
		Joins("left join asset_packages on asset_packages.package_id = packages.id").
		Where("asset_id = ?", uuid)

	// Apply search, if exists
	if c.Query("search") != "" {
		search.Search(c.Query("search"), db)
	}

	// Get data
	var packages []entities.Package
	page, err := paginator.New(db, c).Paginate(&packages)
	if err != nil {
		return err
	}

	return c.JSON(page)
}
