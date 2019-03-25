package dto

// server struct for load config of server
type server struct {
	Domain     string `json:"domain"`
	Port       string `json:"port"`
	Production bool   `json:"production"`
}
