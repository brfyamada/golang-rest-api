package main

import (
	"fmt"

	"github.com/brfyamada/golang-rest-api/models"
	"github.com/brfyamada/golang-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "Getúlio Vargas", Historia: "Getúlio Dornelles Vargas foi um advogado e político brasileiro, líder da Revolução de 1930, que pôs fim à República Velha, depondo seu 13.º e último presidente, Washington Luís, e impedindo a posse do presidente eleito em 1.º de março de 1930, Júlio Prestes"},
		{Nome: "José Fortuna", Historia: "foi um cantor, compositor, autor teatral e ator brasileiro"},
	}

	fmt.Println("Iniciando Servidor GO")
	routes.HandleRequest()
}
