package migrations

import (
	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/database"
)

func Migrate() error {
	if err := database.Connection().AutoMigrate(&entities.Discovery{}); err != nil {
		return err
	}
	if err := database.Connection().AutoMigrate(&entities.Asset{}); err != nil {
		return err
	}
	if err := database.Connection().AutoMigrate(&entities.Profile{}); err != nil {
		return err
	}
	if err := database.Connection().AutoMigrate(&entities.Package{}); err != nil {
		return err
	}
	return nil
}
