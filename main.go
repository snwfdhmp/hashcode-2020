package main

import (
	"fmt"
	"os"
)

var (
	filename = "c_incunabula.txt"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("fatal: %v\n", err)
	}
}

func run() error {
	ctx, err := parseFile("./input/" + filename)
	if err != nil {
		return err
	}

	plan := ctx.CreatePlan()
	file, err := os.OpenFile("./output/sorted_"+filename, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	return plan.Write(file)
}
