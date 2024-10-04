package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"time"
)

func main() {
	// Create a new context
	context, _ := zmq.NewContext()
	defer context.Term()

	// Create a new socket of type REP
	socket, _ := context.NewSocket(zmq.REP)
	defer socket.Close()

	// Bind the socket to an address
	socket.Bind("tcp://*:5555")

	// Wait for messages
	for {
		// Receive a message
		msg, _ := socket.Recv(0)
		fmt.Println("Received ", msg)

		// do some fake "work"
		time.Sleep(time.Second)

		// send reply back to client
		reply := "World"
		socket.Send(reply, 0)
	}
}
