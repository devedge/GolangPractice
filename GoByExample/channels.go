package main

import "fmt"

func main() {
	messages := make(chan string)

	// Fire a 'ping' into the 'messages' channel
	go func() { messages <- "ping" }()

	// Set 'msg' to the value that will be sent through
	// the 'messages' channel
	msg := <- messages
	fmt.Println(msg)
}
