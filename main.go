package main

import (
	"fmt"
	"literature/database"
	"literature/pkg/mysql"
	"literature/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// fmt.Println("Hello Ji")

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println(errEnv)
		panic("failed to load env file")
	}

	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running localhost:5000")

	http.ListenAndServe("localhost:5000", r)

}
