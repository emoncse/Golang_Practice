package routes

import (
	"first_project/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define routes similar to Django urls
	router.HandleFunc("/", controllers.HomeController).Methods("GET")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")

	return router
}
