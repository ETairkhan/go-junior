package handler

import (
	"encoding/json"
	"junior/internal/model"
	"junior/internal/service"
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
	json.NewEncoder(w).Encode(newPerson)
}

func (h *Handler) getPeople(w http.ResponseWriter, r *http.Request) {
	people, err := h.service.GetPeople()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	w.WriteHeader(http.StatusNoContent)
}
