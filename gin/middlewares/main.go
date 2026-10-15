package main


import (
	"log"
	"net/http"
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
	// server.Use(gin.Recovery())
	server.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
    	log.Printf("[PANIC RECOVERED]: %v", recovered)

    	c.JSON(http.StatusInternalServerError, gin.H{
        	"status":  500,
        	"message": "Internal Server Error occurred",
    	})
	}))

	server.Use(config.MyLoggerMiddleware())
	server.Use(config.CutomMiddleware1())
	server.Use(config.CutomMiddleware2())
	server.Use(config.CutomMiddleware3())
	server.GET("/", routers.Welcome)
	server.POST("/panic", routers.RouteWithPanic)


	log.Printf("[ INFO ] Server is being initialized at http://localhost%s", config.GIN_PORT)
	if err:= server.Run(config.GIN_PORT); err != nil {
		log.Fatalf("[ ERROR ] Error: %s", err)																																		
	}
}