package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a new context
	context, _ := zmq.NewContext()
	defer context.Term()

	// Create a new socket of type REQ
	socket, _ := context.NewSocket(zmq.REQ)
	defer socket.Close()

	fmt.Printf("Connecting to hello world server...\n")
	socket.Connect("tcp://localhost:5555")

	for i := 0; i < 10; i++ {
		// Send hello message
		msg := fmt.Sprintf("Hello %d", i)
		socket.Send(msg, 0)
		fmt.Println("Sending", msg)

		// Wait for reply
		reply, _ := socket.Recv(0)
		fmt.Println("Received", reply)
	}
}
