package routers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Gin support two types of params ,
// First -  /user/:name  matches to single path segment
// Second - /user/*actions  matches to everything after prefix


func Welcome(ctx *gin.Context){
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello %s", name),
	})
}

func ReturnRequestsAddress(ctx *gin.Context){
	action := ctx.Param("action")

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("This is what the action parameter holds in that address: %s", action),
	})
}

func ListObjects(ctx *gin.Context) {
	limitParam := ctx.DefaultQuery("limit", "10")
	offsetParam := ctx.DefaultQuery("offset", "0")

	limit,err := strconv.Atoi(limitParam)
	if err != nil{
		limit = 10
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil{
		offset = 0
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User requested to pages %d / %d ", offset, limit),
	})
}