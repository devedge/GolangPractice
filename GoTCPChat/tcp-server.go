package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081") // listen on all interfaces
	conn, _ := ln.Accept()              // accept connection on port

	// loop forever, until ctrl-c
	for {
		// listen for message to process, ending in newline
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output received message
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
