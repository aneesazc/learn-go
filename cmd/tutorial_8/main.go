package main

// channels are used to communicate between goroutines

import "fmt"

func main() {
	// create a channel
	c := make(chan string)

	// send a value to the channel
	go process(c)

	msg := <-c
	fmt.Println(msg)
}

func process(c chan string) {
	c <- "Hello, World!"
}
