package main

import "fmt"

type rect struct {
	l, b int
}

/*
	Area of rectange = l*b
*/
func (r *rect) area() int {
	return r.l * r.b
}

/*
Perimeter of rectangle =2(l+b)
*/
func (r *rect) perimeter() int {
	return 2 * (r.l + r.b)
}
func main() {
	r := rect{2, 3}
	fmt.Println("Area of rectange 2 by 3=", r.area())
	fmt.Println("Perimeter of rectange of 2 by 3=", r.perimeter())
}
