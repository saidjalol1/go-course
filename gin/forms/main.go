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
	server.MaxMultipartMemory = 2 << 20 // max 2mb of upload size


	server.GET("/", routers.Welcome)
	server.POST("/post-form", routers.SubmitFormRequest)
	server.POST("/upload-single-file", routers.UploadOneFile)
	server.POST("/upload-many-files", routers.UploadManyFiles)
	
	
	log.Printf("[INFO]: Starting server at http://localhost%s", config.GIN_PORT)
	if err:= server.Run(config.GIN_PORT);err != nil{
		log.Fatalf("[ERROR]: %s", err)
	}
}