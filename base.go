package main

import "strconv"

// coords ...
type coords struct {
	x, y, z float64
}

// makeCoords ...
func makeCoords(x, y, z float64) coords {
	return coords{
		x: x,
		y: y,
		z: z,
	}
}

// String ...
func (c *coords) String() string {
	s := strconv.FormatFloat(c.x, 'f', 12, 64) + ","
	s += strconv.FormatFloat(c.y, 'f', 12, 64) + ","
	s += strconv.FormatFloat(c.z, 'f', 12, 64)

	return s
}

// point ...
type point struct {
	coords
}

// makePoint ...
func makePoint(x, y, z float64) point {
	return point{makeCoords(x, y, z)}
}

// vector ...
type vector struct {
	coords
}

// makeVector ...
func makeVector(x, y, z float64) vector {
	return vector{makeCoords(x, y, z)}
}
