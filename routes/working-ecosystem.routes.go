package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonnatan0316/go-rest-api-postgres-sql/db"
	"github.com/jhonnatan0316/go-rest-api-postgres-sql/db/models"
)

func GetWe(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var we models.WorkingEcosystem
	db.DB.First(&we, params["id"])

	if we.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Working Ecosystem not found!"))
		return
	}

	json.NewEncoder(writer).Encode(&we)
}

func GetworkingEcosystems(writer http.ResponseWriter, request *http.Request) {
	var workingEcosystems []models.WorkingEcosystem
	db.DB.Find(&workingEcosystems)
	json.NewEncoder(writer).Encode(&workingEcosystems)
}

func CreateWe(writer http.ResponseWriter, request *http.Request) {
	var we models.WorkingEcosystem
	json.NewDecoder(request.Body).Decode(&we)
	createdTask := db.DB.Create(&we)
	err := createdTask.Error

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
	}

	json.NewEncoder(writer).Encode(&we)
}

func DeleteTWe(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var we models.WorkingEcosystem
	db.DB.First(&we, params["id"])

	if we.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Working Ecosystem not found"))
		return
	}

	db.DB.Unscoped().Delete(&we)
	writer.WriteHeader(http.StatusOK)
}
