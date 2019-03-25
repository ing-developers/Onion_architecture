package dto

// app struct for load information of app
type app struct {
	Name      string `json:"name"`
	Lang      string `json:"lang"`
	Version   string `json:"version"`
	PrefixApi string `json:"prefix_api"`
	Debug     bool   `json:"debug"`
}
