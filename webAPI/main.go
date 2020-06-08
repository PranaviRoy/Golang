package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

//Creating a slice of employee objects
var company []employee

//Creating an employee struct to hold employee data
type employee struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Role    string `json:"Role"`
	Address string `json:"Address"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//Function to fetch employee data, given emp id
func returnData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["Id"])
	var index int = -1
	for i := range company{
		index++
		if company[i].Id == id{
			break
		}
	}
	fmt.Fprintln(w, "Employee data:")
	json.NewEncoder(w).Encode(company[index])
}

//Function to fetch all employee data
func returnAllData(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnData")
	json.NewEncoder(w).Encode(company)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/allEmpData", returnAllData).Methods("GET")
	myRouter.HandleFunc("/empData/{id}", returnData).Methods("GET")
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	//Employee objects (dummy data for testing)

	company = []employee{
		employee{
			Id:      1,
			Name:    "Sherlock Holmes",
			Role:    "UI Designer",
			Address: "221B Baker Street"},
		employee{
			Id:      2,
			Name:    "Sirius Black",
			Role:    "Project Manager",
			Address: "12 Grimmauld Place"}}

	handleRequests()
}
