package main

import (
	"log"
	"parameters/config"
	"parameters/routers"

	"github.com/gin-gonic/gin"
)


func main(){
	if config.GIN_MODE == "release"{
		gin.SetMode(gin.ReleaseMode)
	}

	route := gin.Default()

	// Routes with two types of params, you can see how to handle them inside routers.go
	route.GET("/:name", routers.Welcome)
	route.GET("/actions/*action", routers.ReturnRequestsAddress)
	route.GET("/products", routers.ListObjects)

	log.Printf("[INFO] Server is being initialized on http://localhost%s", config.GIN_PORT)
	if err := route.Run(config.GIN_PORT); err != nil{
		log.Fatalf("[ERROR] %v", err)
	}
}