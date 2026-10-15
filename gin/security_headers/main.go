package main

import (
	"log"
	"net/http"
	"security_headers/config"
	"security_headers/routers"

	"github.com/gin-gonic/gin"
)

// KEEP IN MIND: You need to set securty headers to protect
// your application from common securty vulnerabilities

// in this file you can see how to set them and
// what what each secirty headers protects your application from

func main(){

	if config.GIN_MODE == "release"{
		gin.SetMode(gin.ReleaseMode)
	}
	
	expectedHost := "localhost:8000"
	server := gin.Default()
	
	// Setting up securuty headers
	server.Use(func(ctx *gin.Context) {
		if ctx.Request.Host != expectedHost {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":"Invalid Host header"})
			return
		}

		// Prevents clickjacking , diables to render the page in iframe
		ctx.Header("X-Frame-Options", "DENY") 

		// this controls which resources can be loaded by the browser
		// This is the most effective defences against cross site scripting
		ctx.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")

		// Forces the browser to use HTTPS for all future requests to the domain for the specified max-age duration.
		ctx.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		// Controls how much referrer information is sent with outgoing requests. 
		// Without this header, the full URL (including query parameters 
		// that may contain tokens or sensitive data) can leak to third-party sites. 
		// strict-origin sends only the origin (domain) and only over HTTPS.
		ctx.Header("Referrer-Policy", "strict-origin")

		// Prevents MIME-type sniffing attacks. Without this header, 
		// browsers may interpret files as a different content type than declared, 
		// allowing attackers to execute malicious scripts disguised 
		// as harmless file types (e.g., uploading a .jpg that is actually JavaScript).
		ctx.Header("X-Content-Type-Options", "nosniff")

		// Restricts which browser features (geolocation, camera, microphone, etc.) can be used by the page. 
		// This limits the damage if an attacker manages to inject scripts, 
		// since those scripts cannot access sensitive device APIs.
		ctx.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
	})


	server.GET("/", routers.Welcome)

	log.Printf("[ INFO ] Starting the server at http://localhost%s", config.GIN_PORT)
	if err := server.Run(config.GIN_PORT);err != nil{
		log.Fatalf("[ ERROR ] Error occured while starting the server %s", err.Error())
	}
}