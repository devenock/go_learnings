package gochannels

import "fmt"

// Channels are the pipes that connect concurrent goroutines
// You can send values into channels from one goroutine and receive those values into another goroutine.
func GoChannels() {
	fmt.Println("Go Channels!!")

	// Create a new channel with make(chan val-type)
	messages := make(chan string)

	// Channels are typed by the values they convey.

	// send a value to the channel using the (<-) syntax
	go func() { messages <- "ping" }()
	msg := <-messages

	fmt.Println(msg, messages)

	// Channel Buffering
	// By default, channels are unbuffered, meaning that they will only accept sends(chan <-)
	// if there is a corresponding receive(<- chan) ready to receive the sent value.
	// Buffered channels accept a limited number of values without a corresponding receiver for those values.
	students := make(chan string, 3)
	// Because the channel is buffered, we can send these values into the channel without a corresponding concurent receive.
	students <- "buffered"
	students <- "channel"
	fmt.Println(<-students)
	fmt.Println(<-students)
}
