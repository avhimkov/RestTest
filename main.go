package main

import (
	//"fmt"
	//"encoding/json"
	"log"
	"github.com/gorilla/mux"

	"net/http"
)

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request)  {
	
}


func GetPerson(w http.ResponseWriter, r *http.Request)  {

}

func CreatePerson(w http.ResponseWriter, r *http.Request)  {

}

func DeletePerson(w http.ResponseWriter, r *http.Request)  {

}
