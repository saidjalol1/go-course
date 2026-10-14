package routers

import (
	"custom_validators/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Welcome(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Welcome to our API made in GIN !",
	})
}

func FillProfile(ctx *gin.Context){
	var RequestBody config.Profile

	if err := ctx.ShouldBindJSON(&RequestBody); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Profile data : Lastname: %s, Name: %s, Bio: %s", RequestBody.LastName, RequestBody.Name, RequestBody.Bio),
	})
}