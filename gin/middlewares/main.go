package main


import (
	"log"
	"middlewares/config"
	"middlewares/routers"
	"github.com/gin-gonic/gin"
)


func main(){

	if config.GIN_MODE == "release"{
		gin.SetMode(gin.ReleaseMode)
	}
	
	// KEEP IN MIND : The middlewares in Gin are executed in LIFO pattern 
	// which means the first registered middleware is the last to be completed 
	// you can see how to write middleware funcs in config/middlewares.go and at gin's official documentation

	server := gin.New()
	server.Use(config.MyLoggerMiddleware())
	server.GET("/", routers.Welcome)


	log.Printf("[ INFO ] Server is being initialized at http://localhost%s", config.GIN_PORT)
	if err:= server.Run(config.GIN_PORT); err != nil {
		log.Fatalf("[ ERROR ] Error: %s", err)																																		
	}
}