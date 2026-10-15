package routers

import (
	"fmt"
	"middlewares/config"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Welcome(ctx *gin.Context){
	getVarFromMiddleware := ctx.MustGet("example")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("This is the value passed from middleware %s", getVarFromMiddleware),
	})
}


func RouteWithPanic(ctx *gin.Context){
	var reqBody config.RequestBody

	if err := ctx.BindJSON(&reqBody); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	zero := 0
	result := reqBody.Integer / int64(zero)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("[ Result ] %d", result),
	})

}