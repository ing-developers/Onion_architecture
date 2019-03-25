package maria

import (
	"github.com/ing-developers/clean-architecture/be/dal/entities"
	"github.com/ing-developers/go-dal"
)

// Product operation of entity to database
type Product struct{}

// Create operation for create register in database
func (Product) Create(p *entities.Product) error {
	dal, err := Connect()
	if err != nil {
		return err
	}
	data, err := dal.GetRowsQuery("CALL product_create(?)", p.Name)
	if err != nil {
		return err
	}
	return go_dal.ToStruct(data[0], p)
}

