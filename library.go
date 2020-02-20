package main

type Book struct {
	Score int
	ID    int
}

type Library struct {
	Books       []Book
	SignupTime  int
	BooksPerDay int
}

func (l *Library) BookValueSum(numberOfDays int) int {
	numberOfDays -= l.SignupTime
	sum := 0
	for days := 0; days < numberOfDays; days++ {
		for books := 0; books < l.BooksPerDay; books++ {
			sum += l.Books[(l.BooksPerDay*days)+books].Score
		}
	}
	return sum
}

func (l *Library) Sort(books []int) {
	unsortedBooks := make([]Book, 0)
	// remove
	for i := range l.Books {
		present := false
		for j := 0; j < len(books); j++ {
			if l.Books[i].ID == books[j] {
				present = true
			}
		}
		if !present {
			unsortedBooks = append(unsortedBooks, l.Books[i])
		}
	}

	operationHappened := true
	for operationHappened == true {
		operationHappened = false
		for i := 0; i < len(unsortedBooks); i++ {
			for j := i; j < len(unsortedBooks); j++ {
				if unsortedBooks[i].Score > unsortedBooks[j].Score {
					unsortedBooks[i], unsortedBooks[j] = unsortedBooks[j], unsortedBooks[i]
					operationHappened = true
				}
			}
		}
	}
}
