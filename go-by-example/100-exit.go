package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome")
	defer fmt.Println("Exiting")
	os.Exit(3)
}