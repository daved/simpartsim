package simpartsim

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// Particle ...
type Particle interface {
	Point() Coords
	SetPoint(Coords)
	Vector() Coords
	SetVector(Coords)
	String() string
}

// Particles ...
type Particles interface {
	data() []Particle
	dump(int) error
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

// String ...
func (p *SimpleParticle) String() string {
	return p.pt.String()
}

// SimpleParticles ...
type SimpleParticles struct {
	dir string
	d   []Particle
}

// NewSimpleParticles ...
func NewSimpleParticles(dir string, ct int, termination Coords) *SimpleParticles {
	ps := &SimpleParticles{
		dir: dir,
		d:   make([]Particle, ct),
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

// dump ...
func (ps *SimpleParticles) dump(i int) error {
	name := filepath.Join(ps.dir, fmt.Sprintf("particleData-%d.csv", i))
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	if _, err := f.WriteString("X Axis,Y Axis,Z Axis\n"); err != nil {
		return err
	}

	for k := range ps.d {
		if _, err := f.WriteString(ps.d[k].String() + "\n"); err != nil {
			return err
		}
	}

	return nil
}
