package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width  float64
	Height float64
}

type Circle struct {
	Diameter float64
}

func (r Circle) Area() float64 {
	return 100
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perimeter() float64 {
	return r.Width + r.Height
}

func PrintArea(s Shape) {
	fmt.Print(s.Area())
}

func PrintArea2(i interface{}) {
	s := i.(Shape)
	fmt.Println(s.Area())
}

// func (e *PathError)

func main() {
	aRect := Rect{Height: 100, Width: 200}
	PrintArea(aRect)
	PrintArea2(aRect)
}
