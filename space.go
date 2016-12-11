package simpartsim

import "math"

// Space ...
type Space interface {
	Run(ps Particles, frames int, dump bool)
	Termination() Coords
}

// SimpleSpaceOptions ...
type SimpleSpaceOptions struct {
	FrameLen float64
	Size     float64
	Gravity  float64
	Drag     float64
}

func (o *SimpleSpaceOptions) normalize() {
	if o.FrameLen == 0 {
		o.FrameLen = .1
	}
	if o.Size == 0 {
		o.Size = 100.0
	}
	if o.Gravity == 0 {
		o.Gravity = 9.81
	}
	if o.Drag == 0 {
		o.Drag = 9.0
	}
}

// SimpleSpace ...
type SimpleSpace struct {
	origin, termination     Coords
	frameLen, gravity, drag float64
}

// NewSimpleSpace ...
func NewSimpleSpace(opts SimpleSpaceOptions) *SimpleSpace {
	opts.normalize()

	return &SimpleSpace{
		origin:      Coords{0, 0, 0},
		termination: Coords{opts.Size, opts.Size, opts.Size},
		frameLen:    opts.FrameLen,
		gravity:     opts.Gravity,
		drag:        opts.Drag,
	}
}

// Termination ...
func (s *SimpleSpace) Termination() Coords {
	return s.termination
}

// incrementDisplacement ...
func (s *SimpleSpace) incrementDisplacement(p Particle, time, gravity float64) {
	pt, v := p.Point(), p.Vector()

	if pt.y == 0 && v.y == 0 {
		return
	}

	newV := v
	newV.y = v.y - gravity*time

	p.SetVector(newV)
}

// drag ...
func (s *SimpleSpace) incrementDrag(p Particle, time, drag float64) {
	pt, v := p.Point(), p.Vector()

	if pt.y != 0 && v.y != 0 {
		return
	}

	newV := v

	if math.Abs(v.x) < .05 {
		newV.x = 0
	}
	if math.Abs(v.z) < .05 {
		newV.z = 0
	}

	if newV.x != 0 {
		newV.x = v.x * drag * time
	}
	if newV.z != 0 {
		newV.z = v.z * drag * time
	}

	p.SetVector(newV)
}

// incrementLocation ...
func (s *SimpleSpace) incrementLocation(p Particle, time float64) {
	pt, v := p.Point(), p.Vector()
	newPt := pt

	newPt.x = pt.x + v.x*time
	newPt.y = pt.y + v.y*time
	newPt.z = pt.z + v.z*time

	p.SetPoint(newPt)
}

// reflect ...
func (s *SimpleSpace) reflect(a, b float64) float64 {
	diff := a - b
	return b - diff
}

// processCollisions ...
func (s *SimpleSpace) processCollisions(p Particle, origin, termination Coords) {
	pt, v := p.Point(), p.Vector()
	newPt, newV := pt, v

	if pt.x <= origin.x {
		newPt.x = s.reflect(pt.x, origin.x)
		newV.x = v.x * -.7
	}
	if pt.x >= termination.x {
		newPt.x = s.reflect(pt.x, termination.x)
		newV.x = v.x * -.7
	}

	if pt.y <= origin.y {
		newPt.y = s.reflect(pt.y, origin.y)
		newV.y = v.y * -.5

		if newPt.y < .4 && newV.y-newPt.y < 1 {
			newPt.y = 0
			newV.y = 0
		}
	}
	if pt.y >= termination.y {
		newPt.y = s.reflect(pt.y, termination.y)
		newV.y = v.y * -.9
	}

	if pt.z <= origin.z {
		newPt.z = s.reflect(pt.z, origin.z)
		newV.z = v.z * -.7
	}
	if pt.z >= termination.z {
		newPt.z = s.reflect(pt.z, termination.z)
		newV.z = v.z * -.7
	}

	p.SetPoint(newPt)
	p.SetVector(newV)
}

// increment ...
func (s *SimpleSpace) increment(p Particle) {
	s.incrementDisplacement(p, s.frameLen, s.gravity)
	s.incrementLocation(p, s.frameLen)
	s.processCollisions(p, s.origin, s.termination)
	s.incrementDrag(p, s.frameLen, s.drag)
}

// tick ...
func (s *SimpleSpace) tick(ps Particles) []Coords {
	d := ps.data()
	cs := make([]Coords, len(d))

	for k := range d {
		s.increment(d[k])
		cs[k] = d[k].Point()
	}

	return cs
}

// Run ...
func (s *SimpleSpace) Run(ps Particles, frames int, cs chan []Coords) {
	for i := 0; i < frames; i++ {
		cs <- s.tick(ps)
	}
}
