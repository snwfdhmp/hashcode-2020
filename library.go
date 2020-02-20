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
