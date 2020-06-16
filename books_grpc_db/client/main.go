package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang/books_grpc_db/handlers"
)

func main() {
	web := gin.Default()
	web.POST("/addBook", handlers.AddBookHandler)
	web.GET("/fetchBook/:id", handlers.FetchBookHandler)
	web.GET("/fetchBooks", handlers.FetchBooksHandler)
	web.PUT("/updateBook/:id", handlers.UpdateBookHandler)
	web.DELETE("/deleteBook/:id", handlers.DeleteBookHandler)

	if err := web.Run(":4040"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
