package entities

import "github.com/google/uuid"

type Asset struct {
	Base
	Hostname     string `json:"hostname"`
	Address      string `json:"address"`
	SerialNumber string `json:"serial_number"`
	Vendor       string `json:"vendor"`
	Model        string `json:"model"`
	// Packages     []Package `json:"packages"`

	DiscoveryID *uuid.UUID `json:"discovery_id"`
	Discovery   *Discovery `json:"discovery"`
}
