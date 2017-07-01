package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// These encapsulate the requests and a way
// for the goroutines to respond
type readOp struct {
	key int
	resp chan int
}
type writeOp struct {
	key int
	val int
	resp chan bool
}


// See mutexes.go for a non-Golang-style implementation
func main() {

	var readOps uint64 = 0
  var writeOps uint64 = 0

	// Channels to perform reads & writes
  reads := make(chan *readOp)
  writes := make(chan *writeOp)

	// The goroutine that owns the state.
	// In this example, the state is a 'map'
  go func() {
  	var state = make(map[int]int)

		// Handle read and write commands into the state
  	for {
  		select {
  			case read := <- reads:
  				read.resp <- state[read.key]
  			case write := <- writes:
  				state[write.key] = write.val
  				write.resp <- true
  		}
  	}
  }()


  // Create 100 random reader goroutines
  for r := 0; r < 100; r++ {
  	go func() {
  		for {
  			read := &readOp{
  					key: rand.Intn(5),
  					resp: make(chan int)}
  			reads <- read // pipe in a read request
  			<- read.resp  // block for the read response
  			atomic.AddUint64(&readOps, 1)
  			time.Sleep(time.Millisecond)
  		}
  	}()
  }

  // Create 10 random writer goroutines
  for w := 0; w < 10; w++ {
  	go func() {
  		for {
  			write := &writeOp{
  					key: rand.Intn(5),
  					val: rand.Intn(100),
  					resp: make(chan bool)}
  			writes <- write
  			<- write.resp
  			atomic.AddUint64(&writeOps, 1)
  			time.Sleep(time.Millisecond)
  		}
  	}()
  }

  // Let the goroutines run for a full second
  time.Sleep(time.Second)

  readOpsFinal := atomic.LoadUint64(&readOps)
  fmt.Println("readOps:", readOpsFinal)

  writeOpsFinal := atomic.LoadUint64(&writeOps)
  fmt.Println("writeOps:", writeOpsFinal)
}
