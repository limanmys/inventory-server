package discoveries

import (
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/constants"
	"github.com/limanmys/inventory-server/internal/database"
	"github.com/limanmys/inventory-server/internal/paginator"
	"github.com/limanmys/inventory-server/internal/search"
	"github.com/limanmys/inventory-server/internal/validation"
	"github.com/limanmys/inventory-server/pkg/discovery"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Create, Creates a new record
func Create(c *fiber.Ctx) error {
	// Parse body
	var payload entities.Discovery
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	// Validate payload
	err := validation.Validate(payload)
	if err != nil {
		return err
	}

	// Set discovery status
	payload.Status = entities.StatusPending
	payload.Message = "Discovery pending."

	// Create record on database
	if err := database.Connection().Clauses(clause.Returning{}).Create(&payload).Error; err != nil {
		return err
	}

	// Set run id
	runID := uuid.New().String()

	// Create discovery logs record
	if err := database.Connection().Create(&entities.DiscoveryLogs{
		DiscoveryID: payload.ID,
		Filename:    runID,
	}).Error; err != nil {
		return err
	}

	// Start discovery
	go discovery.Start(payload, runID)

	return c.JSON(payload)
}

// Run, runs a discovery
func Run(c *fiber.Ctx) error {
	// Check uuid validity
	uid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	// Get discovery
	var discoveryObject entities.Discovery
	if err := database.Connection().Model(&entities.Discovery{}).Where("id = ?", uid).First(&discoveryObject).Error; err != nil {
		return err
	}

	// Set run id
	runID := uuid.New().String()

	// Create discovery logs record
	if err := database.Connection().Create(&entities.DiscoveryLogs{
		DiscoveryID: discoveryObject.ID,
		Filename:    runID,
	}).Error; err != nil {
		return err
	}

	// Start discovery
	go discovery.Start(discoveryObject, runID)

	return c.JSON("Discovery started successfully.")
}

// Index, Lists all records
func Index(c *fiber.Ctx) error {
	// Set query
	db := database.Connection().Model(&entities.Discovery{}).
		Preload("Profile", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("username", "password")
		})

	// Apply search, if exists
	if c.Query("search") != "" {
		search.Search(c.Query("search"), db)
	}

	// Get data
	var discoveries []entities.Discovery
	page, err := paginator.New(db, c).Paginate(&discoveries)
	if err != nil {
		return err
	}

	return c.JSON(page)
}

// Delete, Deletes existing record
func Delete(c *fiber.Ctx) error {
	// Check uuid validity
	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	// Delete record
	if err := database.Connection().
		Where("id = ?", uuid).Delete(&entities.Discovery{}).Error; err != nil {
		return err
	}

	return c.JSON("Record deleted successfully.")
}

// ReadLog, Reads a latest log file
func ReadLatestLog(c *fiber.Ctx) error {
	// Check uuid validity
	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	// Get logs
	var log entities.DiscoveryLogs
	if err := database.Connection().Model(&entities.DiscoveryLogs{}).
		Where("discovery_id = ?", uuid).Order("updated_at DESC").First(&log).Error; err != nil {
		return err
	}

	// Read content
	content, err := os.ReadFile(path.Join(constants.DEFAULT_LOG_PATH, log.Filename))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"content": string(content)})
}
