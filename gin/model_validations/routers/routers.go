package routers

import (
	"fmt"
	"log"
	"model_validation/config"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


func Welcome(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Welcome to our API made in GIN !",
	})
}

func LoginEndpoint(ctx *gin.Context){
	var requestBody config.LoginRequestBody
	if err:= ctx.BindJSON(&requestBody);err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("[ LOGIN ATTEMPT ] Username: %s, Password: %s", requestBody.Username, requestBody.Password)
	ctx.JSON(http.StatusOK, gin.H{
		"message":"This is a test endpoint to simulate login with model validation !",
	})
}

func UploadProfilePhoto(ctx *gin.Context){
	file, err := ctx.FormFile("profile_image")
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	destionation := filepath.Join(fmt.Sprintf("%s/profile_images/", config.MEDIA_ROOT), filepath.Base(file.Filename))
	ctx.SaveUploadedFile(file, destionation)
	ctx.JSON(http.StatusOK, gin.H{
		"message":"File uploaded successfully !",
	})
}


func CreateProduct(ctx *gin.Context){
	var requestBody config.ProductCreate
	if err:= ctx.BindJSON(&requestBody); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("[ PRODUCT CREATION ATTEMPT ] Product name: %s, Product price: %s", requestBody.Name, requestBody.Price)
	ctx.JSON(http.StatusOK, gin.H{
		"message":"This was the test endpoint to simulate the product creation logic",
	})
}


func FilterProducts(ctx *gin.Context){
	params := ctx.DefaultQuery("category", "clothes")


	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("This endpoint should return the product with the category of %s", params),
	})
}