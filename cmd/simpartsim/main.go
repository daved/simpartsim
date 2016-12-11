package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/daved/simpartsim"
)

func main() {
	dir := ""
	prefix := "particleData"
	parts := 500
	frames := 1000
	opts := simpartsim.SimpleSpaceOptions{
		FrameLen: .1,
		Size:     100.0,
		Gravity:  9.81,
		Drag:     9.0,
	}

	flag.StringVar(&dir, "dir", dir, "output dir (stdout if blank)")
	flag.IntVar(&parts, "parts", parts, "particle count")
	flag.IntVar(&frames, "frames", frames, "frame count")
	flag.Parse()

	spc := simpartsim.NewSimpleSpace(opts)
	ps := simpartsim.NewSimpleParticles(parts, spc.Termination())

	csc := make(chan []simpartsim.Coords)
	go func() {
		spc.Run(ps, frames, csc)
		defer close(csc)
	}()

	i := 0
	for cs := range csc {
		if dir == "" {
			if err := csvToStdout(cs); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			continue
		}

		if err := csvToFiles(dir, prefix, i, cs); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		i++
	}
}

// dump ...
func dumpCSV(w io.Writer, cs []simpartsim.Coords) error {
	if _, err := w.Write([]byte(simpartsim.CoordsFieldNames)); err != nil {
		return err
	}
	if _, err := w.Write([]byte("\n")); err != nil {
		return err
	}

	for k := range cs {
		if _, err := w.Write([]byte(cs[k].String())); err != nil {
			return err
		}
		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
	}

	return nil
}

func csvToFiles(dir, prefix string, iter int, cs []simpartsim.Coords) error {
	name := filepath.Join(dir, prefix+fmt.Sprintf("-%d.csv", iter))
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	if err := dumpCSV(f, cs); err != nil {
		return err
	}

	return nil
}

func csvToStdout(cs []simpartsim.Coords) error {
	return dumpCSV(os.Stdout, cs)
}
