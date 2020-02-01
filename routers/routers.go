package routers

import (
	"github.com/gorilla/mux"
	"github.com/terrymight/go_simple_api_db/controllers"
)

// BasicRoutes GerPeople list
var BasicRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/people", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/api/person/{id}", controllers.GetPerson).Methods("GET")
	router.HandleFunc("/api/person/{id}", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/api/person/{id}", controllers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/api/person/{id}", controllers.DeletePerson ).Methods("DELETE")

}