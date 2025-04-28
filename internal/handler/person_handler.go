package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "junior/internal/model"
    "junior/internal/service"

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
    var p model.Person
    json.NewDecoder(r.Body).Decode(&p)

    newPerson, err := personService.CreatePerson(p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(newPerson)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
    people, err := personService.GetPeople()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(people)
}

func getPersonByID(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, _ := strconv.Atoi(idStr)

    person, err := personService.GetPersonByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(person)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, _ := strconv.Atoi(idStr)

    var p model.Person
    json.NewDecoder(r.Body).Decode(&p)

    updatedPerson, err := personService.UpdatePerson(id, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(updatedPerson)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, _ := strconv.Atoi(idStr)

    err := personService.DeletePerson(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
