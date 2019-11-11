package main

import "fmt"

func main() {

	x := 15
	a := &x
	fmt.Println(a)
	fmt.Println(*a)
}
