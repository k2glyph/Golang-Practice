package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "test text 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println("Received message from channel 1:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout channel message didn't received exceed time")
	}
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "This is the message two"
	}()
	select {
	case msg := <-c2:
		fmt.Println("Received message from channel 2:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout exceed")
	}
}
