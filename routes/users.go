package routes

import (
	"literature/handlers"
	"literature/pkg/mysql"
	"literature/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {

	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
}
