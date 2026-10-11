package main


import (
	"github.com/saidjalol1/go-loadtester/utils"
)

func main(){

	cfg := &utils.Config{
		URL:         "https://httpbin.org/get",
		Requests:    1000,                       
		Concurrency: 5,                       
		Method:      "GET",
	}

	// Execute the load test engine
	utils.RunLoadTest(cfg)
}