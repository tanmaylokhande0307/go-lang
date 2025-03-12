package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// VideoWorker processes messages
func VideoWorker(workerID int, requests <-chan VideoRequest, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range requests {
		fmt.Printf("[Worker %d] Processing video: %s\n", workerID, req.VideoName)
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) // Simulate processing
		req.Response <- fmt.Sprintf("[Worker %d] Completed video: %s", workerID, req.VideoName)
	}
}

// Supervisor manages workers
func Supervisor(requests chan VideoRequest, numWorkers int) chan string {
	responseChan := make(chan string, 10)
	var wg sync.WaitGroup

	// Start worker actors
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go VideoWorker(i, requests, &wg)
	}

	// Close the response channel when all workers are done
	go func() {
		wg.Wait()
		close(responseChan)
	}()

	return responseChan
}

func Supervised_Actor() {
	rand.Seed(time.Now().UnixNano())

	requests := make(chan VideoRequest, 10)
	responseChan := Supervisor(requests, 3) // Start 3 workers

	videos := []string{"Intro.mp4", "Scene1.mp4", "Scene2.mp4", "Scene3.mp4"}
	for i, v := range videos {
		requests <- VideoRequest{ID: i, VideoName: v, Response: responseChan}
	}
	close(requests) // No more messages

	// Collect responses
	for res := range responseChan {
		fmt.Println(res)
	}

	fmt.Println("All videos processed. ✅")
}
