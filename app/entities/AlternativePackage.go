package entities

type AlternativePackage struct {
	Base
	Name        string     `json:"name"`
	URL         string     `json:"url"`
	PackageName string     `json:"package_name"`
	Packages    []*Package `json:"packages"`
}
