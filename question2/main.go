package main

import "fmt"

func main() {
	// slice containing values to send to channel
	values := []int{10, 20, 35, 100, 200, 502}
	c := make(chan int, len(values))
	done := make(chan bool)

	go channelReader(c, done)

	for _, v := range values {
		c <- v // sending value to channel
	}
	close(c)
	<-done // wait for channel reader to finish
}

// Channel reader go-routine
func channelReader(c <-chan int, done chan<- bool) {
	for v := range c {
		fmt.Println("got value : ", v)
	}
	done <- true
}
