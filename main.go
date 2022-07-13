package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonnatan0316/go-rest-api-postgres-sql/db"
	"github.com/jhonnatan0316/go-rest-api-postgres-sql/db/models"
	"github.com/jhonnatan0316/go-rest-api-postgres-sql/routes"
)

func main() {
	db.Conection()

	db.DB.AutoMigrate(models.WorkingEcosystem{})
	db.DB.AutoMigrate(models.Globers{})

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeIndex)
	router.HandleFunc("/globers", routes.CreateGlober).Methods("POST")
	router.HandleFunc("/globers", routes.GetGlobers).Methods("GET")
	router.HandleFunc("/globers/{id}", routes.GetGlober).Methods("GET")
	router.HandleFunc("/globers/{id}", routes.DeleteGlober).Methods("DELETE")

	router.HandleFunc("/working-ecosystems", routes.CreateWe).Methods("POST")
	router.HandleFunc("/working-ecosystems", routes.GetworkingEcosystems).Methods("GET")
	router.HandleFunc("/working-ecosystems/{id}", routes.GetWe).Methods("GET")
	router.HandleFunc("/working-ecosystems/{id}", routes.DeleteTWe).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
