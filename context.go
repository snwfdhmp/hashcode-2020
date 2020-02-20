package main

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
	remainingDays := c.DayMax
	for iSortedLibs := 0; iSortedLibs < len(c.Libraries); iSortedLibs++ {
		for i := range c.Libraries {
			c.Libraries[i].Sort(passedBooks)
			libSums[i] = c.Libraries[i].BookValueSum(remainingDays, passedBooks)
			if libSums[i] > maxSum {
				maxSum = libSums[i]
				maxSumI = i
			}
		}
		plan.SortedLibraries[iSortedLibs] = c.Libraries[maxSumI]
		for i := range c.Libraries[maxSumI].Books {
			passedBooks = append(passedBooks, c.Libraries[maxSumI].Books[i].ID)
		}
		remainingDays -= c.Libraries[maxSumI].SignupTime
		// c.Libraries = append(c.Libraries[:iSortedLibs], c.Libraries[iSortedLibs+1:]...)
		c.Libraries[maxSumI] = c.Libraries[len(c.Libraries)-1]
		c.Libraries[len(c.Libraries)-1] = Library{}
		maxSum = 0
		maxSumI = 0
	}

	return plan
}
