package main

import "fmt"

// Only accepts a channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Accepts a channel for recieves (pings) and 
// a second for sends (pongs)
func pong(pongs chan<- string, pings <-chan string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message") // fire the message into the 'pings' channel
	pong(pongs, pings) // redirect the 'pings' message and fire it into 'pongs'

	fmt.Println(<-pongs) // catch 'pongs' and prints it
}
