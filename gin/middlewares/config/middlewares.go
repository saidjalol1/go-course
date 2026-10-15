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

		ctx.Next()


		// after request finished
		latency := time.Since(t)
		log.Printf("[ LATENCY ] This is the time the request took to complete %s", latency)

		status := ctx.Writer.Status()
		log.Printf("[ STATUS CODE ] THIS IS FROM LOGGER : %d", status)
	}
}