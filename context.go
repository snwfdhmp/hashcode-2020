package main

import "fmt"

type Context struct {
	Libraries []Library
	DayMax    int
}

func (c *Context) CreatePlan() Plan {
	libScores := make(map[int]int)
	for i := range c.Libraries {
		libScores[i] = c.Libraries[i].BookValueSum(c.DayMax)
	}

	fmt.Printf("%#v\n", libScores)
	return Plan{}
}
