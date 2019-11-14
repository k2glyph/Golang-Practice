package main

import "fmt"

func main() {
	messages := make(chan string, 20)
	messages <- "HI"
	messages <- "hellow"
	messages <- "hellow2"
	for {
		fmt.Println(<-messages)
		// msg := <-messages
		// if !msg {
		// 	fmt.Println(msg)
		// }
	}

}
