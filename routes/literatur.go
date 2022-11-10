package routes

import (
	"literature/handlers"
	"literature/pkg/middleware"
	"literature/pkg/mysql"
	"literature/repositories"

	"github.com/gorilla/mux"
)

func LiteraturRoutes(r *mux.Router) {
	LiteraturRepository := repositories.RepositoryLiteratur(mysql.DB)

	h := handlers.HandlerLiteratur(LiteraturRepository)

	// r.HandleFunc("/literatur", middleware.Auth(middleware.UploadFile(h.CreateLiteratur))).Methods("POST")

	// r.HandleFunc("/book", middleware.Auth(middleware.UploadCover(middleware.UploadPDF(h.AddBook)))).Methods("POST")

	r.HandleFunc("/literatur", middleware.Auth(middleware.UploadPDF(h.CreateLiteratur))).Methods("POST")

	r.HandleFunc("/literaturs", h.FindLiteraturs).Methods("GET")
	r.HandleFunc("/literatur/{id}", h.GetLiteratur).Methods("GET")

	
}
