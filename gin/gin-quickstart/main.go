package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gin-quickstart/config"
	"net/http"
)

func main() {
	
	if config.GIN_MODE == "release"{
		gin.SetMode(gin.ReleaseMode)
	}
	
	router := gin.Default()


	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})
	})

	
	log.Printf("[INFO] Starting server on http://localhost%s", config.PORT)
	err := router.Run(config.PORT)
	if err != nil {
		log.Fatalf("[ERROR] Server failed to start: %v", err)
	}

}