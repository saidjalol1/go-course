package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)



func WelocmePage(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Request made with GET method, Welcome to our page",
	})
}


func ProductListing(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Request made with GET method , to get list of products",
	})
}

func ProductAdding(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Request made with POST method, to add product to store",
	})
}

