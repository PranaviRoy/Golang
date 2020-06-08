package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
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
	db, err := gorm.Open("mysql", "user:password@(localhost)/books?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the database!")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Book{})
}

func AllBooks(w http.ResponseWriter, r *http.Request) {

	fmt.Println("All Books Endpoint Hit")

	db, err := gorm.Open("mysql", "user:password@(localhost)/books?parseTime=true")
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
	fmt.Println("New Book Endpoint Hit")

	db, err := gorm.Open("mysql", "user:password@(localhost)/books?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the database!")
	}
	defer db.Close()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "check the structure!")
	}
	var book Book
	json.Unmarshal(reqBody, &book)

	db.Create(&book)
	fmt.Fprintf(w, "Book info added successfully!")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Book Endpoint Hit")

	db, err := gorm.Open("mysql", "user:password@(localhost)/books?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the database!")
	}
	defer db.Close()

	vars := mux.Vars(r)
	var book Book
	db.Where("isbn = ?", vars["isbn"]).First(&book)
	db.Delete(&book)
	fmt.Fprintf(w, "Book info deleted successfully!")
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Book Endpoint Hit")

	db, err := gorm.Open("mysql", "user:password@(localhost)/books?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the database!")
	}
	defer db.Close()
	
	vars := mux.Vars(r)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "check the structure!")
	}
	var book, book1 Book
	json.Unmarshal(reqBody, &book)

	db.Where("isbn = ?", vars["isbn"]).Find(&book1)
	book1.Author = book.Author
	book1.Name = book.Name
	book1.PublicationYear = book.PublicationYear
	//book1 = book
	db.Save(book1)
	fmt.Fprintf(w, "Book info updated successfully!")	
}
