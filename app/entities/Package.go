package entities

type Package struct {
	Base
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Vendor  string   `json:"vendor"`
	Count   int64    `json:"count" gorm:"-"`
	Assets  []*Asset `json:"assets" gorm:"many2many:asset_packages"`
}
