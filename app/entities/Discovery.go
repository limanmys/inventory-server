package entities

import "github.com/google/uuid"

type Discovery struct {
	Base
	IPRange   string     `json:"ip_range" validate:"required"`
	ProfileID *uuid.UUID `json:"profile_id" validate:"required"`
	Profile   *Profile   `json:"profile"`
}
