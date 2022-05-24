package main

import (
	"fmt"

	"github.com/brfyamada/golang-rest-api/database"
	"github.com/brfyamada/golang-rest-api/routes"
)

func main() {

	/*
		models.Personalidades = []models.Personalidade{
			{Id: 1, Nome: "Getúlio Vargas", Historia: "Getúlio Dornelles Vargas foi um advogado e político brasileiro, líder da Revolução de 1930, que pôs fim à República Velha, depondo seu 13.º e último presidente, Washington Luís, e impedindo a posse do presidente eleito em 1.º de março de 1930, Júlio Prestes"},
			{Id: 2, Nome: "José Fortuna", Historia: "foi um cantor, compositor, autor teatral e ator brasileiro"},
		}
	*/

	database.ConnectDB()
	fmt.Println("Iniciando Servidor GO")
	routes.HandleRequest()
}
