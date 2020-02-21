package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var (
	image      = flag.String("image", "", "image (a, b, c, d, e, f)")
	filename   = flag.String("file", "a_example.txt", "file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	images = map[string]string{
		"a": "a_example.txt",
		"b": "b_read_on.txt",
		"c": "c_incunabula.txt",
		"d": "d_tough_choices.txt",
		"e": "e_so_many_books.txt",
		"f": "f_libraries_of_the_world.txt",
	}
)

func init() {
	flag.Parse()

	if *image != "" {
		var ok bool
		*filename, ok = images[*image]
		if !ok {
			fmt.Printf("fatal: image '%s' does not exist.\n", image)
			os.Exit(1)
		}
	}
}

func main() {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if err := run(); err != nil {
		fmt.Printf("fatal: %v\n", err)
	}
}

func run() error {
	ctx, err := parseFile("./input/" + *filename)
	if err != nil {
		return err
	}

	plan := ctx.CreatePlan()
	filepath := "./output/sorted_" + *filename
	if err := os.Remove(filepath); err != nil {
		return err
	}

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	return plan.Write(file)
}
