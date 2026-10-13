package routers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

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

func UploadOneFile(ctx *gin.Context){
	file, err := ctx.FormFile("image-1")
	

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error while uploading file: %s", err.Error()),
		})
		return
	}

	log.Printf("Filename : %s", file.Filename) // need to sanitize the filename , KEEP IN MIND !

	destionatin := filepath.Join("images/", filepath.Base(file.Filename))
	ctx.SaveUploadedFile(file, destionatin)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("File named %s uploaded successfully !", file.Filename),
	})

}

func UploadManyFiles(ctx *gin.Context){

	// form parsing
	form , err := ctx.MultipartForm()
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}

	// Files processing
	files := form.File["files"]
	for _, file := range files{
		destination := filepath.Join("images/", filepath.Base(file.Filename))
		ctx.SaveUploadedFile(file, destination)
	}

	// Return part 
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Files uploaded successfully !",
	})
}