package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/OJOMB/design-patterns-in-golang/adapter-pattern/raster"
	"github.com/OJOMB/design-patterns-in-golang/adapter-pattern/vector"
)

func hash(obj interface{}) [16]byte {
	bytes, _ := json.Marshal(obj)
	return md5.Sum(bytes)
}

func minmax(a, b float64) (float64, float64) {
	if a < b {
		return a, b
	}
	return b, a
}

// the issue is that we can define Vectors/Lines and Images in terms of the types given to us by the vector library
// but then when it comes to drawing, we only have the raster lib which requires a raster style image
// what we need is an adapter that can take vector images and convert them to raster images to be consumed by the raster draw function

type vector2raster struct {
	points []raster.Point
}

var pointCache = map[[16]byte][]raster.Point{}

func (v2r *vector2raster) generatePoints(line vector.Line) (points []raster.Point) {
	left, right := minmax(line.X1, line.X2)
	bottom, top := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := top - bottom

	if dx == 0 {
		for y := bottom; y < top; y++ {
			points = append(points, raster.Point{X: int(left), Y: int(y)})
		}
	} else if dy == 0 {
		for x := left; x < right; x++ {
			points = append(points, raster.Point{X: int(x), Y: int(top)})
		}
	} else {
		// TODO
	}
	return points
}

func (v2r *vector2raster) addLine(line vector.Line) {
	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		v2r.points = append(v2r.points, pts...)
		return
	}
	points := v2r.generatePoints(line)
	v2r.points = append(v2r.points, points...)

	// be sure to add these to the cache
	pointCache[h] = points
}

func (v2r vector2raster) GetPoints() []raster.Point {
	return v2r.points
}

// VectorToRaster - converts a vector image to a raster Image
func VectorToRaster(vi *vector.Image) raster.Image {
	adapter := vector2raster{}
	for _, line := range vi.Lines {
		adapter.addLine(line)
	}

	return adapter // as RasterImage
}

func main() {
	rc := vector.NewRectangle(25., 14., nil)
	a := VectorToRaster(rc) // adapter!
	fmt.Print(raster.DrawPoints(a))
}
