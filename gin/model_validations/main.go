package main

import (
	"log"
	"model_validation/config"
	"model_validation/routers"

	"github.com/gin-gonic/gin"
)

func main(){

	if config.GIN_MODE == "release"{
		gin.SetMode(gin.ReleaseMode)
	}


	server := gin.Default()
	server.GET("/", routers.Welcome)
	server.POST("/login", routers.LoginEndpoint)
	server.POST("/upload-profile-image", routers.UploadProfilePhoto)
	server.POST("/create-products", routers.CreateProduct)
	server.GET("/filter-products", routers.FilterProducts)



	log.Printf("[ INFO ] Starting the server at http://localhost%s", config.GIN_PORT)
	if err := server.Run(config.GIN_PORT); err != nil{
		log.Fatalf("[ ERROR ] Error occured while starting the server: %s", err)
	}
}