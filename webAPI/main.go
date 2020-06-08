package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	for i := range company {
		index++
		if company[i].Id == id {
			break
		}
	}
	fmt.Println("Endpoint Hit: returnData")
	fmt.Fprintln(w, "Employee data:")
	json.NewEncoder(w).Encode(company[index])
}

//Function to fetch all employee data
func returnAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllData")
	json.NewEncoder(w).Encode(company)
}

//Function to add a new Employee
func addEmployee(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var emp employee
	json.Unmarshal(reqBody, &emp)
	// update our global Articles array to include
	// our new Article
	company = append(company, emp)
	fmt.Println("Endpoint Hit: addEmp")
	fmt.Fprintln(w, "Added a new Employee!")
	json.NewEncoder(w).Encode(emp)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/allEmpData", returnAllData).Methods("GET")
	myRouter.HandleFunc("/empData/{id}", returnData).Methods("GET")
	myRouter.HandleFunc("/addEmp", addEmployee).Methods("POST")
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
