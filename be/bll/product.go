package bll

import (
	"github.com/pkg/errors"
	"log"
	"github.com/ing-developers/clean-architecture/be/dal/entities"
	"github.com/ing-developers/clean-architecture/be/dal/interfaces"
	"github.com/ing-developers/clean-architecture/be/dal/maria"
	"github.com/ing-developers/clean-architecture/be/dal/postgres"
	"github.com/ing-developers/clean-architecture/be/dto"
)

type Product struct{}

func (Product) getContext() interfaces.Product {
	switch dto.ConfigServerDB.Engine {
	case "mysql":
		return &maria.Product{}
	case "postgres":
		return &postgres.Product{}
	default:
		log.Println(dto.I18N.ErrNoImplementEngine)
		return nil
	}
}

func (p Product) Create(product *entities.Product) error {
	dal := p.getContext()
	if dal == nil {
		return errors.New(dto.I18N.ErrNoImplementEngine)
	}
	return dal.Create(product)
}
