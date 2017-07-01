package main

import "fmt"
import "time"

func main() {

	// Simulate a bunch of incoming requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Limiter channel that will recieve a value every 200ms
	// This is the regulator
	limiter := time.Tick(time.Millisecond * 200)

	// Handling the flood of requests
	for req := range requests {
		// Block on each 'limiter' channel event, then handle the request
		<- limiter
		fmt.Println("request", req, time.Now())
	}



	// Simulate a bursty limiter to allow short bursts of requests
	// This one allows bursts of up to 3 events
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200ms, try to add a new value to 'burstyLimiter' up
	// to its limit of 3
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requests, where the first 3 will
	// benefit from the burst capacity of the limiter
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
