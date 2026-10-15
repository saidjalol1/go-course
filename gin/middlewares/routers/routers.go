package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Welcome(ctx *gin.Context){
	getVarFromMiddleware := ctx.MustGet("example")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("This is the value passed from middleware %s", getVarFromMiddleware),
	})
}
