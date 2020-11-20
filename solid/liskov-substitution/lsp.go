package main

import "fmt"

type Sized interface {
	GetWidth() float64
	SetWidth(width float64)
	GetHeight() float64
	SetHeight(height float64)
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) GetWidth() float64 {
	return r.width
}

func (r *Rectangle) SetWidth(width float64) {
	r.width = width
}

func (r *Rectangle) GetHeight() float64 {
	return r.height
}

func (r *Rectangle) SetHeight(height float64) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size float64) *Square {
	return &Square{Rectangle{width: size, height: size}}
}

func (s *Square) SetWidth(width float64) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height float64) {
	s.width = height
	s.height = height
}

func Area(s Sized) float64 {
	return s.GetWidth() * s.GetHeight()
}

func main() {
	rect := &Rectangle{width: 3, height: 2}
	fmt.Printf("rectangle area: %.2f\n", Area(rect))

	square := NewSquare(5)
	fmt.Printf("square area: %.2f\n", Area(square))

	square.SetHeight(7)
	fmt.Printf("square area: %.2f\n", Area(square))
}
