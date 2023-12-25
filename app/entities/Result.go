package entities

type Result struct {
	Error  string      `json:"error"`
	Output interface{} `json:"result"`
}

type Arguments struct {
	IPRange  string `json:"ip_range"`
	Username string `json:"string"`
	Password string `json:"password"`
}
