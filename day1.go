package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint32 = 4
	// var d = 4.0
	var s = "Hacker Rank"

	var i2 int
	// var d2 float64
	var s2 string
	fmt.Printf("Default number1: %d The string: %s\n", i, s)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the Number2: ")
	fmt.Scan(&i2)
	// fmt.Scan(&d2)
	fmt.Print("Please enter the Text2: ")
	for scanner.Scan() {
		s2 = scanner.Text()
		break
	}
	fmt.Printf("Number1 and Number2: %d\n", i2+int(i))
	// fmt.Printf("%.1f \n", d2+d)
	fmt.Println("Text1 and Text2: " + s + " " + s2)
}
