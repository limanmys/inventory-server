package entities

import (
	"github.com/google/uuid"
	"github.com/limanmys/inventory-server/internal/database"
)

type DiscoveryStatus string

var (
	DiscoveryStatusPending    DiscoveryStatus = "pending"
	DiscoveryStatusInProgress DiscoveryStatus = "in_progress"
	DiscoveryStatusDone       DiscoveryStatus = "done"
	DiscoveryStatusError      DiscoveryStatus = "error"
)

type Discovery struct {
	Base
	IPRange   string     `json:"ip_range" validate:"required"`
	ProfileID *uuid.UUID `json:"profile_id" validate:"required"`
	Profile   *Profile   `json:"profile"`

	DiscoveryStatus DiscoveryStatus `json:"discovery_status" gorm:"default:pending"`
	Message         string          `json:"message"`
}

func (d *Discovery) UpdateStatus(status DiscoveryStatus, message string) {
	d.DiscoveryStatus = status
	d.Message = message
	database.Connection().Model(d).Save(d)
}
