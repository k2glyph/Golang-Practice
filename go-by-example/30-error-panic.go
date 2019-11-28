package main

import (
	"fmt"
)

type customError struct {
	arg     string
	problem string
}

func (err *customError) Error() string {
	return fmt.Sprintf("%s: %s", err.arg, err.problem)
}
func f1(x int) (int, error) {
	if x == 42 {
		return -1, &customError{arg: "integer", problem: "Can't do anything because number is 42"}
	}
	return x, nil
}
func main() {
	fmt.Println("Panic Error example")
	for i := 0; i < 50; i = i + 2 {
		if val, e := f1(i); e == nil {
			fmt.Println("f1 works", val)
		} else {
			// fmt.Println(e.Error())
			panic(e.Error()) 
		}
	}

}
