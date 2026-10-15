package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func Welcome(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Welcome !",
	})
}