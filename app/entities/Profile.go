package entities

type Profile struct {
	Base
	Username    string       `json:"username" validate:"required"`
	Password    string       `json:"password" validate:"required"`
	Discoveries []*Discovery `json:"discoveries"`
}
