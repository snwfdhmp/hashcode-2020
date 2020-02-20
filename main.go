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

	for i, v := range ctx.Libraries {
		fmt.Printf("Lib %d: %d books\n", i, len(v.Books))
	}

	ctx.CreatePlan()

	return nil
}
