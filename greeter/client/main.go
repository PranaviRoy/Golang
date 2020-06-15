package main

import(
	"log"
	"fmt"
	"net/http"
	"strconv"
	"greeter/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

//GreeterHandler handles the greeting function
func GreeterHandler(ctx *gin.Context) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewGreetingServiceClient(conn)

	times, err := strconv.ParseInt(ctx.Param("times"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter Name"})
		return
	}
	req := &proto.Request{Name: ctx.Param("name"), Times: times}
	if response, err := client.Reply(ctx, req); err == nil{
		ctx.JSON(http.StatusOK, gin.H{
			"reply": fmt.Sprint(response.Reply),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	}
}

func main() {
	g := gin.Default()
	g.GET("/greet/:name/:times", GreeterHandler)

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}