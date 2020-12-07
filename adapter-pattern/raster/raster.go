package raster

import "strings"

// Point - element in raster space
type Point struct {
	X, Y int
}

// Image - set of points that constitute an image
type Image interface {
	GetPoints() []Point
}

// DrawPoints - rasterize points to space
func DrawPoints(img Image) string {
	maxX, maxY := 0, 0
	points := img.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX++
	maxY++

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}
