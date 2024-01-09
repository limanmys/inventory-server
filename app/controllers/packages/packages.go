package packages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/database"
	"github.com/limanmys/inventory-server/internal/paginator"
	"github.com/limanmys/inventory-server/internal/search"
	"github.com/limanmys/inventory-server/pkg/jobs"
	"github.com/limanmys/inventory-server/pkg/reporter"
	"gorm.io/gorm/clause"
)

type PackageWithAssetCount struct {
	entities.Package
	Count int `json:"count"`
}

// Index, returns asset's packages
func Index(c *fiber.Ctx) error {
	// Build sql query
	sub_query := database.Connection().
		Model(&entities.Package{}).
		Select("packages.name", "count(*)", "null as updated_at", "null as deleted_at", "alternative_package_id").
		Joins("inner join asset_packages ap on ap.package_id = packages.id").
		Joins("inner join assets on assets.id = ap.asset_id").
		Group("packages.name").Group("alternative_package_id").Order("count desc")

	db := database.Connection().Preload(clause.Associations).Table("(?) as t1", sub_query)

	// Apply search, if exists
	if c.Query("search") != "" {
		search.Search(c.Query("search"), db)
	}

	// Get data
	var packages []PackageWithAssetCount
	page, err := paginator.New(db, c).Paginate(&packages)
	if err != nil {
		return err
	}

	return c.JSON(page)
}

func Report(c *fiber.Ctx) error {
	// Create new report job
	job, err := jobs.NewJob(entities.FileType(c.Params("file_type")))
	if err != nil {
		return err
	}

	// Build query
	db := database.Connection().
		Model(&entities.Package{}).
		Select("packages.name", "count(*)").
		Joins("inner join asset_packages ap on ap.package_id = packages.id").
		Joins("inner join assets on assets.id = ap.asset_id").
		Group("packages.name").Order("count desc")

	// Create report as go routine
	go reporter.CreatePackageReport(job, db, []string{"name", "count"})

	return c.JSON(fiber.Map{"id": job.ID.String()})
}
