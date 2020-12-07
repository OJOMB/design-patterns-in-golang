/*
Imagine that we're creating a graphical package that has raster AND vector type images
we want all the good stuff in there like circles, rectangles, triangles and lines etc
if we designed this package badly we'd end up with something like

							                         Shape
	__________________________________________________|___________________________________________________
	|		                         |                           |                                       |
	Circle                          Rect                      Triangle                               Pentagon
____|___________________        _____|__________          _______|________________            ___________|_________
|                      |        |               |         |                       |           |                   |
RasterCircle   VectorCircle    RasterRect   VectorRect   RasterTriangle    VectorTriangle    RasterPentagon     VectorPentagon

4 shapes and 2 different renderers so we get this cartesian product explosion of types
we're combining all permutations to end up with 8 subtypes

*/

package main

import "fmt"

// Point - point in cartesian space
type Point struct {
	X, Y float64
}

// Renderer - render interface
type Renderer interface {
	RenderCircle(radius float64, origin Point)
	RenderSquare(length float64, origin Point)
}

// VectorRenderer - vector type rendering
type VectorRenderer struct{}

// RenderCircle - renders a circle
func (v *VectorRenderer) RenderCircle(radius float64, origin Point) {
	fmt.Printf("circle with radius: %.2f rendered at %+v\n", radius, origin)
}

// RenderSquare - renders a square
func (v *VectorRenderer) RenderSquare(length float64, origin Point) {
	fmt.Printf("square with length: %.2f rendered at %+v\n", length, origin)
}

// Shape - base type for all shapes
type Shape struct {
	Name   string
	Sides  int
	Origin Point
}

// MakeShape - returns a pointer to a new Shape
func MakeShape(sides int, name string, originX, originY float64) Shape {
	return Shape{
		Name:   name,
		Sides:  sides,
		Origin: Point{X: originX, Y: originY},
	}
}

// Circle - one sided shape with infinite degrees of symmetry
type Circle struct {
	Shape
	renderer Renderer
	radius   float64
}

// NewCircle - constructor for a circle
func NewCircle(r Renderer, radius, originX, originY float64) *Circle {
	return &Circle{
		MakeShape(1, "circle", originX, originY),
		r,
		radius,
	}
}

// Draw - call renderer to draw circle
func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius, c.Origin)
}

// Square - 4 sided shape with sides of equal length
type Square struct {
	Shape
	renderer Renderer
	length   float64
}

// NewSquare - constructor for a square
func NewSquare(r Renderer, length, originX, originY float64) *Square {
	return &Square{
		MakeShape(4, "square", originX, originY),
		r,
		length,
	}
}

// Draw - call renderer to draw circle
func (c *Square) Draw() {
	c.renderer.RenderSquare(c.length, c.Origin)
}

func main() {
	r := &VectorRenderer{}
	sq := NewSquare(r, 4., 1., 1.)
	sq.Draw()

	c := NewCircle(r, 10, 0., 0.)
	c.Draw()
}
