package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y float64) float64 {
	return x + y
}
func multiple(a, b string) (string, string) {
	return a, b
}
func main() {
	fmt.Println("Welcome to Go!")
	fmt.Println("The square root of 4 is:", math.Sqrt(4))
	fmt.Println("The random number from 0-100:", rand.Intn(100))
	num1, num2 := 5.55, 6.10
	fmt.Println(add(num1, num2))
	a, b := "Hey", "there"
	fmt.Println(multiple(a, b))
}
