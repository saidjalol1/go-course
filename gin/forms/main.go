package main

import (
	"log"
	"forms/config"
	"forms/routers"
	"github.com/gin-gonic/gin"
)


func main(){
	if config.GIN_MODE == "release"{
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.Default()

	server.GET("/", routers.Welcome)
	server.POST("/post-form", routers.SubmitFormRequest)

	log.Printf("[INFO]: Starting server at http://localhost%s", config.GIN_PORT)
	if err:= server.Run(config.GIN_PORT);err != nil{
		log.Fatalf("[ERROR]: %s", err)
	}
}