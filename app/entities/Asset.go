package entities

import "github.com/google/uuid"

type Asset struct {
	Base
	Hostname     string `json:"hostname"`
	Address      string `json:"address"`
	SerialNumber string `json:"serial_number"`
	BiosVersion  string `json:"bios_version"`
	Vendor       string `json:"vendor"`
	Model        string `json:"model"`

	DiscoveryID *uuid.UUID `json:"discovery_id"`
	Discovery   *Discovery `json:"discovery"`
	Packages    []*Package `json:"packages" gorm:"many2many:asset_packages"`
}
