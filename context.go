package main

import (
	"fmt"
	// "math"
	"time"
)

type Context struct {
	Libraries []Library
	DayMax    int
}

func (c *Context) CreatePlan() Plan {
	plan := Plan{}
	plan.SortedLibraries = make([]Library, len(c.Libraries))
	totalScore := 0
	remainingDays := c.DayMax
	for iSortedLibs := 0; iSortedLibs < len(c.Libraries); iSortedLibs++ {
		startTime := time.Now()
		maxRentability := 0.0
		maxSum := 0
		maxSumI := 0
		for i := range c.Libraries {
			c.Libraries[i].Sort()
			sum := c.Libraries[i].BookValueSum(remainingDays)
			rentabilityCoef := 0.00000001
			if c.Libraries[i].SignupTime > 0 {
				rentabilityCoef = float64(c.Libraries[i].SignupTime) / float64(remainingDays)
			}
			rentability := float64(sum) / rentabilityCoef
			// fmt.Printf("rentability: %.2f\n", rentability)
			if rentability > maxRentability {
				maxRentability = rentability
				maxSum = sum
				maxSumI = i
			}
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

		totalScore += maxSum
		// fmt.Printf("\nSum: %d\n", maxSum)
		if iSortedLibs%9 == 0 || maxSum == 0 {
			fmt.Printf("\rProgress : %d/%d (%dms) (current score: %d)", iSortedLibs+1, len(plan.SortedLibraries), time.Now().Sub(startTime).Milliseconds(), totalScore)
		}
		if maxSum == 0 {
			break
		}

		remainingDays -= c.Libraries[maxSumI].SignupTime
		c.Libraries[maxSumI] = c.Libraries[len(c.Libraries)-1]
		c.Libraries = c.Libraries[:len(c.Libraries)-1]
		// if iSortedLibs == 1000 {
		// 	break
		// }
	}
	fmt.Printf("\n")

	return plan
}
