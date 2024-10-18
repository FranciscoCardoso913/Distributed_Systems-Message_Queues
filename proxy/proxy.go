// Simple request-reply broker
//
// Author:  Brendan Mc.
// Requires: http://github.com/alecthomas/gozmq

package main

import (
	zmq "github.com/pebbe/zmq4"
)

func main() {
	context, _ := zmq.NewContext()

	frontend, _ := context.NewSocket(zmq.XSUB)
	backend, _ := context.NewSocket(zmq.XPUB)
	defer frontend.Close()
	defer backend.Close()
	frontend.Connect("tcp://localhost:5559")
	backend.Bind("tcp://*:5560")

	// Initialize poll set
	poller := zmq.NewPoller()
	poller.Add(frontend, zmq.POLLIN)
	poller.Add(backend, zmq.POLLIN)

	for {
		// Poll sockets
		polled, _ := poller.Poll(-1)

		// Check for activity on frontend socket
		for _, item := range polled {
			switch socket := item.Socket; socket {
			case frontend:
				// If frontend is ready, receive and forward to backend
				parts, _ := frontend.RecvMessage(0)

				backend.SendMessage(parts)

			case backend:
				// If backend is ready, receive and forward to frontend
				parts, _ := backend.RecvMessage(0)

				frontend.SendMessage(parts)
			}
		}
	}
}
