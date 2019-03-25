package interfaces

import "github.com/ing-developers/clean-architecture/be/dal/entities"

type Product interface {
	Create(p *entities.Product) error
}
