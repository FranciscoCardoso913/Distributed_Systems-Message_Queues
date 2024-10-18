// 
//   Weather update server
//   Binds PUB socket to tcp://*:5556
//   Publishes random weather updates
// 
package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"math/rand"
	"time"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.PUB)
	defer context.Term()
	defer socket.Close()
	socket.Bind("tcp://*:5559")
	socket.Bind("ipc://weather.ipc")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// loop for a while aparently
	for {

		//  make values that will fool the boss
		zipcode := rand.Intn(100000)
		temperature := rand.Intn(215) - 80
		relhumidity := rand.Intn(50) + 10

		msg := fmt.Sprintf("%d %d %d", zipcode, temperature, relhumidity)

		//  Send message to all subscribers
		socket.Send(msg, 0)
	}
}
