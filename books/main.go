package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Books corner!")
	fmt.Println("Endpoint hit: HomePage!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloBooks)
	myRouter.HandleFunc("/books", AllBooks).Methods("GET")
	myRouter.HandleFunc("/deleteBook/{isbn}", DeleteBook).Methods("DELETE")
	myRouter.HandleFunc("/updateBook/{isbn}", UpdateBook).Methods("PUT")
	myRouter.HandleFunc("/book", NewBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	InitialMigration()
	// Handle Subsequent requests
	handleRequests()
}
