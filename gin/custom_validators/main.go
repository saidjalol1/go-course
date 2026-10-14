package main 

import (
	"log"
	"custom_validators/config"
	"custom_validators/routers"
	"github.com/gin-gonic/gin"
)


func main(){

	server := gin.Default()
	server.GET("/", routers.Welcome)
	server.POST("/fill-profile", routers.FillProfile)

	
	log.Printf("[ INFO ] Server starting at http://localhost%s", config.GIN_PORT)
	if err := server.Run(config.GIN_PORT); err != nil {
		log.Fatalf("[ ERROR ] Error occured while starting the server: %s", err.Error())
	}
}