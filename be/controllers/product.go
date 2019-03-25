package controllers

import (
	"fmt"
	"github.com/ing-developers/clean-architecture/be/bll"
	"github.com/ing-developers/clean-architecture/be/dal/entities"
	"github.com/ing-developers/clean-architecture/be/dto"
	"github.com/ing-developers/go-tools"
	"net/http"
)

type Product struct {}

func (p Product) Create(w http.ResponseWriter, r *http.Request) {
	product := entities.Product{}
	err := tools.DecodeRequest(r, &product)
	fmt.Println(product)
	if err != nil {
		fmt.Println(err)
		tools.Responder(w, false, err, dto.I18N.ErrBadEntity, nil, true, dto.App.Debug)
		return
	}
	if p.ValidateEntity(&product, false) {
		err = bll.Product{}.Create(&product)
		fmt.Println(err)
		if err != nil {
			tools.Responder(w, false, err, dto.I18N.ErrCreateRegisterDB, nil, true, dto.App.Debug)
		} else {
			tools.Responder(w, true, nil, dto.I18N.CrrCreateRegisterDB, product, true, dto.App.Debug)
		}
	} else {
		tools.Responder(w, false, nil, dto.I18N.ErrBadEntity, nil, true, dto.App.Debug)
	}
}

func (Product) ValidateEntity(entity *entities.Product, isUpdate bool) bool {
	var validation = tools.ValidarLongitud(entity.Name, 5, 50)
	if isUpdate {
		validation = validation && entity.ID > 0
	}
	return validation
}


