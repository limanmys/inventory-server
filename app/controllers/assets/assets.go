package assets

import (
	"github.com/gofiber/fiber/v2"
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
