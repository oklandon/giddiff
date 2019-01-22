package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Version struct {
	ID      string            `json:"id,omitempty"`
	Title   string            `json:"title,title"`
	Content map[string]string `json:"content,content"`
}

func BuildComparisonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var version Version
	_ = json.NewDecoder(req.Body).Decode(&version)
	version.ID = params["id"]
	/*
		fetch prev from db
		spin up server
		git init
		save prev
		save post payload
		save git diff as a string
		parse the string into something useful
		insert new version into db
	*/
	json.NewEncoder(w).Encode("everthing is ok")
}

func main() {
	router := mux.NewRouter()
	// people = append(people, Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
	// people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})
	router.HandleFunc("/version", BuildComparisonEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}
