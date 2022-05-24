package routes

import (
	"log"
	"net/http"

	"github.com/brfyamada/golang-rest-api/controllers"
	"github.com/brfyamada/golang-rest-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.FindAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.FindPersonalityById).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.UpdatePersonality).Methods("Put")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonality).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
