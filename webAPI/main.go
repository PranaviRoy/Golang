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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintln(w, "check the structure!")
	}
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
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "check the structure!")
	}
	var emp employee
	json.Unmarshal(reqBody, &emp)
	// update our global Articles array to include
	// our new Article
	company = append(company, emp)
	fmt.Println("Endpoint Hit: addEmp")
	fmt.Fprintln(w, "Added a new Employee!")
	json.NewEncoder(w).Encode(emp)
}

//Function to update an already existing data
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintln(w, "check the parameter passed!")
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "check the structure!")
	}
	var emp employee
	json.Unmarshal(reqBody, &emp)
	var index int = -1
	for i := range company {
		index++
		if company[i].Id == id {
			// company[i].Name = emp.Name
			// company[i].Role = emp.Role
			// company[i].Address = emp.Address
			company[i] = emp
		}
	}
	fmt.Println("Endpoint Hit: updateEmp")
	fmt.Fprintln(w, "Updated employee's data!")
	json.NewEncoder(w).Encode(company)
}

//Function to remove employee data, given emp id
func removeEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintln(w, "check the parameter passed!")
	}
	var index int = -1
	for i := range company {
		index++
		if company[i].Id == id {
			company = append(company[:index], company[index+1:]...)
		}
	}

	fmt.Println("Emdpoint Hit: deleteData")
	fmt.Fprintln(w, "Employee data removed!")
	json.NewEncoder(w).Encode(company)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/allEmpData", returnAllData).Methods("GET")
	myRouter.HandleFunc("/empData/{id}", returnData).Methods("GET")
	myRouter.HandleFunc("/addEmp", addEmployee).Methods("POST")
	myRouter.HandleFunc("/updateEmp/{id}", updateEmployee).Methods("PUT")
	myRouter.HandleFunc("/deleteData/{id}", removeEmployee).Methods("DELETE")
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
