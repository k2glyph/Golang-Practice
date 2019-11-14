package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan string)
	select {
	case msg := <-messages:
		fmt.Println("Received message: ", msg)
	default:
		fmt.Println("No message received")
	}
	msg := "Hi"
	select {
	case messages <- msg:
		fmt.Println("Send messages to messages channel")
	default:
		fmt.Println("No message sent")
	}
	select {
	case msg := <-messages:
		fmt.Println("Received messages: ", msg)
	case sig := <-signals:
		fmt.Println("Received signals:", sig)
	default:
		fmt.Println("No activity")
	}
}
