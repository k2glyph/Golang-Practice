package main

import "fmt"

func zeroVal(i int) {
	i = 0
}
func zeroPointer(i *int) {
	*i = 0
}
func main() {
	var i int = 1
	fmt.Println("initial value of i: ", i)
	zeroVal(i)
	fmt.Println("Zero value: ", i)
	zeroPointer(&i)
	fmt.Println("Zero pointer", i)
	fmt.Println("zero Pointer", &i)
}
