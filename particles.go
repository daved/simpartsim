package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// particle ...
type particle interface {
	increment(time, gravity float64)
	processCollisions(origin, termination point)
	String() string
}

// particles ...
type particles interface {
	data() []particle
	dump(int) error
}

// simpleParticle ...
type simpleParticle struct {
	point
	vector
}

// newSimpleParticle ...
func newSimpleParticle(num, ct float64, termination point) *simpleParticle {
	t := termination
	return &simpleParticle{
		point:  makePoint(num/ct*t.x, rand.Float64()*t.y, rand.Float64()*t.z),
		vector: makeVector(rand.Float64()*5, 0, 0),
	}
}

// incrementDisplacement ...
func (p *simpleParticle) incrementDisplacement(time, gravity float64) {
	if p.point.y == 0 && p.vector.y == 0 {
		return
	}

	p.vector.y -= gravity * time
}

// incrementLocation ...
func (p *simpleParticle) incrementLocation(time float64) {
	p.point = makePoint(
		p.point.x+p.vector.x*time,
		p.point.y+p.vector.y*time,
		p.point.z+p.vector.z*time,
	)
}

// increment ...
func (p *simpleParticle) increment(time, gravity float64) {
	p.incrementDisplacement(time, gravity)
	p.incrementLocation(time)
}

// reflect ...
func (p *simpleParticle) reflect(a, b float64) float64 {
	diff := a - b
	return b - diff
}

// processCollisions ...
func (p *simpleParticle) processCollisions(origin, termination point) {
	if p.point.x < origin.x {
		p.point.x = p.reflect(p.point.x, origin.x)
		p.vector.x *= -.6
	}
	if p.point.x > termination.x {
		p.point.x = p.reflect(p.point.x, termination.x)
		p.vector.x *= -.6
	}

	if p.point.y < origin.y {
		p.point.y = p.reflect(p.point.y, origin.y)
		p.vector.y *= -.6
	}
	if p.point.y > termination.y {
		p.point.y = p.reflect(p.point.y, termination.y)
		p.vector.y *= -.6
	}

	if p.point.z < origin.z {
		p.point.z = p.reflect(p.point.z, origin.z)
		p.vector.z *= -.6
	}
	if p.point.z > termination.z {
		p.point.z = p.reflect(p.point.z, termination.z)
		p.vector.z *= -.6
	}
}

// String ...
func (p *simpleParticle) String() string {
	return p.point.String()
}

// simpleParticles ...
type simpleParticles struct {
	dir string
	d   []particle
}

// newSimpleParticles ...
func newSimpleParticles(dir string, ct int, termination point) *simpleParticles {
	ps := &simpleParticles{
		dir: dir,
		d:   make([]particle, ct),
	}

	for i := 0; i < ct; i++ {
		ps.d[i] = newSimpleParticle(float64(i), float64(ct), termination)
	}

	return ps
}

// data ...
func (ps *simpleParticles) data() []particle {
	return ps.d
}

// dump ...
func (ps *simpleParticles) dump(i int) error {
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
