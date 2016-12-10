package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/daved/simpartsim"
)

func main() {
	dir := "./"
	frames := 1000
	frameLen := .1
	size := 100.0
	gravity := 9.81
	drag := 9.0
	parts := 50
	dump := true

	flag.StringVar(&dir, "dir", dir, "output directory")
	flag.IntVar(&frames, "frames", frames, "frame count")
	flag.IntVar(&parts, "parts", parts, "particle count")
	flag.Parse()

	spc := simpartsim.NewSimpleSpace(size, frameLen, gravity, drag)
	ps := simpartsim.NewSimpleParticles(dir, parts, spc.Termination())

	if err := spc.Run(ps, frames, dump); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
