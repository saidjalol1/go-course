package main

import (
	"log"
	"routing_http_methods/config"
	"routing_http_methods/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	if config.GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/", routers.WelocmePage)
	router.GET("/products", routers.ProductListing)
	router.POST("/product/add", routers.ProductAdding)


	log.Printf("[INFO] Starting server on http://localhost:8080/")

	if err := router.Run(config.GIN_PORT); err != nil{
		log.Fatalf("[ERROR] Error occured while starting the server on https://localhost%s", config.GIN_PORT)
	} 

}