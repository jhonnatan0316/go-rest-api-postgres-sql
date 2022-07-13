package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonnatan0316/go-rest-api-postgres-sql/db"
	"github.com/jhonnatan0316/go-rest-api-postgres-sql/db/models"
)

func GetGlobers(writer http.ResponseWriter, request *http.Request) {

	var globers []models.Globers
	db.DB.Find(&globers)
	json.NewEncoder(writer).Encode(&globers)
}

func GetGlober(writer http.ResponseWriter, request *http.Request) {
	var glober models.Globers
	params := mux.Vars(request)
	db.DB.First(&glober, params["id"])

	if glober.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Glober not found!"))
		return
	}

	db.DB.Model(&glober).Association("WorkingEcosystems").Find(&glober.WorkingEcosystems)
	json.NewEncoder(writer).Encode(&glober)
}

func CreateGlober(writer http.ResponseWriter, request *http.Request) {
	var newGlober models.Globers
	json.NewDecoder(request.Body).Decode(&newGlober)
	createdWe := db.DB.Create(&newGlober)

	error := createdWe.Error
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(error.Error()))
	}

	json.NewEncoder(writer).Encode(&newGlober)
}

func DeleteGlober(writer http.ResponseWriter, request *http.Request) {
	var glober models.Globers
	params := mux.Vars(request)
	db.DB.First(&glober, params["id"])

	if glober.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Glober not found!"))
		return
	}

	db.DB.Delete(&glober)
	writer.WriteHeader(http.StatusOK)
}
