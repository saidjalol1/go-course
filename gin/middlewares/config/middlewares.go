package config


import (
  "log"
  "time"

  "github.com/gin-gonic/gin"
)

func MyLoggerMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		
		t := time.Now()
		ctx.Set("example", "123456")

		// Before request
		log.Println("[ LOG Middleware ] Logging started .... ")
		
		ctx.Next()


		// after request finished
		latency := time.Since(t)
		log.Printf("[ LOG Middleware ] This is the time the request took to complete %s", latency)

		status := ctx.Writer.Status()
		log.Printf("[ STATUS CODE ] THIS IS FROM LOGGER : %d", status)
		log.Println("[ LOG Middleware ] Logging end ... ")
	}
}

func CutomMiddleware1() gin.HandlerFunc{
	return func(ctx *gin.Context){

		log.Printf("[ M1 ] This is the start of M1 ")

		ctx.Next()

		log.Printf("[ M1 ] This is the end of the M1 ")

	}
}

func CutomMiddleware2() gin.HandlerFunc{
	return func(ctx *gin.Context){

		log.Printf("[ M2 ] This is the start of M2 ")

		ctx.Next()

		log.Printf("[ M2 ] This is the end of the M2 ")

	}
}


func CutomMiddleware3() gin.HandlerFunc{
	return func(ctx *gin.Context){

		log.Printf("[ M3 ] This is the start of M3 ")

		ctx.Next()

		log.Printf("[ M3 ] This is the end of the M3 ")

	}
}