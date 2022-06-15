package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, c chan bool) {
	<-c // thread safe access to 'x'
	x = x + 1
	c <- true // unlocking 'x'
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	c := make(chan bool, 1) // channel to act as a mutex
	c <- true
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, c)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
