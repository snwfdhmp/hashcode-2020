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
	maxRentability := 0.0
	totalScore := 0
	maxSumI := 0
	remainingDays := c.DayMax
	for iSortedLibs := 0; iSortedLibs < len(c.Libraries); iSortedLibs++ {
		startTime := time.Now()
		for i := range c.Libraries {
			c.Libraries[i].Sort(passedBooks)
			libSums[i] = c.Libraries[i].BookValueSum(remainingDays, passedBooks)
			rentability := float64(libSums[i]) / float64(max(c.Libraries[i].SignupTime+1, 1))
			// fmt.Printf("rentability: %.2f\n", rentability)
			if rentability > maxRentability {
				maxRentability = rentability
				maxSum = libSums[i]
				maxSumI = i
			}
		}

		totalScore += maxSum
		// fmt.Printf("\nSum: %d\n", maxSum)
		if iSortedLibs%9 == 0 {
			fmt.Printf("\rProgress : %d/%d (%dms) (current score: %d)", iSortedLibs+1, len(plan.SortedLibraries), time.Now().Sub(startTime).Milliseconds(), totalScore)
		}
		if maxSum == 0 {
			break
		}

		plan.SortedLibraries[iSortedLibs] = c.Libraries[maxSumI]
		for i := range c.Libraries[maxSumI].Books {
			for j := range c.Libraries {
				if j == maxSumI {
					continue
				}
				c.Libraries[j].SetBookAsUsed(c.Libraries[maxSumI].Books[i].ID)
			}
		}
		remainingDays -= c.Libraries[maxSumI].SignupTime
		c.Libraries[maxSumI] = c.Libraries[len(c.Libraries)-1]
		c.Libraries = c.Libraries[:len(c.Libraries)-1]
		maxSum = 0
		maxRentability = 0
		maxSumI = 0
		// if iSortedLibs == 1000 {
		// 	break
		// }
	}
	fmt.Printf("\n")

	return plan
}
