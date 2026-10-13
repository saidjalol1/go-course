package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Welcome(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Welcome to our API !",
	})
}


// You can use PostForm("key") and DefaultPostForm("key", "fallback")
// to handle both type of submission (application/x-www-form-urlencoded, and multipart/form-data)
// 
func SubmitFormRequest(ctx *gin.Context){
	message := ctx.PostForm("message")
	postData := ctx.DefaultPostForm("title", "This is the default title")

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("This is the message which is extracted from form submission via PostForm(): %s   ;", message),
		"message_2": fmt.Sprintf("This is the message which is extracted via DefaultPostForm() %s  ;", postData),
	})
}