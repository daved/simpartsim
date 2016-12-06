package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dir := "./"
	frames := 1000
	frameLen := .1
	size := 100.0
	gravity := 9.81
	parts := 50
	dump := true

	flag.StringVar(&dir, "dir", dir, "output directory")
	flag.IntVar(&frames, "frames", frames, "frame count")
	flag.IntVar(&parts, "parts", parts, "particle count")
	flag.Parse()

	spc := newSpace(size, frameLen, gravity)
	ps := newSimpleParticles(dir, parts, spc.termination)

	if err := spc.run(ps, frames, dump); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
