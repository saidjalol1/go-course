package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)


func SendRequest(client *http.Client, cfg *Config) Result {
	fmt.Println("Making request ......")
	var request *http.Request
	var err error


	var bodyReader io.Reader = nil
	if cfg.Method == "POST" || cfg.Method == "PUT" || cfg.Method == "PATCH" {
	    if cfg.Body != "" {
        	bodyReader = bytes.NewReader([]byte(cfg.Body))
    	}
	}

	request, err = http.NewRequest(cfg.Method, cfg.URL, bodyReader)
	if err != nil{
		return  Result{Error: err}
	}

	start := time.Now()

	responce, err := client.Do(request)

	duration := time.Since(start)

	if err != nil{
		return  Result{Duration: duration, Error: err}
	}

	defer responce.Body.Close()

	return  Result{
		StatusCode: responce.StatusCode,
		Duration: duration,
		Error: nil,
	}
}
