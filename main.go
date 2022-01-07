package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"net/http"
)

type Item struct {
	UID string `json: "UID"`
	Name string `json: "Name"`
	Desc string `json: "Desc"`
	Price float64 `json: "Price"`

}

// global variable
var Inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homepage()")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}
