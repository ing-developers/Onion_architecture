package postgres

import (
	"github.com/ing-developers/clean-architecture/be/dal/entities"
	"strconv"
)

// Product operation of entity to database
type Product struct{}

// Create operation for create register in database
func (Product) Create(p *entities.Product) error {
	dal, err := Connect()
	if err != nil {
		return err
	}
	data, err := dal.GetRowsQuery("INSERT INTO public.products(name) VALUES ($1) RETURNING id", p.Name)
	if err == nil {
		p.ID, err = strconv.ParseInt(data[0]["id"], 10, 64)
		return err
	}
	return err
}
