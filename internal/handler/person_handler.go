package handler

import (
	"encoding/json"
	"junior/internal/model"
	"junior/internal/service"
	"junior/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var personService = service.NewPersonService()

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/api/people", createPerson).Methods("POST")
	r.HandleFunc("/api/people", getPeople).Methods("GET")
	r.HandleFunc("/api/people/{id}", getPersonByID).Methods("GET")
	r.HandleFunc("/api/people/{id}", updatePerson).Methods("PUT")
	r.HandleFunc("/api/people/{id}", deletePerson).Methods("DELETE")
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	logger.Log.WithFields(logger.PostGroup()).Info("Creating new person")

	var p model.Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		logger.Log.WithFields(logger.PostGroup()).Error("Failed to decode request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPerson, err := personService.CreatePerson(p)
	if err != nil {
		logger.Log.WithFields(logger.PostGroup()).Error("Failed to create person: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(logger.PostGroup()).Info("Person created successfully with ID: ", newPerson.ID)
	json.NewEncoder(w).Encode(newPerson)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	logger.Log.WithFields(logger.ReqGroup()).Info("Getting all people")

	people, err := personService.GetPeople()
	if err != nil {
		logger.Log.WithFields(logger.ReqGroup()).Error("Failed to get people: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(logger.ReqGroup()).Info("Successfully retrieved ", len(people), " people")
	json.NewEncoder(w).Encode(people)
}

func getPersonByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.WithFields(logger.ReqGroup()).Error("Invalid ID format: ", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	logger.Log.WithFields(logger.ReqGroup()).Info("Getting person with ID: ", id)

	person, err := personService.GetPersonByID(id)
	if err != nil {
		logger.Log.WithFields(logger.ReqGroup()).Error("Failed to get person: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(logger.ReqGroup()).Info("Successfully retrieved person with ID: ", id)
	json.NewEncoder(w).Encode(person)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.WithFields(logger.PutGroup()).Error("Invalid ID format: ", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	logger.Log.WithFields(logger.PutGroup()).Info("Updating person with ID: ", id)

	var p model.Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		logger.Log.WithFields(logger.PutGroup()).Error("Failed to decode request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedPerson, err := personService.UpdatePerson(id, p)
	if err != nil {
		logger.Log.WithFields(logger.PutGroup()).Error("Failed to update person: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(logger.PutGroup()).Info("Successfully updated person with ID: ", id)
	json.NewEncoder(w).Encode(updatedPerson)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.WithFields(logger.DeleteGroup()).Error("Invalid ID format: ", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	logger.Log.WithFields(logger.DeleteGroup()).Info("Deleting person with ID: ", id)

	err = personService.DeletePerson(id)
	if err != nil {
		logger.Log.WithFields(logger.DeleteGroup()).Error("Failed to delete person: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(logger.DeleteGroup()).Info("Successfully deleted person with ID: ", id)
	w.WriteHeader(http.StatusNoContent)
}
