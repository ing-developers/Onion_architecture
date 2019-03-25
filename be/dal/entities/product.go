package entities

// Product map of entity from database
type Product struct {
	ID   int64  `json:"id,string"`
	Name string `json:"name"`
}
