package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {

	// Create zmq sockets
	receiver, err := zmq.NewSocket(zmq.PULL)
	if err != nil {
		fmt.Println("Error creating receiver socket:", err)
		return
	}
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	subscriber, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		fmt.Println("Error creating subscriber socket:", err)
		return
	}
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5556")
	subscriber.SetSubscribe("10001")

	// Process messages from both sockets
	for {
		// ventilator
		for {
			b, err := receiver.Recv(zmq.DONTWAIT)
			if err != nil {
				break // Exit loop if there's no message
			}
			fmt.Printf("%s\n", string(b))
			// fake process task
		}

		// weather server
		for {
			b, err := subscriber.Recv(zmq.DONTWAIT)
			if err != nil {
				break // Exit loop if there's no message
			}
			// process task
			fmt.Printf("found weather = %s\n", string(b))
		}

		// No activity, so sleep for 1 msec
		time.Sleep(time.Millisecond)
	}

	fmt.Println("done")
}
