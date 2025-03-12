package main

import (
	"fmt"
	"math/rand"
	"time"
)


type VideoRequest struct {
	ID        int
	VideoName string
	Response  chan string 
}


func VideoWorkerActor(requests chan VideoRequest) {
	for req := range requests {
		fmt.Printf("[Worker] Processing video: %s\n", req.VideoName)
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) 
		result := fmt.Sprintf("[Worker] Completed video: %s", req.VideoName)
		req.Response <- result 
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	
	requests := make(chan VideoRequest, 5)

	
	go VideoWorkerActor(requests)

	
	responses := make(chan string, 3) 
	videos := []string{"Intro.mp4", "Scene1.mp4", "Scene2.mp4"}

	for i, v := range videos {
		requests <- VideoRequest{ID: i, VideoName: v, Response: responses}
	}

	
	for i := 0; i < len(videos); i++ {
		fmt.Println(<-responses)
	}

	close(requests)  
	close(responses) 
}
