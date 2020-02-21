package main

import (
	"fmt"
	"time"
)

type Context struct {
	Libraries []Library
	DayMax    int
}

func (c *Context) CreatePlan() Plan {
	plan := Plan{}
	plan.SortedLibraries = make([]Library, len(c.Libraries))
	passedBooks := make(map[int]bool, 0) //id of books
	libSums := make([]int, len(c.Libraries))
	maxSum := 0
	maxSumI := 0
	remainingDays := c.DayMax
	for iSortedLibs := 0; iSortedLibs < len(c.Libraries); iSortedLibs++ {
		startTime := time.Now()
		for i := range c.Libraries {
			c.Libraries[i].Sort(passedBooks)
			libSums[i] = c.Libraries[i].BookValueSum(remainingDays, passedBooks)
			if libSums[i] > maxSum {
				maxSum = libSums[i]
				maxSumI = i
			}
		}

		fmt.Printf("\rProgress : %d/%d (%.1fs)", iSortedLibs+1, len(plan.SortedLibraries), time.Now().Sub(startTime).Seconds())
		// if maxSum == 0 {
		// 	break
		// }

		plan.SortedLibraries[iSortedLibs] = c.Libraries[maxSumI]
		for i := range c.Libraries[maxSumI].Books {
			passedBooks[c.Libraries[maxSumI].Books[i].ID] = true
		}
		remainingDays -= c.Libraries[maxSumI].SignupTime
		// c.Libraries = append(c.Libraries[:iSortedLibs], c.Libraries[iSortedLibs+1:]...)
		c.Libraries[maxSumI] = c.Libraries[len(c.Libraries)-1]
		c.Libraries[len(c.Libraries)-1] = Library{}
		maxSum = 0
		maxSumI = 0
		// if iSortedLibs == 1000 {
		// 	break
		// }
	}
	fmt.Printf("\n")

	return plan
}
