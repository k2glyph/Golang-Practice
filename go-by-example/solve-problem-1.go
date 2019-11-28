package main

import "fmt"

// func foo(a, b int) {
// 	fmt.Println("parameter are:", a, b)
// }
// func foo(a, b, c int) {
// 	fmt.Println("parameter are:", a, b, c)
// }
// func foo(params ...int) {
// 	fmt.Printf("Parameter are:")
// 	for param := range params {
// 		fmt.Printf("%d", param)
// 	}
// }

// func foo() {
// 	fmt.Println("Foo called")
// }
func main() {
	// foo(1, 2)
	// foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
	func foo(any interface{}) {
		switch v:=any.(type) {
		case params ...int:
			fmt.Printf("Parameter are:")
			for param := range params {
				fmt.Printf("%d", param)
			}
		default:
			fmt.Println("Foo called")
			return 
		}
	}
}
