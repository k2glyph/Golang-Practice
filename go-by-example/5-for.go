package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 3 {
		fmt.Println("i value:", i)
		i++
	}
	for j := 3; j < 9; j++ {
		fmt.Println("j value:", j)
	}
	k := 0
	for {
		fmt.Println("looping infinite")
		if k > 5 {
			break
		}
		k++
	}
	for j := 0; j <= 100; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
}
