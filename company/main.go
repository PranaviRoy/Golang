package main

import(
	"fmt"
)

//Creating a slice of employee objects
var company []employee

//Creating an employee struct to hold employee data
type employee struct{
	id int;
	name string;
	role string;
	address string;
}

//Function to fetch employee data, given emp id
func getEmployeeDetails(id int) employee{
	var index int = -1
	for i := range company{
		index ++
		if company[i].id == id{
			break
		}
	}
	fmt.Println("Employee data:")
	return company[index]
}

//Function to add a new Employee, takes an employee type argument
func addEmployee(emp employee) []employee{
	company = append(company, emp)
	fmt.Println("Employee added!")
	return company
}

//Function to remove employee data, given emp id
func removeEmployee(id int) []employee{
	var index int = -1
	for i := range company{
		index ++
		if company[i].id == id{
			break
		}
	}
	company = append(company[:index], company[index+1:]...)
	fmt.Println("Employee data removed!")
	return company
}

//Function to update an employee's details, takes employee type argument
func updateEmployee(emp employee) employee{
	var index int = -1
	for i := range company{
		index ++
		if company[i].id == emp.id{
			break
		}
	}
	company[index] = emp
	fmt.Println("Employee data updated!" )
	return company[index]
}

func main(){
	//Employee objects (dummy data for testing)
	employee1 := employee{
		id: 1, 
		name: "Sherlock Holmes", 
		role: "UI Designer", 
		address: "221B Baker Street"}
	
	employee2 := employee{
		id: 2,
		name: "Sirius Black",
		role: "Project Manager",
		address: "12 Grimmauld Place"}	
	
	employee3 := employee{
		id: 3,
		name: "Buffy Summers",
		role: "Python Developer",
		address: "1630 Revello Drive"}

	employee4 := employee{
		id: 4,
		name: "Jon Arbuckle",
		role: "UX Designer",
		address: "711 Maple Street"}

	employee5 := employee{
		id: 5,
		name: "Spongebob Squarepants",
		role: "Concept Designer",
		address: "124 Conch Street"}

	employee6 := employee{
		id: 6,
		name: "Clark Klent",
		role: "Software Developer",
		address: "344 Clinton Street"}


	//Adding Employees
	fmt.Println(addEmployee(employee1))
	fmt.Println("")
	fmt.Println(addEmployee(employee2))
	fmt.Println("")
	fmt.Println(addEmployee(employee3))
	fmt.Println("")
	fmt.Println(addEmployee(employee4))
	fmt.Println("")
	fmt.Println(addEmployee(employee5))
	fmt.Println("")
	fmt.Println(addEmployee(employee6))
	fmt.Println("")

	//Removing employee with employee id 5
	fmt.Println(removeEmployee(5))
	fmt.Println("")
	
	//Fetching the data of employee with id 4
	fmt.Println(getEmployeeDetails(4))
	fmt.Println("")

	//creating another employee object to test update function
	employee6Updated := employee{
		id: 6,
		name: "Clark Klent",
		role: "Software Developer",
		address: "344 Clinton St., Apt. 3B, Metropolis, USA"}

	//Updating data of employee with id 6
	fmt.Println(updateEmployee(employee6Updated))
}