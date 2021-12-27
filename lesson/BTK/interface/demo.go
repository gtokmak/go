package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s Shape) {
	fmt.Println(s.area())
}
func main() {
	r := rectangle{width: 3, height: 5}
	school(r)

}
