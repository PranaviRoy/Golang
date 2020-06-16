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

	if err := web.Run(":4040"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
