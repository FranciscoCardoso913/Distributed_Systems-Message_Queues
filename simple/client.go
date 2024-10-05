package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a new context
	context, _ := zmq.NewContext()
	defer context.Term()

	// Create and connect socket
	socket := createAndConnectSocket(context)

	for i := 0; i < 10; i++ {
		// Send hello message
		msg := fmt.Sprintf("Hello %d", i)
		_, err := socket.Send(msg, 0)
		if err != nil {
			fmt.Println("Error sending message, recreating socket:", err)
			socket = recreateSocket(context, socket) // Recreate socket
			continue
		}
		fmt.Println("Sending", msg)

		// Poll for a reply with timeout
		reply, err := receiveReplyWithTimeout(socket, 2*time.Second)
		if err != nil {
			fmt.Println("No response from server, recreating socket:", err)
			socket = recreateSocket(context, socket) // Recreate socket
		} else {
			fmt.Println("Received", reply)
		}
	}
}

// Create and connect the socket
func createAndConnectSocket(context *zmq.Context) *zmq.Socket {
	socket, _ := context.NewSocket(zmq.REQ)
	socket.Connect("tcp://localhost:5555")
	fmt.Println("Connected to server")
	return socket
}

// Recreate the socket in case of error
func recreateSocket(context *zmq.Context, oldSocket *zmq.Socket) *zmq.Socket {
	// Close the old socket
	oldSocket.Close()

	// Create and connect a new socket
	time.Sleep(1 * time.Second) // Sleep before trying again
	return createAndConnectSocket(context)
}

// Receive reply with timeout
func receiveReplyWithTimeout(socket *zmq.Socket, timeout time.Duration) (string, error) {
	poller := zmq.NewPoller()
	poller.Add(socket, zmq.POLLIN)

	polledSockets, err := poller.Poll(timeout)
	if err != nil {
		return "", err
	}

	if len(polledSockets) > 0 {
		reply, err := socket.Recv(0)
		if err != nil {
			return "", err
		}
		return reply, nil
	}

	return "", fmt.Errorf("timeout after %v", timeout)
}
