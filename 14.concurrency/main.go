package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go greet("John", done)
	go greet("Jane", done)
	go slowGreet("Jimmy", done)
	go greet("Jill", done)
	// wait for the done channel to receive a message and then print a message
	// block the main goroutine until the done channel receives a message
	// so basically iy will call hello jimmy because it will wait for the done channel to receive a message
	<-done
	<-done
	<-done
	<-done
	// we put many <-done because we want to receive the message from the done channel many times, otherwise it will have race condition which only print out one message
	fmt.Println("Done")

	// this is the same as the above but with a for loop
	dones := make([]chan bool, 4)
	for i := 0; i < 4; i++ {
		dones[i] = make(chan bool)
	}
	go greet("John", dones[0])
	go greet("Jane", dones[1])
	go slowGreet("Jimmy", dones[2])
	go greet("Jill", dones[3])
	for i := 0; i < 4; i++ {
		<-dones[i]
	}
	fmt.Println("Done")

	// there is another way to do this, which is to close channel
	// doneChannel := make(chan bool)
	// go greet("John", doneChannel)
	// go greet("Jane", doneChannel)
	// go slowGreet("Jimmy", doneChannel)
	// go greet("Jill", doneChannel)
	// // this will print out the message from the doneChannel until it is closed
	// for name := range doneChannel {
	// 	fmt.Println("Hello,", name)
	// }

}

func greet(name string, doneChannel chan bool) {
	fmt.Println("Hello,", name)
	doneChannel <- true
}

func slowGreet(name string, doneChannel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello,", name)
	// send a message to the channel, which will be received by the main goroutine
	doneChannel <- true
	// close(doneChannel) // explicitly close the channel this is for whenever you know the channel is not going to be used anymore
}
