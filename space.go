package main

// space ...
type space struct {
	origin, termination     point
	frameLen, gravity, drag float64
}

// newSpace ...
func newSpace(size, frameLen, gravity, drag float64) *space {
	return &space{
		origin:      makePoint(0, 0, 0),
		termination: makePoint(size, size, size),
		frameLen:    frameLen,
		gravity:     gravity,
		drag:        drag,
	}
}

// tick ...
func (s *space) tick(ps particles) {
	d := ps.data()
	for k := range d {
		d[k].increment(s.frameLen, s.gravity)
		d[k].processCollisions(s.origin, s.termination)
		d[k].drag(s.frameLen, s.drag)
	}
}

// run ...
func (s *space) run(ps particles, frames int, dump bool) error {
	for i := 0; i < frames; i++ {
		s.tick(ps)

		if dump {
			if err := ps.dump(i); err != nil {
				return err
			}
		}
	}

	return nil
}
