package simpartsim

import "strconv"

const (
	// CoordsFieldNames ...
	CoordsFieldNames = "X Axis,Y Axis,Z Axis"
)

// Coords ...
type Coords struct {
	x, y, z float64
}

// String ...
func (c *Coords) String() string {
	s := strconv.FormatFloat(c.x, 'f', 12, 64) + ","
	s += strconv.FormatFloat(c.y, 'f', 12, 64) + ","
	s += strconv.FormatFloat(c.z, 'f', 12, 64)

	return s
}
