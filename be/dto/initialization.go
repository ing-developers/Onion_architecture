package dto

import (
	"fmt"
	"github.com/husobee/vestigo"
	"github.com/ing-developers/go-dal"
	"github.com/ing-developers/go-tools"
	"log"
)

// ConfigServerDB var for load config connection to server database
var ConfigServerDB go_dal.ServerDB

// App var for load config of app
var App app

// Server var for load config of server
var Server server

// I18N var for load messages of app
var I18N i18n

// Router var for routes / endpoints of app
var Router *vestigo.Router

func init() {
	// Config server database
	err := tools.Decode("./serverDB.json", &ConfigServerDB)
	if err != nil {
		log.Println("Error al leer el archivo de conexion con el servidor de base de datos: ", err)
	}
	// Config App
	err = tools.Decode("./app.json", &App)
	if err != nil {
		log.Println("Error al leer el archivo de configuracion de la aplicacion: ", err)
	} else {
		// Config i18n
		err = tools.Decode(fmt.Sprintf("./i18n/%s.json", App.Lang), &I18N)
		if err != nil {
			log.Println("Error al leer el archivo de idioma de la aplicacion: ", err)
		}
	}
	// Router App
	Router = vestigo.NewRouter()
	Router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:      []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{"content-type"},
		ExposeHeaders:    []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	})

	// Config server
	err = tools.Decode("./server.json", &Server)
	if err != nil {
		log.Println("Error al leer el archivo de configuracion del servidor: ", err)
	}
}
