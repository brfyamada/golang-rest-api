package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/brfyamada/golang-rest-api/database"
	"github.com/brfyamada/golang-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func FindAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func FindPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)
	if personalidade.Id == 0 {
		log.Printf("Registro não encontrado")
		w.WriteHeader(404)
		return
	}
	json.NewEncoder(w).Encode(personalidade)

}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {

	var personalidade models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&personalidade)
	if err == io.EOF || err != nil {
		log.Printf("Requisição invalida")
		w.WriteHeader(404)
		return
	}
	database.DB.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)

}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&personalidade)
	if err == io.EOF || err != nil {
		log.Printf("Requisição invalida")
		w.WriteHeader(404)
		return
	}

	var personalidadeFromDB models.Personalidade
	database.DB.First(&personalidadeFromDB, id)
	if personalidadeFromDB.Id == 0 {
		log.Printf("Registro não encontrado")
		w.WriteHeader(404)
		return
	}
	personalidade.Id = personalidadeFromDB.Id

	database.DB.Save(&personalidade)

}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	if personalidade.Id == 0 {
		log.Printf("Registro não encontrado")
		w.WriteHeader(404)
		return
	}

	database.DB.Delete(&personalidade)

}
