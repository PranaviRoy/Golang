package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/books_grpc_db/connector"
	proto "github.com/golang/books_grpc_db/services"
	"google.golang.org/protobuf/types/known/emptypb"
)

//AddBookHandler creates a new record in the database
func AddBookHandler(ctx *gin.Context) {
	client := connector.Connection()

	newBook := proto.Book{}
	reqBody, err := ioutil.ReadAll(ctx.Request.Body) //request body is defined to read the .json file body
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Println("Kindly check the book structure")
	}

	res := json.Unmarshal(reqBody, &newBook) //pass the json body in proto format to newCluster
	fmt.Println(res)
	//req := &proto.RequestCreateCluster{OrgId: 2, UsrId: 2, PolicyId: 123, ClusterName: "Cl-1", Status: "Active", Location: "India"}
	if response, err := client.AddBook(ctx, &newBook); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

//FetchBookHandler handles the get request to fetch book data
func FetchBookHandler(ctx *gin.Context) {
	client := connector.Connection()

	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("Error in parsing cluster id param")
	}
	req := &proto.Book{Id: int64(bookID)}
	if response, err := client.FetchBook(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": response,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

//FetchBooksHandler handles the request to return the details of all the books in the table
func FetchBooksHandler(ctx *gin.Context) {
	client := connector.Connection()

	req := &emptypb.Empty{}
	if response, err := client.FetchBooks(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": response,
		})
		fmt.Println("Succesfully returning book details!")
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Some error occured while fetching data!")
	}
}

//UpdateBookHandler handles the request to modify the details of a particular book
func UpdateBookHandler(ctx *gin.Context) {
	client := connector.Connection()

	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("Error in parsing book id param")
	}
	updatedBook := proto.Book{Id: int64(bookID)}
	reqBody, err := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Println("Kindly check the book structure")
	}

	if res := json.Unmarshal(reqBody, &updatedBook); res != nil {
		fmt.Println(res)
	}

	if response, err := client.UpdateBook(ctx, &updatedBook); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

//DeleteBookHandler deletes a particular tuple
func DeleteBookHandler(ctx *gin.Context) {
	client := connector.Connection()

	clusterID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("Error in parsing cluster id param")
	}
	req := &proto.Book{Id: int64(clusterID)}
	if response, err := client.DeleteBook(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Result": fmt.Sprint(response.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
