package entities

import (
	"github.com/google/uuid"
	"github.com/limanmys/inventory-server/internal/database"
)

type Status string

var (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
	StatusError      Status = "error"
)

type Discovery struct {
	Base
	IPRange   string     `json:"ip_range" validate:"required"`
	ProfileID *uuid.UUID `json:"profile_id" validate:"required"`
	Profile   *Profile   `json:"profile"`

	Status  Status `json:"discovery_status" gorm:"default:pending"`
	Message string `json:"message"`
}

func (d *Discovery) UpdateStatus(status Status, message string) {
	d.Status = status
	d.Message = message
	database.Connection().Model(d).Save(d)
}

type DiscoveryLogs struct {
	Base
	DiscoveryID *uuid.UUID `json:"discovery_id"`
	Discovery   *Discovery `json:"discovery"`
	Filename    string     `json:"filename"`
}
