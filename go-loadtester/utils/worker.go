package utils

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)


func RunLoadTest(cfg *Config){
	startTime := time.Now()

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns: cfg.Concurrency,
			MaxIdleConnsPerHost: cfg.Concurrency,
		},
	}

	jobs := make(chan int, cfg.Concurrency)
	results := make(chan Result, cfg.Requests)

	var wg sync.WaitGroup

	for i := 1; i <= cfg.Concurrency; i++{
		wg.Add(1)
		go func ()  {
			defer wg.Done()

			for range jobs{
				res := SendRequest(client, cfg)
				results <- res
			}
		}()
	}

	for i := 1; i <= cfg.Requests; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var totalRequests int
	var successCount int
	var errorCount int
	var totalDuration time.Duration

	for res := range results {
		totalRequests++
		if res.Error != nil {
			errorCount++
		} else {
			successCount++
			totalDuration += res.Duration
		}
	}

	testDuration := time.Since(startTime)


	fmt.Println("\n--- LOAD TEST RESULTS ---")
	fmt.Printf("Total Requests Sent: %d\n", totalRequests)
	fmt.Printf("Successful Requests: %d\n", successCount)
	fmt.Printf("Failed Requests:     %d\n", errorCount)
	fmt.Printf("Total Time Elapsed:  %v\n", testDuration)
	
	if successCount > 0 {
		avgLatency := totalDuration / time.Duration(successCount)
		fmt.Printf("Average Latency:     %v\n", avgLatency)
	}
}