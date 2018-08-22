package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var people []Person

func main() {

	people = append(people, Person{Id: "1", FirstName: "Leonardo", LastName: "Marquezani", Address: &Address{City: "Sao Paulo", State: "SP"}})
	people = append(people, Person{Id: "2", FirstName: "Cintia", LastName: "Brum", Address: &Address{City: "Pres. prudente", State: "SP"}})
	people = append(people, Person{Id: "3", FirstName: "Rafael", LastName: "Silva", Address: &Address{City: "Parana", State: "PR"}})

	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {}

func CreatePerson(w http.ResponseWriter, r *http.Request) {}

func DeletePerson(w http.ResponseWriter, r *http.Request) {}

type Person struct {
	Id        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
