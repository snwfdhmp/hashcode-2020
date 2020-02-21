package main

type Book struct {
	Score int
	ID    int
}

type Library struct {
	ID          int
	Books       []Book
	SignupTime  int
	BooksPerDay int
}

func (l *Library) BookValueSum(numberOfDays int, passedBooks map[int]bool) int {
	numberOfDays -= l.SignupTime
	sum := 0
	for iBook := 0; iBook < l.BooksPerDay*numberOfDays && iBook < len(l.Books); iBook++ {
		if _, ok := passedBooks[l.Books[iBook].ID]; ok {
			continue
		}
		sum += l.Books[iBook].Score
	}
	return sum
}

func (l *Library) Sort(books map[int]bool) {
	unsortedBooks := make([]Book, 0)
	junkBooks := make([]Book, 0)

	// remove
	for i := range l.Books {
		if _, ok := books[l.Books[i].ID]; ok {
			junkBooks = append(junkBooks, l.Books[i])
		} else {
			unsortedBooks = append(unsortedBooks, l.Books[i])
		}
	}

	operationHappened := true
	for operationHappened == true {
		operationHappened = false
		for i := 0; i < len(unsortedBooks)-1; i++ {
			if unsortedBooks[i].Score < unsortedBooks[i+1].Score {
				unsortedBooks[i], unsortedBooks[i+1] = unsortedBooks[i+1], unsortedBooks[i]
				operationHappened = true
			}
		}
	}

	l.Books = append(unsortedBooks, junkBooks...)
}
