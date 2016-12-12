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
	Reset()
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
		pt:   Coords{id / ct * t.X, rand.Float64() * t.Y, rand.Float64() * t.Z},
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
	u []Particle
}

// NewSimpleParticles ...
func NewSimpleParticles(ct int, termination Coords) *SimpleParticles {
	ps := &SimpleParticles{
		d: make([]Particle, ct),
		u: make([]Particle, ct),
	}

	for i := 0; i < ct; i++ {
		p := NewSimpleParticle(float64(i), float64(ct), termination)
		ps.d[i] = p
		pc := *p
		ps.u[i] = &pc
	}

	return ps
}

// data ...
func (ps *SimpleParticles) data() []Particle {
	return ps.d
}

// Reset ...
func (ps *SimpleParticles) Reset() {
	for k := range ps.u {
		ps.d[k].SetPoint(ps.u[k].Point())
		ps.d[k].SetVector(ps.u[k].Vector())
	}
}
