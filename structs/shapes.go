package structs

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base    float64
	Height  float64
	SideOne *float64
	SideTwo *float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (t Triangle) Area() float64 {
	return t.Height * t.Base / 2
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func (t Triangle) Perimeter() float64 {
	if t.SideOne == nil || t.SideTwo == nil {
		fmt.Print("no sides were provided to calculate perimeter of triangle")
		return 0
	}
	return t.Base + *t.SideOne + *t.SideTwo
}
