package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func List(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Create(&personality)
	json.NewEncoder(w).Encode(personality)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.Delete(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(personality)
}
