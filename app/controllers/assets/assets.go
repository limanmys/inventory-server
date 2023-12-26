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
	db := database.Connection().Model(&entities.Asset{})

	// Apply search, if exists
	if c.Query("search") != "" {
		search.Search(c.Query("search"), db)
	}

	// Get data
	var assets []entities.Asset
	page, err := paginator.New(db, c).Paginate(&assets)
	if err != nil {
		return err
	}

	return c.JSON(page)
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

	// Get data
	var packages []entities.Package
	page, err := paginator.New(db, c).Paginate(&packages)
	if err != nil {
		return err
	}

	return c.JSON(page)
}
