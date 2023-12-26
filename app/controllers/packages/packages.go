package packages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/database"
	"github.com/limanmys/inventory-server/internal/paginator"
	"github.com/limanmys/inventory-server/internal/search"
)

// Index, returns asset's packages
func Index(c *fiber.Ctx) error {
	// Build sql query
	sub_query := database.Connection().
		Model(&entities.Package{}).
		Select("packages.name", "count(*)", "null as updated_at", "null as deleted_at").
		Joins("inner join asset_packages ap on ap.package_id = packages.id").
		Joins("inner join assets on assets.id = ap.asset_id").
		Group("packages.name")

	db := database.Connection().Table("(?) as t1", sub_query)

	// Apply search, if exists
	if c.Query("search") != "" {
		search.Search(c.Query("search"), db)
	}

	// Get data
	var packages []map[string]interface{}
	page, err := paginator.New(db, c).Paginate(&packages)
	if err != nil {
		return err
	}

	return c.JSON(page)
}
