package main

import (
	"github.com/NDLPhat/demo_golang/grpc-ytb/proto"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)
	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}
		
		req := &proto.Request{A: int64(a), B: int64(b)}
		// response, err := client.Add(ctx, req)
		response, err := client.Add(ctx, req)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.result)
			})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()
			})
		}

	})
	
	g.GET("/mult/:a/:b", func(ctx *gin.Context) {

	})

	if err := g.Run("8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
