package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var (
	filename   = flag.String("file", "a_example.txt", "file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

func init() {
	flag.Parse()
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
	file, err := os.OpenFile("./output/sorted_"+*filename, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	return plan.Write(file)
}
