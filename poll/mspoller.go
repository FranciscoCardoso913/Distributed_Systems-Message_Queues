package main

import (
	"fmt"
	"log"

	zmq "github.com/pebbe/zmq4"
)

func main() {

	// Create zmq sockets
	receiver, err := zmq.NewSocket(zmq.PULL)
	if err != nil {
		log.Fatal(err)
	}
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	subscriber, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		log.Fatal(err)
	}
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5556")
	subscriber.SetSubscribe("")

	subscriberpt, err2 := zmq.NewSocket(zmq.SUB)
	if err2 != nil {
		log.Fatal(err)
	}
	defer subscriberpt.Close()
	subscriberpt.Connect("tcp://localhost:5559")
	subscriberpt.SetSubscribe("")

	// Setup poller
	poller := zmq.NewPoller()
	poller.Add(receiver, zmq.POLLIN)
	poller.Add(subscriber, zmq.POLLIN)
	poller.Add(subscriberpt, zmq.POLLIN)

	// Process messages from both sockets
	for {
		polled, err := poller.Poll(-1)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range polled {
			switch socket := item.Socket; socket {
			case receiver:
				fmt.Println("rec")
				msg, err := receiver.Recv(0)
				if err != nil {
					log.Println("Error receiving from receiver:", err)
				} else {
					fmt.Println("Received task:", msg)
				}
			case subscriber:
				// Process weather update
				msg, err := subscriber.Recv(0)
				if err != nil {
					log.Println("Error receiving from subscriber:", err)
				} else {
					fmt.Println("Received USA weather update:", msg)
				}

			case subscriberpt:
				// Process weather update
				msg, err := subscriberpt.Recv(0)
				if err != nil {
					log.Println("Error receiving from subscriber:", err)
				} else {
					fmt.Println("Received Portugal weather update:", msg)
				}
			}
		}
	}

	fmt.Println("done")
}
