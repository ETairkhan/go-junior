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

type Handler struct {
	service *service.PersonService
}

func NewHandler(service *service.PersonService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes(r *mux.Router) {
	r.HandleFunc("/api/people", h.createPerson).Methods("POST")
	r.HandleFunc("/api/people", h.getPeople).Methods("GET")
	r.HandleFunc("/api/people/{id}", h.getPersonByID).Methods("GET")
	r.HandleFunc("/api/people/{id}", h.updatePerson).Methods("PUT")
	r.HandleFunc("/api/people/{id}", h.deletePerson).Methods("DELETE")
}

func (h *Handler) createPerson(w http.ResponseWriter, r *http.Request) {
	var p model.Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newPerson, err := h.service.CreatePerson(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Log.Info("Successfully created person with ID: ", newPerson.ID)
	json.NewEncoder(w).Encode(newPerson)

}

func (h *Handler) getPeople(w http.ResponseWriter, r *http.Request) {
	gender := r.URL.Query().Get("gender")
	nationality := r.URL.Query().Get("nationality")

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	people, err := h.service.GetFilteredPeople(gender, nationality, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.Infof("Successfully retrieved %d people (filters: gender=%s, nationality=%s, page=%d, limit=%d)",
	len(people), gender, nationality, page, limit)
json.NewEncoder(w).Encode(people)

}

func (h *Handler) getPersonByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	person, err := h.service.GetPersonByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Log.Infof("Successfully retrieved person with ID: %d", id)
json.NewEncoder(w).Encode(person)

}

func (h *Handler) updatePerson(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	var p model.Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedPerson, err := h.service.UpdatePerson(id, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Log.Infof("Successfully updated person with ID: %d", id)
json.NewEncoder(w).Encode(updatedPerson)

}

func (h *Handler) deletePerson(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	err = h.service.DeletePerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Log.Infof("Successfully deleted person with ID: %d", id)
w.WriteHeader(http.StatusNoContent)

}
