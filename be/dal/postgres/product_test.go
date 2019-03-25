package postgres

import (
	"fmt"
	"github.com/ing-developers/clean-architecture/be/dal/entities"
	"log"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	p := entities.Product{
		Name: "Salmon",
	}
	err := Product{}.Create(&p)
	if err != nil {
		log.Println("Error al crear el producto: ", err)
	} else {
		fmt.Println(p)
	}
}
