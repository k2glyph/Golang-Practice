package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	perimeter() float32
}
type rect struct {
	width, height float32
}

func (r rect) area() float32 {
	return r.width * r.height
}
func (r rect) perimeter() float32 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println("Geometry:", g)
	fmt.Println("Area:=", g.area())
	fmt.Println("Perimeter=", g.perimeter())
}
func main() {
	r := rect{width: 2, height: 3}
	measure(r)
	c := circle{radius: 2}
	measure(c)
}
