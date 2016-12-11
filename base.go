package simpartsim

import "strconv"

const (
	// CoordsFieldNames ...
	CoordsFieldNames = "X Axis,Y Axis,Z Axis"
)

// Coords ...
type Coords struct {
	X, Y, Z float64
}

// String ...
func (c *Coords) String() string {
	s := strconv.FormatFloat(c.X, 'f', 12, 64) + ","
	s += strconv.FormatFloat(c.Y, 'f', 12, 64) + ","
	s += strconv.FormatFloat(c.Z, 'f', 12, 64)

	return s
}
