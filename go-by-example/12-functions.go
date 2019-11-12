package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Multi return functions
func values() (int, int, int) {
	return 1, 2, 3
}

// variadic function
/*
Variadic functions can be called with any number of trailing arguments.
For example, fmt.Println is a common variadic function.
*/
func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/*
Go supports anonymous functions, which can form closures.
Anonymous functions are useful when you want to define a function inline without having to name it.
*/

/*
Go supports recursive functions. Here’s a classic factorial example.
*/
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	fmt.Println(plus(10, 20))
	fmt.Println(plusPlus(10, 20, 30))
	fmt.Println(values())
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	values := []int{1, 2, 3, 4, 5, 6, 7, 929309, 4959, 607}
	fmt.Println(sum(values...))
	nextSeq := intSeq()
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	newSeq := intSeq()
	fmt.Println(newSeq())
	fmt.Println(newSeq())
	fmt.Println(newSeq())
	fmt.Println(factorial(5))
	fmt.Println(factorial(0))
	fmt.Println(factorial(7))
}
