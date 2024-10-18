package main

import (
	"fmt"
	"log"

	zmq "github.com/pebbe/zmq4"
)

func main() {


	subscriber, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		log.Fatal(err)
	}
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5560")
	subscriber.SetSubscribe("")
	// Setup poller
	poller := zmq.NewPoller()
	poller.Add(subscriber, zmq.POLLIN)

	// Process messages from both sockets
	for {
		polled, err := poller.Poll(-1)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range polled {
			switch socket := item.Socket; socket {
			case subscriber:
				// Process weather update
				msg, err := subscriber.Recv(0)
				if err != nil {
					log.Println("Error receiving from subscriber:", err)
				} else {
					fmt.Println("Received USA weather update:", msg)
				}
			}
		}
	}

	fmt.Println("done")
}
