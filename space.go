package simpartsim

import "math"

// Space ...
type Space interface {
	Run(ps Particles, frames int, cs chan []Coords)
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

	if pt.Y == 0 && v.Y == 0 {
		return
	}

	newV := v
	newV.Y = v.Y - gravity*time

	p.SetVector(newV)
}

// drag ...
func (s *SimpleSpace) incrementDrag(p Particle, time, drag float64) {
	pt, v := p.Point(), p.Vector()

	if pt.Y != 0 && v.Y != 0 {
		return
	}

	newV := v

	if math.Abs(v.X) < .05 {
		newV.X = 0
	}
	if math.Abs(v.Z) < .05 {
		newV.Z = 0
	}

	if newV.X != 0 {
		newV.X = v.X * drag * time
	}
	if newV.Z != 0 {
		newV.Z = v.Z * drag * time
	}

	p.SetVector(newV)
}

// incrementLocation ...
func (s *SimpleSpace) incrementLocation(p Particle, time float64) {
	pt, v := p.Point(), p.Vector()
	newPt := pt

	newPt.X = pt.X + v.X*time
	newPt.Y = pt.Y + v.Y*time
	newPt.Z = pt.Z + v.Z*time

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

	if pt.X <= origin.X {
		newPt.X = s.reflect(pt.X, origin.X)
		newV.X = v.X * -.7
	}
	if pt.X >= termination.X {
		newPt.X = s.reflect(pt.X, termination.X)
		newV.X = v.X * -.7
	}

	if pt.Y <= origin.Y {
		newPt.Y = s.reflect(pt.Y, origin.Y)
		newV.Y = v.Y * -.5

		if newPt.Y < .4 && newV.Y-newPt.Y < 1.2 {
			newPt.Y = 0
			newV.Y = 0
		}
	}
	if pt.Y >= termination.Y {
		newPt.Y = s.reflect(pt.Y, termination.Y)
		newV.Y = v.Y * -.9
	}

	if pt.Z <= origin.Z {
		newPt.Z = s.reflect(pt.Z, origin.Z)
		newV.Z = v.Z * -.7
	}
	if pt.Z >= termination.Z {
		newPt.Z = s.reflect(pt.Z, termination.Z)
		newV.Z = v.Z * -.7
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
