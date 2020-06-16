package handlers

import (
	"fmt"
	"net/http"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/golang/books_grpc_db/connector"
	proto "github.com/golang/books_grpc_db/services"
)

//AddBookHandler creates a new record in the database 
func AddBookHandler(ctx *gin.Context) {
	client := connector.Connection()

	newBook := proto.Request{}
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
