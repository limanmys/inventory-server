package jobs

import (
	"errors"
	"os"

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
	db := database.Connection().Model(&entities.Job{})

	// Apply search, if exists
	if c.Query("search") != "" {
		search.Search(c.Query("search"), db)
	}

	// Get data
	var jobs []entities.Job
	page, err := paginator.New(db, c).Paginate(&jobs)
	if err != nil {
		return err
	}

	return c.JSON(page)
}

// Download report
func Download(c *fiber.Ctx) error {
	// Check uuid validity
	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	// Download report
	var job entities.Job
	if err := database.Connection().
		Model(&job).Where("id = ?", uuid).First(&job).Error; err != nil {
		return err
	}

	// Check job has any path data
	if job.Path == "" {
		return errors.New("report does not exists, please set status & message")
	}

	// Check is path exists
	if _, err := os.Stat(job.Path); err != nil {
		if os.IsNotExist(err) {
			return errors.New("report file does not exists")
		} else {
			return err
		}
	}
	return c.SendFile(job.Path)
}
