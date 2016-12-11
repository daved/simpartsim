package simpartsim

import "math/rand"

// Particle ...
type Particle interface {
	Point() Coords
	SetPoint(Coords)
	Vector() Coords
	SetVector(Coords)
}

// Particles ...
type Particles interface {
	data() []Particle
}

// SimpleParticle ...
type SimpleParticle struct {
	pt   Coords
	vctr Coords
}

// NewSimpleParticle ...
func NewSimpleParticle(id, ct float64, termination Coords) *SimpleParticle {
	t := termination

	return &SimpleParticle{
		pt:   Coords{id / ct * t.x, rand.Float64() * t.y, rand.Float64() * t.z},
		vctr: Coords{rand.Float64() * 5, 0, 0},
	}
}

// Point ...
func (p *SimpleParticle) Point() Coords {
	return p.pt
}

// SetPoint ...
func (p *SimpleParticle) SetPoint(cs Coords) {
	p.pt = cs
}

// Vector ...
func (p *SimpleParticle) Vector() Coords {
	return p.vctr
}

// SetVector ...
func (p *SimpleParticle) SetVector(cs Coords) {
	p.vctr = cs
}

// SimpleParticles ...
type SimpleParticles struct {
	d []Particle
}

// NewSimpleParticles ...
func NewSimpleParticles(ct int, termination Coords) *SimpleParticles {
	ps := &SimpleParticles{
		d: make([]Particle, ct),
	}

	for i := 0; i < ct; i++ {
		ps.d[i] = NewSimpleParticle(float64(i), float64(ct), termination)
	}

	return ps
}

// data ...
func (ps *SimpleParticles) data() []Particle {
	return ps.d
}
