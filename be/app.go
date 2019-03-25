package main

import (
	"fmt"
	"github.com/ing-developers/clean-architecture/be/controllers"
	"github.com/ing-developers/clean-architecture/be/dto"
	"log"
	"net/http"
)

func main() {
	// Information server and app
	log.Println("Iniciando servidor...")
	log.Println("Direccion:", dto.Server.Domain)
	log.Println("Puerto:", dto.Server.Port)
	log.Println("...Informacion de aplicacion...")
	log.Println("App:", dto.App.Name)
	log.Println("Lenguaje:", dto.App.Lang)
	log.Println("Version:", dto.App.Version)
	// Load routes api
	loadRoutesApi()
	// Inicialization of server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", dto.Server.Domain, dto.Server.Port), dto.Router))
}

func loadRoutesApi()  {
	path := "product"
	dto.Router.Post(dto.App.PrefixApi+path, controllers.Product{}.Create)
}
