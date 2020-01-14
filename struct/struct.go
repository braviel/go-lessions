package structs

import "math"

//Rectangle represent rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle represent circle shape
type Circle struct {
	Radius float64
}

//Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

//Shape interface
type Shape interface {
	Area() float64
}

//Perimeter calc perimeter of shape
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

//Area calc area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Area calc area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Area calc Triangle Area
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2.0
}
