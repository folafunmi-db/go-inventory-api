package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Item struct {
	UID   string  `json: "UID"`
	Name  string  `json: "Name"`
	Desc  string  `json: "Desc"`
	Price float64 `json: "Price"`
}

// global variable
var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homepage()")
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)

	// delete the item
	_deleteItemAtUID(params["uid"])

	// create it with new data
	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(item)

}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtUID(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemAtUID(uid string) {
	for index, item := range inventory {
		if item.UID == uid {
			inventory = append(inventory[:index], inventory[index+1:]...)
		}
	}
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function called: inventory()")

	json.NewEncoder(w).Encode(inventory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println("Function called: createItem()")

	var item Item

	_ = json.NewDecoder(r.Body).Decode(&item)

	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(item)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	inventory = append(inventory, Item{
		UID:   "0",
		Name:  "Cheese",
		Desc:  "A fine block of cheese",
		Price: 5,
	})
	inventory = append(inventory, Item{
		UID:   "1",
		Name:  "Milk",
		Desc:  "A jug of milk",
		Price: 15,
	})
	handleRequests()
}
