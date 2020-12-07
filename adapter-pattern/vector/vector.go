package vector

// Vector - Point in vector space
type Vector struct {
	X, Y float64
}

// NewVector - returns a pointer to a new vector
func NewVector(x, y float64) *Vector {
	return &Vector{X: x, Y: y}
}

// Line - Line in vector space
type Line struct {
	X1, Y1, X2, Y2 float64
}

// Image - One to many lines in vector space constituing an image
type Image struct {
	Lines []Line
}

// NewRectangle - returns an image of a rectangle from the given origin
func NewRectangle(width, height float64, origin *Vector) *Image {
	if origin == nil {
		origin = NewVector(0, 0)
	}
	return &Image{
		[]Line{
			{origin.X, origin.Y, origin.X + width, origin.Y},
			{origin.X, origin.Y, origin.X, origin.Y + height},
			{origin.X + width, origin.Y, origin.X + width, origin.Y + height},
			{origin.X, origin.Y + height, origin.X + width, origin.Y + height},
		},
	}
}
