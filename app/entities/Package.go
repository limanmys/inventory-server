package entities

import "github.com/google/uuid"

type Package struct {
	Base
	Name                 string              `json:"name"`
	Version              string              `json:"version"`
	Vendor               string              `json:"vendor"`
	Assets               []*Asset            `json:"assets" gorm:"many2many:asset_packages"`
	AlternativePackageID *uuid.UUID          `json:"alternative_package_id"`
	AlternativePackage   *AlternativePackage `json:"alternative_package"`
}
