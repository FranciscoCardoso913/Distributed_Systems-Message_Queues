// Hello World client
// Connects REQ socket to tcp://localhost:5559
// Sends "Hello" to server, expects "World" back
//
// Author:  Brendan Mc.
// Requires: http://github.com/alecthomas/gozmq

package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	context, _ := zmq.NewContext()

	// Socket to talk to clients
	requester, _ := context.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect("tcp://localhost:5559")

	for i := 0; i < 50; i++ {
		requester.Send("Hello", 0)
		reply, _ := requester.Recv(0)
		fmt.Printf("Received reply %d [%s]\n", i, reply)
	}
}
