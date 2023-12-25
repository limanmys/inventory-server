package entities

type Package struct {
	Base
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Vendor  string   `json:"vendor"`
	Assets  []*Asset `json:"assets" gorm:"many2many:asset_packages"`
}
