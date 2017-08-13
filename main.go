package main

import (
	//"fmt"
	//"encoding/json"
	"log"
	"github.com/gorilla/mux"

	"net/http"
	"encoding/json"
)
func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

	people = append(people, Person{ID:"1", FirstName:"Jon", LastName:"Doe", Address:&Address{City:"City X", State:"State X"}})
	people = append(people, Person{ID:"2", FirstName:"Koko", LastName:"Doe", Address:&Address{City:"City Z", State:"State Y"}})
	people = append(people, Person{ID:"3", FirstName:"Francis", LastName:"Sunday"})
}

func GetPeople(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	for _ , item := range people {
		if item.ID == params["id"]
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	for index, item := range people{
		if item.ID == params["id"] {
			people = append(people[:index], people[index + 1]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

type Person struct {
	ID           string `json:"id,omitempty"`
	FirstName    string `json:"firstname,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Address      *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

