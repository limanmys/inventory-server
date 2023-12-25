package entities

import "github.com/google/uuid"

type Asset struct {
	Base
	Hostname     string `json:"hostname"`
	Address      string `json:"address"`
	SerialNumber string `json:"serial_number"`
	Vendor       string `json:"vendor"`

	DiscoveryID *uuid.UUID `json:"discovery_id"`
	Discovery   *Discovery `json:"discovery"`
}
