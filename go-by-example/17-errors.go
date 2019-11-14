package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't  work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg     int
	problem string
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, problem: "Can't work with it"}
	}
	return arg + 5, nil
}
func (e *argError) Error() string {
	return fmt.Sprintf("%d, %s", e.arg, e.problem)
}
func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 works: ", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 works:", r)
		}
	}
	_, err := f2(42)
	if ae, ok := err.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.problem)
	}

}
