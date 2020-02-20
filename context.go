package main

import "fmt"

type Context struct {
	Libraries []Library
	DayMax    int
}

func (c *Context) CreatePlan() Plan {
	plan := Plan{}
	plan.SortedLibraries = make([]Library, len(c.Libraries))

	passedBooks := make([]int, 0) //id of books
	libSums := make([]int, len(c.Libraries))
	maxSum := 0
	maxSumI := 0
	for iSortedLibs := 0; iSortedLibs < len(c.Libraries); iSortedLibs++ {
		for i := range c.Libraries {
			c.Libraries[i].Sort(passedBooks)
			libSums[i] = c.Libraries[i].BookValueSum(c.DayMax)
			if libSums[i] > maxSum {
				maxSum = libSums[i]
				maxSumI = i
			}
		}

		plan.SortedLibraries[iSortedLibs] = c.Libraries[maxSumI]
		for i := range c.Libraries[maxSumI].Books {
			passedBooks = append(passedBooks, c.Libraries[maxSumI].Books[i].ID)
		}
		// c.Libraries = append(c.Libraries[:iSortedLibs], c.Libraries[iSortedLibs+1:]...)
		c.Libraries[iSortedLibs] = c.Libraries[len(c.Libraries)-1]
		c.Libraries[len(c.Libraries)-1] = Library{}
		fmt.Printf("Library %d found => %d\n", iSortedLibs, maxSumI)
	}

	fmt.Printf("%s\n", plan.String())

	return Plan{}
}
