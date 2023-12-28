package profiles

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/database"
	"github.com/limanmys/inventory-server/internal/paginator"
	"github.com/limanmys/inventory-server/internal/search"
	"github.com/limanmys/inventory-server/internal/validation"
	"github.com/limanmys/inventory-server/pkg/aes"
	"gorm.io/gorm/clause"
)

// Create, Creates a new record
func Create(c *fiber.Ctx) error {
	// Parse body
	var payload entities.Profile
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	// Validate payload
	err := validation.Validate(payload)
	if err != nil {
		return err
	}

	// Encrypt profile data
	payload, err = aes.EncryptProfile(payload)
	if err != nil {
		return err
	}

	// Create record on database
	if err := database.Connection().Clauses(clause.Returning{}).Create(&payload).Error; err != nil {
		return err
	}

	payload.Username = ""
	payload.Password = ""

	return c.JSON(payload)
}

// Index, Lists all records
func Index(c *fiber.Ctx) error {
	// Set query
	db := database.Connection().Model(&entities.Profile{}).
		Preload(clause.Associations).Omit("username", "password")

	// Apply search, if exists
	if c.Query("search") != "" {
		search.Search(c.Query("search"), db)
	}

	// Get data
	var profiles []entities.Profile
	page, err := paginator.New(db, c).Paginate(&profiles)
	if err != nil {
		return err
	}

	return c.JSON(page)
}

// Update, Updates existing record
func Update(c *fiber.Ctx) error {
	// Parse body
	var payload entities.Profile
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	// Check uuid validity
	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	// Encrypt profile data
	payload, err = aes.EncryptProfile(payload)
	if err != nil {
		return err
	}

	// Update record
	if err := database.Connection().
		Model(&entities.Profile{}).
		Where("id = ?", uuid).Updates(&payload).Error; err != nil {
		return err
	}

	payload.Username = ""
	payload.Password = ""

	return c.JSON(payload)
}

// Delete, Deletes existing record
func Delete(c *fiber.Ctx) error {
	// Check uuid validity
	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	// Delete record
	if err := database.Connection().Unscoped().
		Where("id = ?", uuid).Delete(&entities.Profile{}).Error; err != nil {
		return err
	}

	return c.JSON("Record deleted successfully.")
}
