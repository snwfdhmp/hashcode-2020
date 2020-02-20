package main

import (
	"fmt"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("fatal: %v\n", err)
	}
}

func run() error {
	ctx, err := parseFile("./input/b_read_on.txt")
	if err != nil {
		return err
	}

	ctx.CreatePlan()

	return nil
}
