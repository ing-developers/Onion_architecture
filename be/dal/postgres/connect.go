package postgres

import (
	"github.com/ing-developers/clean-architecture/be/dto"
	"github.com/ing-developers/go-dal"
	_ "github.com/lib/pq"
	"log"
)

func Connect() (*go_dal.DAL, error) {
	engine, err := go_dal.Connect(dto.ConfigServerDB)
	if err != nil {
		log.Println(dto.I18N.ErrConnectDB, err)
		return nil, err
	}
	if engine.Connected {
		return engine, err
	} else {
		log.Println(dto.I18N.ErrConnectDB, err)
		return nil, err
	}
}

