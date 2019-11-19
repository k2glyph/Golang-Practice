package main

import(
	"fmt"
)
type point struct {
	x,y int
}
func main() {
	p:=point{x: 10, y:12}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n",p)
	fmt.Printf("%#v\n",p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("Binary %b\n", 10)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n", 10)
	fmt.Printf("%f\n", 8.8)
}