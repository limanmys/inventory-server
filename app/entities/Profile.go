package entities

type Profile struct {
	Base
	Name        string       `json:"name" gorm:"unique" validate:"required"`
	Username    string       `json:"username" validate:"required"`
	Password    string       `json:"password" validate:"required"`
	Discoveries []*Discovery `json:"discoveries"`
}
