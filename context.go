package main

import "fmt"

type Context struct {
	Libraries []Library
	DayMax    int
}

func (c *Context) CreatePlan() Plan {
	passedBooks := make([]int, 0) //id of books
	libSums := make([]int, len(c.Libraries))
	maxSum := 0
	maxSumI := 0
	for i := range c.Libraries {
		c.Libraries[i].Sort(passedBooks)
		libSums[i] = c.Libraries[i].BookValueSum(c.DayMax)
		if libSums[i] > maxSum {
			maxSum = libSums[i]
			maxSumI = i
		}
	}

	c.Libraries[maxSumI]

	fmt.Printf("%#v\n", c.Libraries)
	fmt.Printf("Max Library 1 : \n", maxSumI)

	return Plan{}
}
