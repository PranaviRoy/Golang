package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

//Book model creates the book relation in our db
type Book struct {
	gorm.Model
	Name            string `json:"Name"`
	Author          string `json:"Author"`
	PublicationYear string `json:"Year"`
	ISBN            string `json:"ISBN"`
}

// our initial migration function
func InitialMigration() {
	db, err := gorm.Open("mysql", "user:password@(localhost)/book?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the database!")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Book{})
}

func AllBooks(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w, "All Books Endpoint Hit")

	db, err := gorm.Open("mysql", "user:password@(localhost)/books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the database!")
	}
	defer db.Close()

	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func NewBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "New Book Endpoint Hit")

	db, err := gorm.Open("mysql", "user:password@(localhost)/books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the database!")
	}
	defer db.Close()

	vars := mux.Vars(r)
    name := vars["name"]
    email := vars["email"]

    db.Create(&User{Name: name, Email: email})
    fmt.Fprintf(w, "New User Successfully Created")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Delete Book Endpoint Hit")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Update Book Endpoint Hit")
}
